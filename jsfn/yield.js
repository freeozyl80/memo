function* generatorFunction() {
  yield 'Hello';
  yield 'World';
  yield 'from';
  yield 'Node.js';
}

const generator = generatorFunction();

console.log(generator.next().value); // Output: Hello
console.log(generator.next().value); // Output: World
console.log(generator.next().value); // Output: from
console.log(generator.next().value); // Output: Node.js
console.log(generator.next().value); // Output: undefined