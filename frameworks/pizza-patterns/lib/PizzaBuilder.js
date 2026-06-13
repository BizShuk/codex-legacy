const Pizza = require('./pizza/Pizza');

const DoughFactory = require('./DoughFactory');
const ToppingFactory = require('./ToppingFactory');
const SauceFactory = require('./SauceFactory');

class PizzaBuilder {

    constructor () {
        this.doughFactory = new DoughFactory();
        this.toppingFactory = new ToppingFactory();
        this.sauceFactory = new SauceFactory();
    }

    create (doughType, toppingType, sauceType) {
        let dough = this.doughFactory.create(doughType);
        let topping = this.toppingFactory.create(toppingType);
        let sauce = this.sauceFactory.create(sauceType);
        let pizza = new Pizza(dough, topping, sauce);

        return pizza;
    }

}


module.exports = PizzaBuilder;