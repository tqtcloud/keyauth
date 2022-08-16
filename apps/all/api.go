package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github/tqtcloud/keyauth/apps/book/api"
	_ "github/tqtcloud/keyauth/apps/token/api"
	_ "github/tqtcloud/keyauth/apps/user/api"
)
