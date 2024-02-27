import React, {Component} from 'react'

class B extends React.Component {
  constructor(props){
    super(props)
    // this.changeMsg = this.changeMsg.bind(this);
    this.state = {
      msg: 'hi' + Date.now()
    }
  }
  static getDerivedStateFromProps() {
    console.log('初始状态，获取props', arguments)
  }
  shouldComponentUpdate() {
    console.log('是否需要更新', arguments)
    return true
  }
  changeMsg = () => {
    this.setState({
      msg: 'hi' + Date.now()
    })
    this.setState({
      msg: 'h ～ i' + Date.now()
    })
    this.setState((prev) => ({
      msg: prev.msg + '..',
    }));
  }
  render(){
    return(
      <div className="demo-text" ref={this.refIns} onClick={this.changeMsg}>{this.state.msg}</div>
    )
  }
  getSnapshotBeforeUpdate() { 
    console.log('#enter getSnapshotBeforeUpdate');
    return 'foo';
  }
  componentDidMount() {
    this.refIns = React.createRef();
    console.log('componentDidMount', arguments)

  }
  componentDidUpdate() {
    console.log(this.refIns)
    console.log('componentDidUpdate', arguments)
  }
}

const HocComponent = function(Component) {
  class Enhance extends React.Component{
    constructor(){
      super()
    }
    componentDidMount() {
      console.log('HOC组件更新 ====================================')
      console.log(this.myref)
    }
    render() {
        const {myref, ...rest} = this.props;
        // 把 forwardedRef 赋值给 ref
        return (
          <>
            <Component {...rest} ref={myref}/>
          </>
      );
    }
  }
  // React.forwardRef 方法会传入 props 和 ref 两个参数给其回调函数
  // 所以这边的 ref 是由 React.forwardRef 提供的
  function forwardRef(props, ref) {
    console.log(ref, 123)
    return <Enhance {...props} myref={ref} />
  }
  return React.forwardRef(forwardRef);
}

const HocComponent2 = function(Component) {
  return class Enhance extends Component {

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

export default CompFunction