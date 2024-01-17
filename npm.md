# 好用的npm包

- [压测专用包](https://www.npmjs.com/package/autocannon)
- [refelect-metadata](https://jkchao.github.io/typescript-book-chinese/tips/metadata.html#%E5%9F%BA%E7%A1%80)
- [utils-decorators](https://jkchao.github.io/typescript-book-chinese/tips/metadata.html#%E5%9F%BA%E7%A1%80)
- [inversify: 近一步封装了对 依赖注入的实例化的过程](https://www.npmjs.com/package/inversify)
```javascript
import { injectable, inject, Container } from 'inversify';

// Defining an interface for the dependency
// which enforces the presence of 'doSomething()' method.
interface MyDependency {
  doSomething(): void;
}

// Implementing the interface with the @injectable decorator to 
// signify that it's an injectable class.
@injectable()
class MyDependencyImpl implements MyDependency {
  // Implementing the 'doSomething()'.
  doSomething(): void {
    console.log('Doing something...');
  }
}

// Creating a new instance of the Inversify container to manage dependency injection.
const container = new Container();

// Binding the MyDependency interface to the MyDependencyImpl class, 
// meaning when MyDependency is requested,
// the container will provide an instance of MyDependencyImpl.
container.bind<MyDependency>(MyDependency).to(MyDependencyImpl);

// Getting an instance of MyDependency from the container,
// which automatically resolves it to MyDependencyImpl
const instance = container.get<MyDependency>(MyDependency);

// Calling the 'doSomething()' method on the instance.
instance.doSomething();
```
- [nodemaile](https://www.npmjs.com/package/nodemailer)

- [concurrently 并行插件](https://www.npmjs.com/package/concurrently)

- [常用node的一些相关包](https://semaphoreci.medium.com/best-practices-for-securing-node-js-applications-in-production-d24b7c4981d)

- [node监控插件](https://clinicjs.org/doctor/)

- [请求频次限制](https://www.npmjs.com/package/rate-limiter-flexible)

- apm 库推荐
	- [signoz](https://github.com/SigNoz/signoz/blob/develop/README.zh-cn.md)
	- sentry
	- promtheus
	- relic
	- elastic
	- [logging, traces, metrices](https://www.atatus.com/blog/logging-traces-metrics-observability/)

- 