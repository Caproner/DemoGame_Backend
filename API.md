# 接口文档

## 测试用接口

### 简单的测试接口

```
GET IP:17263/api/test/echo
```

#### 请求参数

空

#### 返回格式

```
{
    "ret": 0,         // 非0则表示失败
    "msg": "ok",      // 当ret非0时这里给出错误信息
    "data": {
        "msg": "rsp ok, this is demo backend"       // 测试信息
    }
}
```