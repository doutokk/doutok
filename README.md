# douyin-commerce
抖音电商项目

## 增加模块教程
1. 在 app 包创建一个子模块包
2. go mod init
进入子模块包，执行 go mod init
如 go mod init github.com/doutokk/doutok/app/auth
3. 在 go.work 加入子模块包
4. 进入项目根目录
make gen svc=auth