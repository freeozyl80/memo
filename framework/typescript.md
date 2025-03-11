interface vs type

```typescript
interface User {
  name: string
  age: number
}

interface SetUser {
  (name: string, age: number): void;
}

type User = {
  name: string
  age: number
};

type SetUser = (name: string, age: number)=> void;

```


区别：

1. type 可以声明基本类型别名，联合类型，元组等类型

interface 是方法集，type 是数据属性结构