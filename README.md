# firstAPI

### Beego Framework

### JWT Authtication

#### 逻辑：
    1. 检查token的有效期
    2. 如果过期 则需要重新登录
    3. 如果验证成功 但是即将过期则更新token 并返回app
