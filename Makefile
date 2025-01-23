## 运行时使用命令 make 命令名称 svc=服务名
.PHONY: gen
gen: ## make gen svc=product
	@scripts/gen.sh ${svc}

.PHONY: gen-server
gen-server: ## make gen-server svc=product
	@mkdir -p app/${svc} && cd app/${svc} && cwgo server -I ../../idl  --type RPC --module github.com/PengJingzhao/douyin-commerce/app/${svc} --service ${svc}  --idl ../../idl/${svc}.proto && go work use . && go mod tidy

.PHONY: gen-client
gen-client: ## make gen-client svc=product
	@cd rpc_gen && cwgo client --type RPC --service ${svc} --module github.com/PengJingzhao/douyin-commerce/rpc_gen  -I ../idl  --idl ../idl/${svc}.proto


.PHONY: gen-model
gen-model: ## make gen-model svc=user
	@cd app/${svc} && cwgo model c ../../idl --idl ../../idl/${svc}.proto

.PHONY: tidy
tidy: ## make tidy svc=auth
	@cd app/${svc} && go mod tidy
