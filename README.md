# Android API接口文档
 

**1.手机号查询用户信息接口**

请求方式：GET 

用户请求格式：https://134.175.6.153:8899/user_basic/query/手机号码

用户请求范例：https://134.175.6.153:8899/user_basic/query/18620638848

返回 Status 200 OK 手机号码存在，查询成功

返回 Status 201 Created 该手机号码不存在




**2.修改密码接口**

请求方式：PUT

用户请求格式：https://134.175.6.153:8899/user_basic/update/手机号码/旧密码/新密码

用户请求范例：https://134.175.6.153:8899/user_basic/update/15626171233/123/456

返回 status 200 OK 修改成功

返回 status 201 Created 手机号码不存在或旧密码错误