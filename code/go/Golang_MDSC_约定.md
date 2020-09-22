# Golang MDSC 约定

`Model`和`Ctrl`是必须的，如果`Ctrl`方法要单元测试或复用则需抽象出`Service`。

有数据库操作可创建`DAO`，`Model`、`DAO`和`Service`使用单元测试，

`Ctrl`使用`RESTClient`的`.http`文件测试。

## 分层结构

* `Model` (模型)
  * 可以包含少量方法；
  * 不引用外部资源；
  * 功能：
    * 数据传输；
    * 持久化；
    * 字段校验；
    * `json`、`xml`格式定义；
  * 如无必要，转换`json`、`xml`格式时无需改变字段大小写，减少无谓的工作量；
* `DAO`(数据操作接口)
  * 读写数据库；
  * 通用创建、获取、保存不要重复实现，所有方法必须有原因：
    * 包含`sql`；
    * 复合语句；
    * 异常处理；
  * 可被`ctrl`和`service`调用；
  * 不调用其他资源；
  * 数据库操作对象通过参数传入，方便事务控制；
  * 方法小写开头，避免跨包调用；
* `Service` (服务)
  * 提供被多方使用的公共方法；
  * 名称以`Service`结尾；
  * 会被多个`Ctrl`或`Service`调用；
  * 可创建、输入、输出`Model`；
  * 可调用其他`Service`、`DAO`；
  * 不可调用`Ctrl`；
  * 主要的单元测试对象，覆盖率要足够高；
  * 提供一个`New`方法创建服务，方法的参数是需要引用的其他服务；
  * 两个`Service`避免循环引用，如果无法避免，可赋值引用：
* `Ctrl` (控制器)
  * 对外提供`WEB`交互服务；
  * 名称以`Ctrl`结尾；
  * 包含一个路由方法(`Party`)，定义`RESTful`调用名称、路径、模式、参数等；
  * `Ctrl`不能被`Model`和`Service`调用；
  * `Ctrl`可以调用其他`Ctrl`，但只能用于路由分发；
  * 为每个`Ctrl`都提供同名的`http`测试文件；
  * 提供一个`New`方法创建控制器，方法的参数是需要引用的服务；
  * 方法开头使用`get`、`post`、`put`、`delete`、`patch`等说明调用方式；
* 对象间依赖使用 [wire](https://github.com/google/wire)

## Ctrl & RESTful

* `Ctrl`包含一个路由方法(`Party`)；
* 路由指定的执行方法使用模式开头(`get`、`post`、`put`、`patch`、`delete`)；
* 避免多级路径，使用扁平结构，这样更灵活，方便管理和重构：
  * 错误：`GET /trades/32/orders/5/items/1`；
  * 正确: `GET /items/1?trade=32&order=5`；
* 模式定义：
  * `get` 查询、统计、列表等，无信息被修改；
  * `post` 创建、增加，主数据总量增长；
  * `put` 编辑、修改，主体信息被改变，总量不变但可能会变成另外一个实体；
  * `patch` 状态变更、执行命令，局部信息修改，主体未大改变被加工微调；
  * `delete` 删除、作废，失效不再关注；
* `get` 无主键是列表，有主键是根据主键查询单条记录
* `get` 有路径参数，是根据路径参数统计、分析；
* `patch` 主键写在路径上，参数放在`params`中：
  * `PATCH /trades/32/send?express=SF`
  * `PATCH /trades/32/confirm?mode=phone`
* `patch` 数据可以是结构化的

```http
PATCH /trades/32/tags
content-type: application/json

[1,3]
```
