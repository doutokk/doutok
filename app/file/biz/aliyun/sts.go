package aliyun

import (
	"fmt"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func main() {
	// 从环境变量中获取步骤 1.1 生成的 RAM 用户的访问密钥（AccessKey ID 和 AccessKey Secret）。
	accessKeyId := os.Getenv("ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")
	// 从环境变量中获取步骤 1.3 生成的 RAM 角色的 RamRoleArn。
	roleArn := os.Getenv("RAM_ROLE_ARN")

	// 创建权限策略客户端。
	config := &openapi.Config{
		// 必填，步骤 1.1 获取到的 AccessKey ID。
		AccessKeyId: tea.String(accessKeyId),
		// 必填，步骤 1.1 获取到的 AccessKey Secret。
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Sts
	config.Endpoint = tea.String("sts.cn-hangzhou.aliyuncs.com")
	client, err := sts20150401.NewClient(config)
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		return
	}

	// 使用 RAM 用户的 AccessKey ID 和 AccessKey Secret 向 STS 申请临时访问凭证。
	request := &sts20150401.AssumeRoleRequest{
		// 指定 STS 临时访问凭证过期时间为 3600 秒。
		DurationSeconds: tea.Int64(3600),
		// 从环境变量中获取步骤 1.3 生成的 RAM 角色的 RamRoleArn。
		RoleArn: tea.String(roleArn),
		// 指定自定义角色会话名称，这里使用和第一段代码一致的 examplename
		RoleSessionName: tea.String("examplename"),
	}
	response, err := client.AssumeRoleWithOptions(request, &util.RuntimeOptions{})
	if err != nil {
		fmt.Printf("Failed to assume role: %v\n", err)
		return
	}

	// 打印 STS 返回的临时访问密钥（AccessKey ID 和 AccessKey Secret）、安全令牌（SecurityToken）以及临时访问凭证过期时间（Expiration）。
	credentials := response.Body.Credentials
	fmt.Println("AccessKeyId: " + tea.StringValue(credentials.AccessKeyId))
	fmt.Println("AccessKeySecret: " + tea.StringValue(credentials.AccessKeySecret))
	fmt.Println("SecurityToken: " + tea.StringValue(credentials.SecurityToken))
	fmt.Println("Expiration: " + tea.StringValue(credentials.Expiration))
}
