## 面试题模板（由于换电脑了，之前的面试题模板没了 T.T）

### React面试题

1. SSR端
 * 讲一讲ssr的过程
 * node端如何发起请求
 * 

2. 优化端
 * 性能优化端：
 	https://v2ex.com/t/724286#reply26
 	useRef: 存储某些状态而不需要重新渲染来达成，比如button的 disable 的属性
  createRef()
 * react-loadable 分割代码

 

3. react 事件

4. react  高阶组件的case设计：
  * 弹窗的实现（）： 有多个弹窗，维护多个弹窗的状态（但是弹窗是不会共存的）
  * abTest 是否显示组件
  * 停留时长关闭弹窗

 5. react 弹窗实现：
  
  * 快捷调用问题： class 添加 static 方法
  * 实现一个弹窗组件时每次都要引入 Modal，都要重写一次 show 方法: 高阶组件
  参考文章(https://kreja.github.io/study/2018/08/20/modal.html)
  
6. react 如何获取 改变value 的 history

7. react 为什么不 在 useEffect 里面有异步操作？

  useEffect 作为 Hooks 中一个很重要的 Hooks，可以让你在函数组件中执行副作用操作。 它能够完成之前 Class Component 中的生命周期的职责。它返回的函数的执行时机如下

  （首次渲染不会进行清理，会在下一次渲染，清除上一次的副作用。卸载阶段也会执行清除操作。）

   参考文章(https://www.jianshu.com/p/e1ff909aca88)

   useEffect 依赖 时object.is比较

   useEffect 中断异步操作方法： 1. return ajax abort; 2.return lock值

### 前端可视化面试题

canvas , svg , webgl ：

	共同点：都是有效的图形工具，对于数据较小的情况下，都很又高的性能，它们都使用 JavaScript 和 HTML；它们都遵守万维网联合会 (W3C) 标准。

	svg优点： 矢量图，不依赖于像素，无限放大后不会失真。 以dom的形式表示，事件绑定由浏览器直接分发到节点上。 svg缺点： dom形式，涉及到动画时候需要更新dom，性能较低。

	canvas优点： 定制型更强，可以绘制绘制自己想要的东西。 非dom结构形式，用JavaScript进行绘制，涉及到动画性能较高。

	canvas缺点： 事件分发由canvas处理，绘制的内容的事件需要自己做处理。 依赖于像素，无法高效保真，画布较大时候性能较低

canvas渲染较大画布的时候性能会较低？为什么？
因为canvas依赖于像素，在绘制过程中是一个一个像素去绘制的，当画布足够大，像素点也就会足够多，那么想能就会足够低。

假设我现在有5000个圆，完全绘制出来，点击某一个圆，该圆高亮，另外4999个圆设为半透明，分别说说用svg和canvas怎么实现？

### 几个有争议的问题

vue 弹窗的实现

v-if v-show

tab 切换，造成数据不一致


大数据分页的逻辑，筛选后页面没有重置 ？

react useref 是干嘛的

```code 
# 定时器（观察组件存活时间，但是组件重新渲染不重置，但是必须是展示出来才开始计时）
const App = () => {
  const timer = useRef();

  useEffect(() => {
    timer.current = setInterval(() => {
      console.log('触发了');
    }, 1000);
  },[]);

  const clearTimer = () => {
    clearInterval(timer.current);
  }

  return (
    <>
      <Button onClick={clearTimer}>停止</Button>
    </>)
}

```

### 罗列几个问题
```javascript
function Counter() {
  const [count, setCount] = React.useState(0);

  function increment() {
    setCount((n) => n + 1);
  }
  return <ChildComponent count={count} onClick={increment} />;
}
// usecallback 的解决场景，其实hooks写法会重新渲染，但是 class写法 会么 ？
// props 或 state 更新的时候， static getDerivedStateFromProps() > shouldComponentUpdate() > render() > getSnapshotBeforeUpdate() > componentDidUpdate()
```
* usememo
```javascript
const expensive = useMemo(() => {
  // 昂贵计算
}, [count])

const callback = useCallback(() => {
  // 昂贵计算
}, [count])
```


### 高阶函数：(https://juejin.cn/post/6844904050236850184)
1.  操作props，条件渲染 ？ 比如暂无数据的视图
2.  反向劫持，通过extends ? 

案例： 不同的数据源，大致相同的列表渲染
案例： 白名单，ABTEST
案例： 性能的追踪