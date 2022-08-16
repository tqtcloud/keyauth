package utils

import "github.com/infraboard/mcube/exception"

// IsAccessTokenExpiredError 访问令牌是否过期错误
func IsAccessTokenExpiredError(err error) bool {
	if err == nil {
		return false
	}

	e, ok := err.(exception.APIException)
	if !ok {
		return false
	}

	return e.ErrorCode() == exception.AccessTokenExpired && e.Namespace() == exception.GlobalNamespace.String()
}
