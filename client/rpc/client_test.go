package rpc_test

import (
	"context"
	"fmt"
	"github.com/tqtcloud/keyauth/apps/audit"
	"github.com/tqtcloud/keyauth/apps/policy"
	"github.com/tqtcloud/keyauth/apps/token"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tqtcloud/keyauth/client/rpc"

	mcenter "github.com/infraboard/mcenter/client/rpc"
)

// keyauth 客户端
// 需要配置注册中心的地址
// 获取注册中心的客户端，使用注册中心的客户端 查询 keyauth的地址
func TestBookQuery(t *testing.T) {
	should := assert.New(t)

	conf := mcenter.NewDefaultConfig()
	conf.Address = os.Getenv("MCENTER_ADDRESS")
	conf.ClientID = os.Getenv("MCENTER_CDMB_CLINET_ID")
	conf.ClientSecret = os.Getenv("MCENTER_CMDB_CLIENT_SECRET")
	fmt.Println(conf.ClientSecret)

	// 传递Mcenter配置, 客户端通过Mcenter进行搜索, New一个用户中心的客户端
	keyauthClient, err := rpc.NewClient(conf)

	// 使用SDK 调用Keyauth进行 凭证的校验
	// c.Token().ValidateToken()

	// 进行服务功能注册
	// keyauthClient.Endpoint().RegistryEndpoint()

	// 鉴权校验
	// keyauthClient.Policy().ValidatePermission()

	if should.NoError(err) {
		resp, err := keyauthClient.Token().ValidateToken(
			context.Background(),
			token.NewValidateTokenRequest("D2K74FPU7rCAGggT0M5rgvKn"),
		)
		should.NoError(err)
		fmt.Println(resp)
	}
}

// 测试鉴权
func TestValidatePermission(t *testing.T) {
	should := assert.New(t)

	conf := mcenter.NewDefaultConfig()
	conf.Address = os.Getenv("MCENTER_ADDRESS")
	conf.ClientID = os.Getenv("MCENTER_CDMB_CLINET_ID")
	conf.ClientSecret = os.Getenv("MCENTER_CMDB_CLIENT_SECRET")

	// 传递Mcenter配置, 客户端通过Mcenter进行搜索, New一个用户中心的客户端
	keyauthClient, err := rpc.NewClient(conf)
	if should.NoError(err) {
		req := policy.NewValidatePermissionRequest()
		req.Username = "member"
		req.Service = "cmdb"
		req.Resource = "secret"
		req.Action = "delete"

		p, err := keyauthClient.Policy().ValidatePermission(context.TODO(), req)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(p)
	}
}

func TestAuditOperate(t *testing.T) {
	should := assert.New(t)

	conf := mcenter.NewDefaultConfig()
	conf.Address = os.Getenv("MCENTER_ADDRESS")
	conf.ClientID = os.Getenv("MCENTER_CDMB_CLINET_ID")
	conf.ClientSecret = os.Getenv("MCENTER_CMDB_CLIENT_SECRET")

	// 传递Mcenter配置, 客户端通过Mcenter进行搜索, New一个用户中心的客户端
	keyauthClient, err := rpc.NewClient(conf)
	if should.NoError(err) {
		req := audit.NewOperateLog("member", "secret", "delete")
		p, err := keyauthClient.Audit().AuditOperate(context.TODO(), req)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(p)
	}
}

func init() {
	// 提前加载好 mcenter客户端, resolver需要使用
	err := mcenter.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
