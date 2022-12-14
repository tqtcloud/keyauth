package all

import (
	_ "github.com/tqtcloud/keyauth/apps/audit/impl"
	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	_ "github.com/tqtcloud/keyauth/apps/book/impl"
	_ "github.com/tqtcloud/keyauth/apps/endpoint/impl"
	_ "github.com/tqtcloud/keyauth/apps/policy/impl"
	_ "github.com/tqtcloud/keyauth/apps/role/impl"
	_ "github.com/tqtcloud/keyauth/apps/token/impl"
	_ "github.com/tqtcloud/keyauth/apps/user/impl"
)
