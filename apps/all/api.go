package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/tqtcloud/keyauth/apps/book/api"
	_ "github.com/tqtcloud/keyauth/apps/policy/api"
	_ "github.com/tqtcloud/keyauth/apps/role/api"
	_ "github.com/tqtcloud/keyauth/apps/token/api"
	_ "github.com/tqtcloud/keyauth/apps/user/api"
)
