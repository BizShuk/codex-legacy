const BDough = require('../ingredients/dough/BDough');
const BTopping = require('../ingredients/topping/BTopping');
const BSauce = require('../ingredients/sauce/BSauce');

class BPizza {
    constructor () {
        this.dough = new BDough();
        this.topping = new BTopping();
        this.sauce = new BSauce();
    }
}


module.exports = BPizza;