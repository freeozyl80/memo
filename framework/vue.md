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



- vue3 的提升

  可以看到，组件内部只有一个动态节点，剩余一堆都是静态节点，所以这里很多 diff 和遍历其实都是不需要的，造成性能浪费，因此，Vue3在编译阶段，做了进一步优化。主要有如下

  * diff算法优化： 静态标记
  * 静态提升： createVnode 的时候，已经声明了
  * 事件监听缓存
  * SSR优化

  : 当静态内容大到一定量级时候，会用createStaticVNode方法在客户端去生成一个static node，这些静态node，会被直接innerHtml，就不需要创建对象，然后根据对象渲染

  vue2中采用 defineProperty来劫持整个对象，然后进行深度遍历所有属性，给每个属性添加getter和setter，实现响应式

  vue3采用proxy重写了响应式系统，因为proxy可以对整个对象进行监听，所以不需要深度遍历

  可以监听动态属性的添加
  可以监听到数组的索引和数组length属性
  可以监听删除属性

  defineProperity: 

  1. 遍历
  2. 递归
  3. 数组 X
  4. 对象属性 add / del

- compositon api 的

  函数化

- api 组件：

  vue2:  Vue.extend()
  vue3: createVnode








- vue context 的概念，对比 react

- vue 实现的 api 弹窗，以及 vue2 和 vue3 的区别

- vue template 是否需要编译

- proxy 的使用

- vue 绑定与监听的问题， vue data属性做了什么，为什么


https://vue3js.cn/interview/vue/data.html