# react 相关指南

- react 一切皆组件（函数）

- 声明周期
  
  * getDerivedStateFromProps
  * shouldComponentUpdate
  * render
  * getSnapshotBeforeUpdate
  * componentDidUpdate
  * componentWillUnmount
  * getSnapshotBeforeUpdate

  有关shouldComponentUpdate
  Purememo整个组件会重新渲染，如果不适用shouldComponentUpdate那么Title，Count会一直渲染，但是使用shouldComponentUpdate后Title则不会发生渲染。


- react 副作用
  [参考](https://juejin.cn/post/7175843021752598589)


  第1次：render（渲染） -> useEffect
  第2次：state（状态变化）-> render（渲染）-> 清理函数 -> useEffect
  第n次：state（状态变化）-> render（渲染）-> 清理函数 -> useEffect


- setState 执行机制：
  1. 异步执行，不能立马返回值，需要在 setState(xxx, () => '返回值') ， 合并返回最后一个
  2. 同步执行， setTimeout 中，可以立马获得值 ， 不会合并

- react 事件流程
  1. React 所有事件都挂载在 document 对象上
  2. 当真实 DOM 元素触发事件，会冒泡到 document 对象后，再处理 React 事件
  3. 所以会先执行原生事件，然后处理 React 事件
  4. 最后真正执行 document 上挂载的事件

- react 组件
  1. 无状态组件

- context:
```javascript
const PriceContext = React.createContext('price')

<PriceContext.Provider value={100}>
</PriceContext.Provider>

<PriceContext.Consumer>
    { /*这里是一个函数*/ }
    {
        price => <div>price：{price}</div>
    }
</PriceContext.Consumer>

```
- react keys

  同样，并不是拥有key值代表性能越高，如果说只是文本内容改变了，不写key反而性能和效率更高

  主要是因为不写key是将所有的文本内容替换一下，节点不会发生变化

  而写key则涉及到了节点的增和删，发现旧key不存在了，则将其删除，新key在之前没有，则插入，这就增加性能的开销

- react 高阶组件

  如果向一个高阶组件添加refe引用，那么ref 指向的是最外层容器组件实例的，而不是被包裹的组件，如果需要传递refs的话，则使用React.forwardRef，如下

  1. 正向属性代理：

    我们可以理解为父子组件关系，父组件对子组件进行一系列强化操作。

  2. 反向继承：

   方便获取组件内部状态，比如state，props ,生命周期,绑定的事件函数等
   es6继承可以良好继承静态属性。我们无须对静态属性和方法进行额外的处理。

  [推荐文章](https://juejin.cn/post/6940422320427106335)


 - setState 和 useState:

   setState是浅合并

 - react css 引入
   webpack xxx.module.css 可以实现 scoped

 - css 动画：


  1. CSSTransition：在前端开发中，结合 CSS 来完成过渡动画效果
  2. SwitchTransition：两个组件显示和隐藏切换时，使用该组件
  3. TransitionGroup：将多个动画组件包裹在其中，一般用于列表中元素的动画


- redux

 1. redux: 

   * 生成 reducer
   * createStore

   * Provider 顶层
   * connect(mapStateToProps, mapDispatchToProps)(MyComponent) connect方法将store上的getState和 dispatch包装成组件的props


- immuntable 不可以变的
  setState 可以不是一个整体     ============================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


- react 渲染效率
  类组件: shouldComponentUpdate自比较， pureComponent浅比较

  数据更新：尽量避免愿数据改变的操作

- react diff

 1. tree 层级： 只增/删
 2. components 层级： 同一个类组件，则会往下比较

- react fiber 架构

  在react中，主要做了以下的操作：

  为每个增加了优先级，优先级高的任务可以中断低优先级的任务。然后再重新，注意是重新执行优先级低的任务
  增加了异步任务，调用requestIdleCallback api，浏览器空闲的时候执行
  dom diff树变成了链表，一个dom对应两个fiber（一个链表），对应两个队列，这都是为找到被中断的任务，重新执行

   一个 fiber就是一个 JavaScript对象，包含了元素的信息、该元素的更新操作队列、类型，


  该实现过程是基于 Fiber节点实现，作为静态的数据结构来说，每个 Fiber 节点对应一个 React element，保存了该组件的类型（函数组件/类组件/原生组件等等）、对应的 DOM 节点等信息。

  作为动态的工作单元来说，每个 Fiber 节点保存了本次更新中该组件改变的状态、要执行的工作。

  每个 Fiber 节点有个对应的 React element，多个 Fiber节点根据如下三个属性构建一颗树：


- react 编译过程:

  1. babel 
    
  当首字母为小写时，其被认定为原生 DOM 标签，createElement 的第一个变量被编译为字符串

  当首字母为大写时，其被认定为自定义组件，createElement 的第一个变量被编译为对象

  2. type

  原生标签节点
  文本节点
  函数组件
  类组件

  3. react 性能优化方式：

  - 避免使用内联动函数，每次新增空间
  使用 React Fragments 避免额外标记

  使用 Immutable

  懒加载组件

  事件绑定方式

  服务端渲染

- useReducer 可以管理多个依赖



## react 性能优化高阶篇
```javascript
// 场景一
function Box() {
  const [position, setPosition] = useState({ left: 0, top: 0 });
  const [size, setSize] = useState({ width: 100, height: 100 });
  // 重复渲染次数变多, 异步的环境，这里有误
}
// 场景二
function Counter() {
  const [count, setCount] = React.useState(0);

  function increment() {
    setCount((n) => n + 1);
  }
  return <ChildComponent count={count} onClick={increment} />;
}
// 在 count 状态更新的时候， Counter 组件会重新执行，这个时候会重新创建一个新的函数 increment。这样传递给 ChildComponent 的 onClick 每次都是一个新的函数，从而导致 ChildComponent 组件的 React.memo 失效

// 解决办法：
function usePersistFn<T extends (...args: any[]) => any>(fn: T) {
  const ref = React.useRef<Function>(() => {
    throw new Error('Cannot call function while rendering.');
  });
  ref.current = fn;
  return React.useCallback(ref.current as T, [ref]);
}
const increment = usePersistFn(() => {
  setCount((n) => n + 1);
});
// 或者使用 useCallback
const increment = React.useCallback(() => {
  setCount((n) => n + 1);
}, []);

```

```javascript
function useWhyDidYouUpdate(name, props) {
  const previousProps = useRef();

  useEffect(() => {
    if (previousProps.current) {
      const allKeys = Object.keys({ ...previousProps.current, ...props });
      const changesObj = {};
      allKeys.forEach(key => {
        if (previousProps.current[key] !== props[key]) {
          changesObj[key] = {
            from: previousProps.current[key],
            to: props[key]
          };
        }
      });

      if (Object.keys(changesObj).length) {
        console.log('[why-did-you-update]', name, changesObj);
      }
    }
    previousProps.current = props;
  });
}

const Counter = React.memo(props => {
  useWhyDidYouUpdate('Counter', props);
  return <div style={props.style}>{props.count}</div>;
});

```

- 使用React.lazy与React.Suspense后出现的问题 实现动态加载


- react 弹窗的实现
  1. 正常模式，实现一个 类弹窗，通过 visible state 进行控制
  评论： 实现一个弹窗组件时每次都要引入 Modal，都要重写一次 show 方法


  2. ReactDOM.render 动态加载， 通过装饰器，用 高阶组件，


  * 快捷调用问题： class 添加 static 方法
  * 实现一个弹窗组件时每次都要引入 Modal，都要重写一次 show 方法: 高阶组件
  参考文章(https://kreja.github.io/study/2018/08/20/modal.html)


- react 函数组件 vs 类组件经典问题

  [链接](https://juejin.cn/post/7033750813986324510?utm_source=gold_browser_extension)

```jsx

class CompClass extends Component {

 showMessage = () => {
  console.log("点击的这一刻，props中info为 " + this.props.info);
 };

 handleClick = () => {
   setTimeout(this.showMessage, 3000);
   console.log(`当前props中的info为${this.props.info},一致就说明准确的关联到了此时的render结果`)
 };

 render() {
  return <div onClick={this.handleClick}>
     <div>点击类组件</div>
  </div>;
 }
}

function CompFunction(props) {
  
  const showMessage = () => {
    console.log("点击的这一刻，props中info为 " + props.info);
  };
 
  const handleClick = () => {
    setTimeout(showMessage, 3000);
    console.log(`当前props中的info为${props.info},一致就说明准确的关联到了此时的    render结果`)
  };

  return <div onClick={handleClick}>点击函数组件</div>;
}


export default function App() {
 const [info, setInfo] = useState(0);
 return (
   <div>
     <div onClick={()=>{
      setInfo(info+1)
     }}>父组件的info信息>> {info}</div>
     <CompFunction info = {info}></CompFunction>
     <CompClass info = {info}></CompClass>
   </div>
 );
}

// 函数组件CompFunction会输出：0
// 类组件CompClass会输出：5

// 函数组件执行，就会形成一个闭包，可以形象地说成render结果，其中包括props，而点击事件的处理函数同样也包括在内，那它无论是立即执行还是延迟执行，都应该与触发执行的那一刻的render结果（你也可以理解为那一刻的快照）相关联。
//所以回调函数showMessage所应该log出的info，应该为事件触发的那一刻render结果中的info，也就是"1"，无论外部的info怎么变。
// 而类组件就会输出info的最新值，也就是"5"。

/*
  每次 Render 的内容都会形成一个快照并保留下来，因此当状态变更而 Rerender 时，就形成了 N 个 Render 状态，而每个 Render 状态都拥有自己固定不变的 Props 与 State。
*/

/*
我们之前聊了，函数组件最大的硬伤就是"次次重来，无法延续" ，很难让它具备跟类组件那样的能力，比如用状态和生命周期函数，而如今hooks的加持，很好的粉碎了被类组件克制的枷锁。
*/

```

  引申 fiber问题:
  reconciler: 
    1. reconciliation: 通过diff获得变动的结果。
    2. commit: 将变动作用到画面上（side effect即副作用，如dom操作）。

  Fiber就是按照vdom来拆分的，一个vdom节点对应一个Fiber节点，最后形成一个链表结构的fiber tree，

  1) 自顶root向下，流转子节点b1
  2)b1开始beginWork，工作目标根据情况diff处理，获得变动结果（effectList），然后判断是是否有子节点，没有那结束工作completeWork，然后流转到兄弟节点b2
  3) b2开始工作，然后判断有子节点c1，那就流转到c1
  4) c1工作完了，completeWork获得effectList，并提交给b2
  5) 然后b2完成工作，流转给b3,那么b3就按照这套路子，往下执行了，最后执行到了最底部d2
  6) 最后随着光标的路线，一路整合各节点的effectList，最后抵达Root节点，第1阶段-reconciliation结束，准备进入Commit阶段

  双循环树

  1) 根据current fiber treeclone出workinProgress fiber tree，每clone一个workinProgress fiber都会尽可能的复用备用fiber节点（曾经的current fiber）
  2) 当构建完整个workinProgress fiber tree的时候，current fiber tree就会退下去，作为备用fiber节点树，然后workinProgress fiber tree就会扶正，成为新的current fiber tree
  3) 然后就将已收集完变动结果（effect list）的新current fiber tree，送去commit阶段，从而更新画面

  alternate 指针关联两颗树

  current workingProgress 的hooks 的存储结构 决定了 不能 在hooks 里面写条件语句


   
回到上面的问题，可以用useref解决


继续

```jsx
let A = ((props) => {
  const [count,setCount] = useState({a:1})
  
  return <div onClick={()=>{
    count.a = Date.now();
    setCount(count)
  }}>
    测试是否刷新
  </div>
})
// 不刷新
React内部会针对传入的参数进行浅比较，引用类型的数据比较的是其指向的地址，而不是内容，切记，所以光内容变了没用。
// 再谈setState
// 1. 第二个参数是一个函数，可以在状态值设置生效后进行回调，我们就可以在这里面拿到最新的状态值。
// 2. setState具备浅合并功能，比如state是{a:1,b:2,c:{e:0}},setState({c:{f:0},d:4}),state就会合并成{a:1,b:2,c:{f:0},d:4}
// 3. 一定会引发刷新

//set函数

// 没有第二个参数，但是可以借助useEffect组合实现,也还好
// 没有合并功能，设置啥就是啥。。。,不过自己动手优化一下也是可以的。
// 设置相同的状态是不会触发刷新的，这一点无需进行配置。
// 也是异步的，可以考虑函数的方式来更新状态
```

越来越经典
[useImperativeHandle](https://zh-hans.react.dev/reference/react/useImperativeHandle#exposing-a-custom-ref-handle-to-the-parent-component)

用来 解决 forwardRef 会把所有暴露给父组件的情况下，只暴露少量句柄
```jsx
// 除此之外还有一个有趣的用处，就是可以结合forwardRef和useImperativeHandle这两个api，让函数组件可以像类组件那样暴露函数给其他节点使用
let A = forwardRef((props, ref) => {
  (ref, () => {
    return {
      test: () => {
        console.log("123")
      }
    }
  })
  const [info, setInfo] = useState(0);
  return <div>
    {info}
  </div>
})

export default function App() {
  const ref = useRef(null);
  return (
    <div onClick={() => {
      ref.current.test()
    }}>
      <A ref={ref}>
      </A>
    </div>
  );
}


```

有关 useReducer 和 context 实现 状态管理

```jsx
import {
  createContext,
  useContext,
  useReducer,
} from "react";
export const TestContext = createContext({})
const TAG_1 = 'TAG_1'

const reducer = (state, action) => {
  const { payload, type } = action;
  switch (type) {
    case TAG_1:
      return { ...state, ...payload };
      dedault: return state;
  }
};

export const A = (props) => {
  const [data, dispatch] = useReducer(reducer, { info: "本文作者" });
  return (
    <TestContext.Provider value={{ data, dispatch }}>
      <B></B>
    </TestContext.Provider>
  );
};

let B = () => {
  const { dispatch, data } = useContext(TestContext);
  let handleClick = ()=>{
    dispatch({
        type: TAG_1,
        payload: {
          info: "闲D阿强",
        },
      })
  }
  return (
    <div>
      <input
        type="button"
        value="测试context"
        onClick={handleClick}
      />
      {data.info}
    </div>
  );
};

```


- react 编译过程的经典示例（https://mp.weixin.qq.com/s/zhaVGY1yrQB1fECiQMntaQ）


  * vdom 转 fiber 的流程叫做 reconcile，我们常说的 diff 算法就是在 reconcile 这个过程中。

  * hooks 挂载：

   就是第一个节点挂在 fiber 节点的 memoizedState 属性上，后面的挂在上个节点的 next 属性上。

   总结：

    react 渲染流程分为 render 和 commit 阶段。

    render 阶段执行 vdom 转 fiber 的 reconcile，commit 阶段更新 dom，执行 effect 等副作用逻辑。

    commit 阶段分为 before mutation、mutation、layout 3 个小阶段。

    hook 的数据就是保存在 fiber.memoizedState 的链表上的，每个 hook 对应一个链表节点。

    hook 的执行分为 mountXxx 和 updateXxx 两个阶段，第一次会走 mountXxx，创建 hook 链表，之后执行 updateXxx。

    我们看了 useRef、useMemo、useCallback 的实现原理，这几个 hook 都比较简单。其中后两个 hook 是作为 props 时为了减少不必要渲染的时候用的。

    useState 和 useEffect 就和渲染流程有关了：

    useEffect 在 render 阶段会把 effect 放到 fiber.updateQueue 的环形链表上，然后在 commit 阶段遍历所有 fiber 的 updateQueue，取出 effect 异步执行。

    useLayoutEffect 和 useEffect 差不多，只是 effect 链表是在 layout 阶段同步执行的。

    useState 的 mountState 阶段返回的 setXxx 是绑定了几个参数的 dispatch 函数。执行它会创建  hook.queue 记录更新，然后标记从当前到根节点的 fiber 的 lanes 和 childLanes 需要更新，然后调度下次渲染。

    下次渲染执行到 updateState 阶段会取出 hook.queue，根据优先级确定最终的 state，最后返回来渲染。

    这样就实现了 state 的更新和重新渲染。

    这就是 react hooks 特别是 useState 和 useEffect 的实现原理。理解它们不单单要理解 hook 的存储结构，还要理解 react 的整个渲染流程。


```jsx
const WatchCount = () => {
  const [count, setCount] = useState(0);
  
  useEffect(() => {
    setInterval(function log() {
      console.log(`Count: ${count}`);
    }, 2000);
  }, []);
  
  const handleClick = () => setCount(count => count + 1);
  
  return (
    <>
      <button onClick={handleClick}>+</button>
      <div>Count: {count}</div>
    </>
  );
}

useState 是一个闭包，之前的函数组件也是一个必包
```


