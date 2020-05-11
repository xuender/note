# hash 算法

## SipHash

让输出更随机化，防止哈希洪水(hash flooding)问题，速度快，可以生成long

## HighwayHash

据说和 SipHash 一样优秀，速度更快。

## 比较

单位纳秒

|  数据 | SipHash | HighwayHash |
| ----: | ------: | ----------: |
|     5 |    91.7 |         164 |
|    50 |     120 |         176 |
|   500 |     359 |         218 |
|  5000 |    2782 |         593 |
| 50000 |   26563 |        4261 |

hash数据较少的时候 `siphash` 微弱优势，数据较大时 `highwayhash`优势明显
