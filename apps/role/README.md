# 角色绑定
> 创建api：http://{{127.0.0.1}}:8050/keyauth/api/v1/role

## 创建指定服务功能角色
```json
{
    "name": "member",
    "description": "成员",
    "permissions": [
        {
            "service": "cmdb",
            "featrues": [
                {
                    "resource": "secret",
                    "action": "list"
                },
                {
                    "resource": "secret",
                    "action": "get"
                }
            ]
        }
    ]
}
```

## 创建admin角色，跳过所有鉴权
```json
{
    "name": "admin",
    "description": "admin角色跳过所有鉴权",
    "permissions": [
        {
            "service": "cmdb",
            "allow": true
        }
    ]
}
```