/* Apriori algo : https://en.wikipedia.org/wiki/Association_rule_learning#Apriori_algorithm
 * 
 * Find relation between items in all transactions , more higher frequency more related.
 *
 * constructor: 
 *      data = transaction list = [transaction 1, transaction 2,...] , transaction = [item,item,...]
 * 
 * usage:
 *      1. get instance by `var ins = apriori(data)`
 *      2. calculate with `ins.calc(miniFreq)` , miniFreq is a number for how frequenc you want to filter with item combination(item_set).
 *         Each time trigger this method will get bigger item_set with frequency
 *      3. last one will return some combination with frequencies.
 *         e.g.
 *         [ { item_set: [ 'p1', 'p3', 'p2' ], freq: 1 },
 *           { item_set: [ 'p1', 'p3', 'p5' ], freq: 1 },
 *           { item_set: [ 'p1', 'p2', 'p5' ], freq: 1 },
 *           { item_set: [ 'p3', 'p2', 'p5' ], freq: 2 } ]
 */
var apriori = function(data) {

    var list = data, // transactions list , each one has some items
        items = {}, // { "item":[tid,tid] } , item in which transactions
        miniFreq = 1, // minimum frequency for item_set
        combine_level = 1, // start from 1 , because item_set has one item at least.
        cur_associated_result = []; // current assocaited result ,  [ { item_set:[prop1,prop2],freq:number },... ]

    if (!(list instanceof Array)) {
        return null;
    }

    // initialize paremeters
    var reset = function() {
        miniFreq = 1;
        combine_level = 1;
        cur_associated_result = [];

        // initial cur_associated_result
        for (var key in items) {
            cur_associated_result.push({
                item_set: [key],
                freq: items[key].length
            });
        }
    }

    // Initialization
    list.forEach(function(t, tid, a) {
        if (t instanceof Array) {
            t.forEach(function(p, i2, a2) {
                if (typeof p == "string") {
                    items[p] ? items[p].push(tid) : items[p] = [tid];
                }
            })
        }
    });
    reset();


    // manual output
    var show = function() {
        console.log(list);
        console.log(items);
        console.log(cur_associated_result);
    }


    // make combination of items 
    // e.g.
    //      {1,2,3},1   =>  [[1],[2],[3]]
    //      {1,2,3},2   =>  [[1,2],[2,3],[1,3]]
    //      {1,2,3},3   =>  [[1,2,3]]
    //
    var combine_items = function(_set, count) {
        var ary = [..._set];
        var _item_set = [];
        var _all_item_set = [];

        if (ary.length < count || count <= 0) {
            return [];
        }

        // transfer index which store in array value to real value
        function transfer_index_to_value(_item_set) {
            var _after = [];
            _item_set.forEach(function(e, i, a) {
                _after.push(ary[e]);
            });
            return _after;
        }

        var i = 0,
            j = i;

        while (i + count <= ary.length) {
            if (j >= ary.length) {
                j = _item_set.pop();
                j++;
                if (_item_set.length == 0) {
                    i = j;
                }
                continue;
            }

            if (_item_set.length < count) {
                _item_set.push(j);
                j++;
            }

            if (_item_set.length >= count) {
                _all_item_set.push(transfer_index_to_value(_item_set));
                j = _item_set.pop() + 1;
            }

        }

        return _all_item_set;
    }

    // get dintinct items from current associated result
    function get_distinct_items(_ms) {
        var _dict_property = new Set();

        cur_associated_result.forEach(function(e, i, a) {
            if (e.freq >= _ms) {
                e.item_set.forEach(function(e2, i2, a2) {
                    _dict_property.add(e2);
                });
            }
        });
        return _dict_property;
    }

    // Calculate next combine level associated result
    //
    // input: _ms , new miniFreq
    // output: cur_associated_result , higher assocaited item_sets.
    var calc = function(_ms) {
        var _dict_property = get_distinct_items(_ms),
            _item_sets = [];
        miniFreq = _ms;

        // combine distinct items to next combine level , each item_set has one item -> two items
        combine_level++;
        _item_sets = combine_items(_dict_property, combine_level);
        if (_item_sets.length <= 0) {
            combine_level--;
            return cur_associated_result;
        }

        // check freq for each item_set 
        var _cur_associated_result = [];
        _item_sets.forEach(function(e, i, a) {
            var _interaction = [];
            var _item_set = [];

            e.forEach(function(e2, i2, a2) {
                _item_set.push(e2);
                if (_interaction.length > 0) {
                    _interaction = _interaction.filter(x => items[e2].indexOf(x) >= 0);
                } else {
                    _interaction = items[e2];
                }
            });

            _cur_associated_result.push({
                item_set: _item_set,
                freq: _interaction.length
            });
        });

        cur_associated_result = _cur_associated_result;
        return cur_associated_result;
    };

    // retrun instance
    return {
        getListCount: () => list.length,
        getPropertyCount: () => items.length,
        getMiniFreq: () => miniFreq,
        get_cur_associated: () => cur_associated_result,
        calc: calc,
        reset: reset,
        show: show,
        combine_items: combine_items
    }

}

/*
console.log("Sample:");

var test1 = apriori([
    ["p1", "p3", "p4"],
    ["p2", "p3", "p5"],
    ["p1", "p2", "p3", "p5"],
    ["p2", "p5"]
]);

console.log("Transaction list:", test1.getListCount());
console.log("Current associated result:", test1.get_cur_associated());
console.log("Cureent minimum frequency (for last result):", test1.getMiniFreq());

var r1 = test1.calc(2);
console.log("Current associated result:", r1);
var r1 = test1.calc(2);
console.log("Current associated result:", r1);

console.log("Test casts of combine_items:");
// for combine_items test
var t2 = test1.combine_items([1, 2, 3, 4, 5], 0);
console.log("case [1,2,3,4,5],0:", t2);
var t2 = test1.combine_items([1, 2, 3, 4, 5], 1);
console.log("case [1,2,3,4,5],1:", t2);
var t2 = test1.combine_items([1, 2, 3, 4, 5], 2);
console.log("case [1,2,3,4,5],2:", t2);
var t2 = test1.combine_items([1, 2, 3, 4, 5], 3);
console.log("case [1,2,3,4,5],3:", t2);
var t2 = test1.combine_items([1, 2, 3, 4, 5], 4);
console.log("case [1,2,3,4,5],4:", t2);
var t2 = test1.combine_items([1, 2, 3, 4, 5], 5);
console.log("case [1,2,3,4,5],5:", t2);
var t2 = test1.combine_items([1, 2, 3, 4, 5], 6);
console.log("case [1,2,3,4,5],6:", t2);
*/
