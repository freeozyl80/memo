# 微前端
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


# ci /cd 
# node
# node ssr

# vue 和 react ， pull 和 push 的区别

# golang 实现电商网站，电商网站的一些东西

# 短链接服务
# 设计模式
