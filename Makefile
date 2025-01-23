##@ Build

.PHONY: gen
gen: ## gen client code of {svc}. example: make gen svc=product
	@scripts/gen.sh ${svc}

.PHONY: gen-client
gen-client: ## gen client code of {svc}. example: make gen-client svc=product 为什么这里公用一个modulename,为什么要加上github.com
	@cd rpc_gen && cwgo client --type RPC --service ${svc} --module github.com/PengJingzhao/douyin-commerce/rpc_gen  -I ../idl  --idl ../idl/${svc}.proto

.PHONY: gen-server
gen-server: ## gen service code of {svc}. example: make gen-server svc=product 不加这个--pass会生成一个kitex_gen
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/PengJingzhao/douyin-commerce/app/${svc} --pass "-use github.com/puzzlehh/kill_system/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/${svc}.proto

.PHONY: gen-model
gen-model: ## Generate model code with CRUD operations. Example: make gen-model svc=user
	@cd app/${svc} && cwgo model c ../../idl --idl ../../idl/${svc}.proto
