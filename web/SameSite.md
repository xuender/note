# Cookie 的 SameSite 属性

* Strict 完全禁止第三方 Cookie，
* Lax 导航到目标网址的 Get 请求除外
* None

设置必须增加 `Secure`，只有`https`环境才有效

`Set-Cookie: key=value; SameSite=None; Secure`
