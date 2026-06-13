const ASauce = require('./ingredients/sauce/ASauce');
const BSauce = require('./ingredients/sauce/BSauce');
const CSauce = require('./ingredients/sauce/CSauce');


class SauceFactory {

    constructor () {

    }

    create (SauceType) {
        let sauce;
        switch (SauceType) {
            case 'A': sauce = new ASauce(); break;
            case 'B': sauce = new BSauce(); break;
            case 'C': sauce = new CSauce(); break;
            default:
                throw new Error('No Sauce Created');
        }
        return sauce;
    }

}


module.exports = SauceFactory;