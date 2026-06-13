module.exports = () => {
    var isUndefined, isNull=null
        ,isNumber=2, isZero=0, isNumberObject=new Number(3)
        , isString="abc", isEmptyString="", isStringObject=new String("")
        , isArray=[1,2,3], isEmptyArray=[]
        , isObject={"1":1}, isEmptyObject={}
        , isFunction=function(){}
        , isNotANumber=NaN;


    function parent(){};
    var parent1 = function(){};
    //// below can't be instanceof
    // var parent2 = () => {};
    // var parent3 = {};
    var parent_i = new parent();

    var son1 = {}; // false of instanceof parent
    son1.prototype = parent.prototype;

    function son_t(){}
    var son2 = new son_t(); // false of instanceof parent
    // son2.__proto__ = parent.prototype; // make son2 instanceof parent true
    // son2.__proto__ = parent_1;         // make son2 instanceof parent true
    son_t.prototype = new parent();
    var son3 = new son_t(); // true of instanceof parent


    return {
        'data':[
            isUndefined,
            isNull, isNotANumber,
            isNumber, isZero, isNumberObject,
            isString, isEmptyString, isStringObject,
            isArray, isEmptyArray,
            isObject, isEmptyObject,
            isFunction,
            son1,son2,son3
        ],
        'instance':[
            Object,
            Function,
            Array,
            Number,
            String,
            parent,
            son_t
        ]
    }
};