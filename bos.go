package main

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/auth"         //导入认证模块
	"github.com/baidubce/bce-sdk-go/services/bos" //导入BOS服务模块
	"github.com/baidubce/bce-sdk-go/services/sts" //导入STS服务模块
)

func main() {
	// 创建STS服务的Client对象，Endpoint使用默认值

	stsClient, err := sts.NewClient(AK, SK)
	if err != nil {
		fmt.Println("create sts client object :", err)
		return
	}

	// 获取临时认证token，有效期为60秒，ACL为空
	sts, err := stsClient.GetSessionToken(60, "")
	if err != nil {
		fmt.Println("get session token failed:", err)
		return
	}
	fmt.Println("GetSessionToken result:")
	fmt.Println("  accessKeyId:", sts.AccessKeyId)
	fmt.Println("  secretAccessKey:", sts.SecretAccessKey)
	fmt.Println("  sessionToken:", sts.SessionToken)
	fmt.Println("  createTime:", sts.CreateTime)
	fmt.Println("  expiration:", sts.Expiration)
	fmt.Println("  userId:", sts.UserId)

	// 使用申请的临时STS创建BOS服务的Client对象，Endpoint使用默认值
	bosClient, err := bos.NewClient(sts.AccessKeyId, sts.SecretAccessKey, "")
	if err != nil {
		fmt.Println("create bos client failed:", err)
		return
	}
	stsCredential, err := auth.NewSessionBceCredentials(
		sts.AccessKeyId,
		sts.SecretAccessKey,
		sts.SessionToken)
	if err != nil {
		fmt.Println("create sts credential object failed:", err)
		return
	}
	bosClient.Config.Credentials = stsCredential
}
