const ADough = require('./ingredients/dough/ADough');
const BDough = require('./ingredients/dough/BDough');
const CDough = require('./ingredients/dough/CDough');


class DoughFactory {

    constructor () {

    }

    create (doughType) {
        let dough;
        switch (doughType) {
            case 'A': dough = new ADough(); break;
            case 'B': dough = new BDough(); break;
            case 'C': dough = new CDough(); break;
            default:
                throw new Error('No dough Created');
        }
        return dough;
    }

}


module.exports = DoughFactory;