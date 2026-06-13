const ADough = require('../ingredients/dough/ADough');
const ATopping = require('../ingredients/topping/ATopping');
const ASauce = require('../ingredients/sauce/ASauce');

class APizza {
    constructor () {
        this.dough = new ADough();
        this.topping = new ATopping();
        this.sauce = new ASauce();
    }
}


module.exports = APizza;