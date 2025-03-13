function debounce(func, wait) {
  let timeout;
  return function(...args) {
    const context = this;
    clearTimeout(timeout);
    timeout = setTimeout(() => func.apply(context, args), wait);
  };
}

// 示例用法
const myDebouncedFunction = debounce(() => {
  console.log('Function executed!');
}, 300);

// 调用防抖函数
myDebouncedFunction();
myDebouncedFunction();
myDebouncedFunction();
// 只有最后一次调用会在300毫秒后执行

// ----------------- 深拷贝--------------------

function deepClone(obj) {
  if (obj === null || typeof obj !== 'object') {
    return obj;
  }

  if (obj instanceof Date) {
    return new Date(obj);
  }

  if (obj instanceof Array) {
    return obj.map(item => deepClone(item));
  }

  if (obj instanceof Object) {
    const copy = {};
    for (const key in obj) {
      if (obj.hasOwnProperty(key)) {
        copy[key] = deepClone(obj[key]);
      }
    }
    return copy;
  }

  throw new Error('Unable to copy object!');
}

// 示例用法
const original = {
  name: 'Japhy',
  details: {
    age: 30,
    hobbies: ['coding', 'music']
  },
  date: new Date()
};

const copied = deepClone(original);
console.log(copied);
console.log(copied.details === original.details); // false
console.log(copied.date === original.date); // false
// ----------------- for 循环 闭包问题--------------------

for (var i = 0; i < 3; i++) {
  setTimeout(function() {
    console.log(i);
  }, 1000);
}

// 解决方法：使用 let 替代 var
for (let i = 0; i < 3; i++) {
  setTimeout(function() {
    console.log(i);
  }, 1000);
}

// 解决方法：使用 IIFE (立即执行函数表达式)
for (var i = 0; i < 3; i++) {
  (function(i) {
    setTimeout(function() {
      console.log(i);
    }, 1000);
  })(i);
}