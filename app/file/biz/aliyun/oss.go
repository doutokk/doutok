package aliyun

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type GetPolicyReq struct {
	FileName   string
	ExpireTime time.Time
}

type GetPolicyResp struct {
	Host                string
	Key                 string
	Policy              string
	Signature           string
	SignatureVersion    string
	SignatureCredential string
	Date                string
	SecurityToken       string
}

func GetPolicy(req GetPolicyReq) (resp GetPolicyResp) {
	// 假设已通过 AssumeRole 获取临时凭证
	stsToken := GetSTSToken()
	accessKeyId := stsToken.AccessKeyId
	accessKeySecret := stsToken.AccessKeySecret
	securityToken := stsToken.SecurityToken

	// 配置信息
	bucketName := c.Bucket
	region := c.RegionID
	date := time.Now().UTC().Format("20060102")
	xOssDate := time.Now().UTC().Format("20060102T150405Z")

	// 构造 Policy
	policy := map[string]interface{}{
		"expiration": req.ExpireTime.Format("2006-01-02T15:04:05.000Z"),
		"conditions": []interface{}{
			map[string]string{"bucket": bucketName},
			[]interface{}{"starts-with", "$key", req.FileName},
			map[string]string{"x-oss-signature-version": "OSS4-HMAC-SHA256"},
			map[string]string{"x-oss-credential": fmt.Sprintf("%s/%s/%s/oss/aliyun_v4_request", accessKeyId, date, region)},
			map[string]string{"x-oss-date": xOssDate},
			map[string]string{"x-oss-security-token": securityToken},
			[]interface{}{"content-length-range", 0, 10485760},
		},
	}

	// 将 Policy 序列化为 JSON 字符串
	policyJSON, err := json.Marshal(policy)
	if err != nil {
		fmt.Printf("Failed to marshal policy: %v	", err)
		return
	}

	// Base64 编码 Policy
	encodedPolicy := base64.StdEncoding.EncodeToString(policyJSON)

	// 计算 SigningKey
	signingKey, err := calculateSigningKey(accessKeySecret, date, region, "oss")
	if err != nil {
		fmt.Printf("Failed to calculate signing key: %v", err)
		return
	}

	// 计算 Signature
	signature := calculateSignature(signingKey, encodedPolicy)

	// 输出结果
	fmt.Println("Policy:", encodedPolicy)
	fmt.Println("Signature:", signature)
	resp = GetPolicyResp{
		Host:                fmt.Sprintf("https://%s.%s.aliyuncs.com", bucketName, region),
		Key:                 req.FileName,
		Policy:              encodedPolicy,
		Signature:           signature,
		SignatureVersion:    "OSS4-HMAC-SHA256",
		SignatureCredential: fmt.Sprintf("%s/%s/%s/oss/aliyun_v4_request", accessKeyId, date, region),
		Date:                xOssDate,
		SecurityToken:       securityToken,
	}
	return
}

func calculateSigningKey(secretKey, date, region, service string) ([]byte, error) {
	dateKey := hmacSHA256([]byte("aliyun_v4"+secretKey), []byte(date))
	dateRegionKey := hmacSHA256(dateKey, []byte(region))
	dateRegionServiceKey := hmacSHA256(dateRegionKey, []byte(service))
	signingKey := hmacSHA256(dateRegionServiceKey, []byte("aliyun_v4_request"))
	return signingKey, nil
}

func hmacSHA256(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func calculateSignature(signingKey []byte, stringToSign string) string {
	h := hmac.New(sha256.New, signingKey)
	h.Write([]byte(stringToSign))
	return hex.EncodeToString(h.Sum(nil))
}
