const PizzaFactory = require('./lib/PizzaFactory');
const PizzaBuilder = require('./lib/PizzaBuilder');

const pizzaFactory = new PizzaFactory();


let Apizza = pizzaFactory.createPizza('A');
let Bpizza = pizzaFactory.createPizza('B');
let Cpizza = pizzaFactory.createPizza('C');


console.log(Apizza);
console.log(Bpizza);
console.log(Cpizza);


const pizzaBuilder = new PizzaBuilder();

let C1pizza = pizzaBuilder.create('A','B','C');
let C2pizza = pizzaBuilder.create('B','B','C');
let C3pizza = pizzaBuilder.create('C','B','C');
let C4pizza = pizzaBuilder.create('A','B','C');

