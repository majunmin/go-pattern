# chain of responsibility

责任链模式


将业务处理使用链式的模式执行,当业务可以以流程的方式执行, 或者 根据上下文(context)信息 仅可以被一个handler 处理,可以使用  责任链模式.
责任链模式应用比较广泛, 

java中的
- spring中的 Filter 实现
- spring中的 Interceptor 实现
- alibaba sentinel 中的 流量过滤 Filter.


实现方式 一般有两种模式
一种是 一个 ChainManager 内部维护一个 数组/链表 来表示handler之间的顺序, 进行业务处理时可以使用   遍历数组/链表的方式
另一种是 一个 handler 内部维护一个 next(handler)，来进行链式调用.

## 应用场景
1. 敏感词过滤(sex, sensitive...)