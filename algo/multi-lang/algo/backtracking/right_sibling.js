var tree = require('../tree/tree.js')

var right_sibling = function(root) {
        var walk_stack = [];
        var pre_stack = [];
        walk(root, walk_stack,pre_stack);
};

function walk(node,w,p){
    if (!(node instanceof tree.treeNode)) {
        return false;
    }
    if (p.length !== 0){ 
        var pointed = p.pop();
        pointed.next=node;
    }

    w.push(node);
    if (node.left)    walk(node.left,w,p);
    if (node.right)   walk(node.right,w,p);
    p.push(w.pop());
}


exports.right_sibling = right_sibling



var nodes = tree.buildtree([1,2,3,4,5,6,7,8,9,10,11,12,13,14,15])
right_sibling(nodes[0])
right_sibling({})
tree.printtree(nodes[0])

