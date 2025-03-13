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
# 接口和类型有什么区别

- 接口
  定义一个类应该包含哪些方法

- 类型


```typescript
type Job = {
    time: Number
}


interface Working {
    name: string;
    greet: (greeting: string) => string;
}
class Employee implements Working {
    name: string;
    job: Job;

    constructor(name: string, job: Job) {
        this.name = name;
        this.job = job;
    }

    greet(greeting: string): string {
        return `${greeting}, my name is ${this.name} and I work for ${this.job.time} hours.`;
    }
}

const job: Job = { time: 8 };
const employee = new Employee('Alice', job);
console.log(employee.greet('Hello'));

```
