# 原型模式

如果一个对象的创建成本比较大, 而同一个类的不同对象之间差别不大(大部分字段都相同), 这种场景下,可以使用原型模式利用已有对象(原型)进行复制(拷贝)的方式创建新对象,以达到节省创建时间的目的. 
这种基于原型创建对象的方式叫做`原型模式`

- 对象的创建成本比较大?

创建对象的过程,实际上包含分配内存空间,给变量复制这一过程,这个过程本身不会花费太多时间,对于大多数场景,可以忽略.

如果对象的创建需要经过复杂的计算才能得到(排序,计算hash值),或者需要从 RPC, 网络, 数据库, 文件系统 等慢速IO设备中读取,这种场景下就可以使用原型模式。
从其他对象中拷贝得到,而不是在每次创建新对象的时候,都重复执行这些耗时操作.

## 深拷贝与浅拷贝
