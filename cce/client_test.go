package cce

import (
	"github.com/baidu/baiducloud-sdk-go/bce"
	"os"
)

// Modify with your AccessKeyID and SecretAccessKey
var (
	TestAccessKeyID     = os.Getenv("AccessKeyID")
	TestSecretAccessKey = os.Getenv("SecretAccessKey")
	TestRegion          = os.Getenv("Region")
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		config := bce.NewConfigWithParams(TestAccessKeyID, TestSecretAccessKey, TestRegion)
		config.Protocol = "https"
		testClient = NewClient(config)
	}
	return testClient
}
