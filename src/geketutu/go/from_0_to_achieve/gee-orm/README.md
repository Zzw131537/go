### ORM (对象关系映射)

将面向对象语言程序中的对象自动持久化到数据库中

如何根据任意类型的指针，得到其对应的结构体的信息。这涉及到了 Go 语言的反射机制(reflect)，通过反射，可以获取到对象对应的结构体名称，成员变量、方法等信息，

#### day1-sql数据库
 
sqlite 的一些用法

实现一个简单的log库，具备一下特性

1. 支持日志分级
2. 不同层级的日志显示时使用不同的颜色区分
3. 显示打印日志代码对用的文件名和行号
   

#### day2-对象表结构的映射

1. 使用dialect 隔离不同数据库之间的差异，便于扩展
2. 使用反射reflect获取任意struct 对象的名称和字段
3. 数据库的创建与删除
   
Schema 对象和表的转换
