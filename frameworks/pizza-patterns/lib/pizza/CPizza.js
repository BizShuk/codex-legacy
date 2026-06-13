const CDough = require('../ingredients/dough/CDough');
const CTopping = require('../ingredients/topping/CTopping');
const CSauce = require('../ingredients/sauce/CSauce');

class CPizza {
    constructor () {
        this.dough = new CDough();
        this.topping = new CTopping();
        this.sauce = new CSauce();
    }
}


module.exports = CPizza;