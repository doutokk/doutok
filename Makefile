export ROOT_MOD=github.com/PengJingzhao/douyin-commerce
# 项目相关变量
SERVICE_NAME := douyin-commerce      # 微服务名称
MODULE_NAME := douyin-commerce       # Go 模块名称
BUILD_DIR := ./bin                     # 构建输出目录
GO := go                               # Go 命令的别名
KITEX := kitex                         # Kitex 工具的别名
HERTZ := hertz                         # Hertz 工具的别名
PROTO_DIR := ./idl                     # Proto 文件目录
MAIN_FILE := ./cmd/main.go             # 项目的主入口

# 构建参数
GO_FLAGS := -ldflags="-s -w"           # 构建标志 (优化后的 Go 二进制)
GO_TEST_FLAGS := -v -cover            # 测试标志

# 默认目标
.PHONY: all
all: build

# 生成代码（Kitex 和 Hertz）
.PHONY: generate
generate:
	@echo "Generating Kitex and Hertz code..."
	$(KITEX) -module $(MODULE_NAME) -service $(SERVICE_NAME) -use protobuf $(PROTO_DIR)
	$(HERTZ) -module $(MODULE_NAME) -idl $(PROTO_DIR)

# 构建项目
.PHONY: build
build: generate
	@echo "Building the project..."
	$(GO) build $(GO_FLAGS) -o $(BUILD_DIR)/$(SERVICE_NAME) $(MAIN_FILE)

# 运行微服务
.PHONY: run
run:
	@echo "Running the service..."
	$(GO) run $(MAIN_FILE)

# 运行测试
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test $(GO_TEST_FLAGS) ./...

# 清理构建输出文件
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# 更新依赖
.PHONY: deps
deps:
	@echo "Updating dependencies..."
	$(GO) mod tidy
	$(GO) mod vendor

# 静态检查（如代码格式化、lint）
.PHONY: lint
lint:
	@echo "Running lint checks..."
	golangci-lint run

# Docker 构建（可选）
.PHONY: docker
docker:
	@echo "Building Docker image..."
	docker build -t $(SERVICE_NAME):latest .


.PHONY: gen-cart
gen-cart:
	@cd cart-service && cwgo server --type RPC  --server_name cart --module  ${ROOT_MOD}/cart-service  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../idl  --idl ../idl/cart.proto
	@cd rpc_gen && cwgo client --type RPC  --server_name cart --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/cart.proto