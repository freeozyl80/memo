# vue 问题罗列

- vue 编译流程

* 将模板template转为ast结构的JS对象
* 用ast得到的JS对象拼装render和staticRenderFns函数
* render和staticRenderFns函数被调用后生成虚拟VNODE节点，该节点包含创建DOM节点所需信息
* vm.patch函数通过虚拟DOM算法利用VNODE节点创建真实DOM节点

- vue3 和 vue2 的区别

- v-if 和 v-show 的区别

* 编译过程：v-if切换有一个局部编译/卸载的过程，切换过程中合适地销毁和重建内部的事件监听和子组件；v-show只是简单的基于css切换
* v-for优先级是比v-if高


- Vue 不允许在已经创建的实例上动态添加新的响应式属性, defineProperity 为对新对象添加， 可以使用 vue.set


- vue 双向绑定的原理

	* new Vue()首先执行初始化，对data执行响应化处理，这个过程发生Observe中

	* 同时对模板执行编译，找到其中动态绑定的数据，从data中获取并初始化视图，这个过程发生在Compile中

	* 由于data的某个key在⼀个视图中可能出现多次，所以每个key都需要⼀个管家Dep来管理多个Watcher

	* 将来data中数据⼀旦发生变化，会首先找到对应的Dep，通知所有Watcher执行更新函数



- vue context 的概念，对比 react

- vue 实现的 api 弹窗，以及 vue2 和 vue3 的区别

- vue template 是否需要编译

- proxy 的使用

- vue 绑定与监听的问题， vue data属性做了什么，为什么


https://vue3js.cn/interview/vue/data.html