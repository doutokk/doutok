package main

import (
	"context"
	"embed"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/doutokk/doutok/app/frontend/biz/router"
	"github.com/doutokk/doutok/app/frontend/conf"
	"github.com/doutokk/doutok/app/frontend/infra/mtl"
	"github.com/doutokk/doutok/app/frontend/infra/rpc"
	"github.com/doutokk/doutok/app/frontend/middleware"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	hertzotelprovider "github.com/hertz-contrib/obs-opentelemetry/provider"
	hertzoteltracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	oteltrace "go.opentelemetry.io/otel/trace"
)

//go:embed template/*
var templateFS embed.FS

//go:embed static/*
var staticFS embed.FS

// releaseStaticFiles 将嵌入的 static 文件释放到目标目录，若目录不存在则创建
func releaseStaticFiles(targetDir string) error {
	// 如果目标文件夹存在则跳过释放，可以根据需要检测是否为空
	if info, err := os.Stat("static"); err == nil && info.IsDir() {
		return nil
	} else if os.IsNotExist(err) {
		if err := os.MkdirAll("static", 0755); err != nil {
			return err
		}
	}

	// 遍历 embedded 的 static 文件夹
	return fs.WalkDir(staticFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		destPath := filepath.Join(".", path)
		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}
		content, err := staticFS.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(destPath, content, 0644)
	})
}

func main() {
	_ = godotenv.Load()

	// 释放 static 文件到本地文件夹（仅在启动时生效）
	if err := releaseStaticFiles("./static"); err != nil {
		log.Fatalf("释放 static 文件失败: %v", err)
	}

	mtl.InitMtl()
	rpc.InitClient()
	address := conf.GetConf().Hertz.Address

	p := hertzotelprovider.NewOpenTelemetryProvider(
		hertzotelprovider.WithSdkTracerProvider(mtl.TracerProvider),
		hertzotelprovider.WithEnableMetrics(false),
	)
	defer p.Shutdown(context.Background())
	tracer, cfg := hertzoteltracing.NewServerTracer(hertzoteltracing.WithCustomResponseHandler(func(ctx context.Context, c *app.RequestContext) {
		c.Header("shop-trace-id", oteltrace.SpanFromContext(ctx).SpanContext().TraceID().String())
	}))

	h := server.New(server.WithHostPorts(address), server.WithTracer(
		hertzprom.NewServerTracer(
			"",
			"",
			hertzprom.WithRegistry(mtl.Registry),
			hertzprom.WithDisableServer(true),
		),
	),
		tracer,
	)

	// 使用 html/template 从 embed.FS 中解析模板文件
	tmpl, err := template.ParseFS(templateFS, "template/*")
	if err != nil {
		log.Fatalf("模板解析失败: %v", err)
	}
	h.SetHTMLTemplate(tmpl)
	h.Delims("{{", "}}")

	h.Use(hertzoteltracing.ServerMiddleware(cfg))
	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)

	h.GET("sign-in", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-in", utils.H{
			"title": "Sign in",
			"next":  c.Query("next"),
		})
	})
	h.GET("sign-up", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-up", utils.H{
			"title": "Sign up",
		})
	})
	h.GET("/redirect", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "about", utils.H{
			"title": "Error",
		})
	})
	if os.Getenv("GO_ENV") != "online" {
		h.GET("/robots.txt", func(ctx context.Context, c *app.RequestContext) {
			c.Data(consts.StatusOK, "text/plain", []byte(`User-agent: *
Disallow: /`))
		})
	}

	h.Static("/static", "./")

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	store, err := redis.NewStore(100, "tcp", conf.GetConf().Redis.Address, "", []byte("114514"))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{MaxAge: 86400, Path: "/"})
	rs, err := redis.GetRedisStore(store)
	if err == nil {
		rs.SetSerializer(sessions.JSONSerializer{})
	}
	h.Use(sessions.New("cloudwego-shop", store))

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	logger := hertzzap.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelInfo)
	hlog.SetOutput(os.Stdout)

	// recovery
	h.Use(recovery.Recovery())

	h.OnShutdown = append(h.OnShutdown, mtl.Hooks...)

	// cores
	h.Use(cors.Default())
	middleware.RegisterMiddleware(h)
}
