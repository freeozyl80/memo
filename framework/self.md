# 微前端
	 micro-template 模板文件
	 micro-app   create Vue 实例， Router实例
	 micro-view  Server端渲染容器， Client端渲染容器
	 micro-node


# 构建

 - babel 按需加载

 	* @babel/core 核心，还在插件

 		1. @babel/parser 解释器
 		2. @babel/traverse 访问者模式，遍历ast节点
 		3. @babel/generator 转化源代码

 		// Note:babel配置中preset的配置的执行顺序是相反的


 		使用 entry 的方式，假使我们代码中仅使用到了 Array.prototype.includes 方法但是它会将目标浏览器中所有不兼容的 polyfill  全部实现并且引入进项目，造成包体积过大。 针对这种情况，Babel-Preset-Env useBuiltIns 提供了另一种配置方式:usage。 所谓 usage 即是表示它会分析我们的源代码，判断如果目标浏览器不支持并且我们代码中使用到的情况下才会在模块内进行引入对应的 polyfil。

 		无论是 entry 还是 usage 本质上都是通过注入浏览器不支持的 polyfill 在对应的全局对象上增加功能实现，这样无疑是会污染全局环境。


 		@babel/plugin-transform-runtime 编译过程中一些重复的工具函数变成外部模块引入的方


# 设计模式：

  多态： 

  继承：

  * 策略模式
  * 组合模式
  * 装饰者模式

  * 模板模式

# 设计原则：

  单一职责原则

  开闭原则： 对拓展开放，对修改关闭
  总结： Alert Class , check => tps, 
  		1. check => alertHandler  ; addALertHandler （tpsHandler, errorHanlder）
  		2. context 信息
 
  里氏替换:
  总结: 在上面的代码中，子类 SecurityTransporter 的设计完全符合里式替换原则，可以替换父类出现的任何位置，并且原来代码的逻辑行为不变且正确性也没有被破坏。


  接口隔离原则：

      组合方法


  依赖反转/控制反转/依赖注入

  1. 依赖反转： vue ssr asyncData的思想

  框架提供了一个可扩展的代码骨架，用来组装对象、管理整个执行流程。程序员利用框架进行开发的时候，只需要往预留的扩展点上，添加跟自己业务相关的代码，就可以利用框架来驱动整个程序流程的执行

  2. 依赖注入: 

  不通过 new() 的方式在类内部创建依赖类对象，而是将依赖的类对象在外部创建好之后，通过构造函数、函数参数等方式传递（或注入）给类使用

  3. 依赖注入框架： context


  依赖反转原则：

	高层模块（high-level modules）不要依赖低层模块（low-level）。高层模块和低层模块应该通过抽象（abstractions）来互相依赖。除此之外，抽象（abstractions）不要依赖具体实现细节（details），具体实现细节（details）依赖抽象（abstractions）


 4. 高内聚，松耦合

# 短链接服务

  1. shortUrl, longUrl 插入，需要判断有则取，无则建
  2. 多台 redis 生成不同 号码类型的 自增 id 号码



# 手写几个排序















# ci /cd 
# node
# node ssr： react ssr

# vue 和 react ， pull 和 push 的区别

# golang 实现电商网站，电商网站的一些东西

gmp: https://learnku.com/articles/41728

# 短链接服务
# 设计模式
