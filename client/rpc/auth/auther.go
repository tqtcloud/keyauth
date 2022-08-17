package auth

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/tqtcloud/keyauth/apps/token"
	"github.com/tqtcloud/keyauth/client/rpc"
)

func NewKeyauthAuther(client *rpc.ClientSet, serviceName string) *KeyauthAuther {
	return &KeyauthAuther{
		auth:        client.Token(),
		log:         zap.L().Named("http.auther"),
		serviceName: serviceName,
	}
}

// 有Keyauth提供的 HTTP认证中间件
type KeyauthAuther struct {
	log         logger.Logger
	auth        token.ServiceClient
	serviceName string
}
