package casbin

import (
	_ "embed"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/doutokk/doutok/common/utils"
)

var (
	enforcer *casbin.Enforcer
	err      error
	//go:embed model.conf
	modelFile []byte
	//go:embed policy.csv
	policyFile []byte
)

func Init() {

	utils.WriteFile("auth_model.conf", modelFile)
	utils.WriteFile("auth_policy.csv", policyFile)
	// 鉴权中间件
	enforcer, err = casbin.NewEnforcer("auth_model.conf", "auth_policy.csv")
	if err != nil {
		panic("auth init failed")
	}
}

// 这里还需要uri和method
func CheckAuthByRBAC(sub string, obj string, act string) bool {
	hlog.Info("sub: ", sub, " obj: ", obj, " act: ", act)

	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		hlog.Error("checkAuthByRBAC failed: ", err)
		return false
	}
	if !ok {
		return false
	}
	return true
}
