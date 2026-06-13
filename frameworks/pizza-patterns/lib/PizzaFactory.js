const APizza = require('./pizza/APizza');
const BPizza = require('./pizza/BPizza');
const CPizza = require('./pizza/CPizza');

class PizzaFactory {

    constructor () {
    }

    createPizza (pizzaType) {
        let pizza;
        switch (pizzaType) {
            case 'A': pizza = new APizza(); break;
            case 'B': pizza = new BPizza(); break;
            case 'C': pizza = new CPizza(); break;
            default:
                throw new Error('No Pizza Created');
        }
        return pizza;
    }
}


module.exports = PizzaFactory;