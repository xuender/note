# jwt

## jwt 自动续期

* 到期时间不应太长
* jwt的payload中要增加个版本号
* 到期后进行自动续期
  * 续期时检查版本号
  * 如果用户修改过密码等身份信息，版本号增加
  * 如果需要续期的jwt版本号与用户版本号不符，拒绝续期
