const ATopping = require('./ingredients/topping/ATopping');
const BTopping = require('./ingredients/topping/BTopping');
const CTopping = require('./ingredients/topping/CTopping');


class ToppingFactory {

    constructor () {

    }

    create (toppingType) {
        let topping;
        switch (toppingType) {
            case 'A': topping = new ATopping(); break;
            case 'B': topping = new BTopping(); break;
            case 'C': topping = new CTopping(); break;
            default:
                throw new Error('No topping Created');
        }
        return topping;
    }

}


module.exports = ToppingFactory;