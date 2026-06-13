
var buildtree = function(arr) {
    var nodes = [];
    for (var i = 0, len = arr.length; i < len; i++) {
        nodes.push(new treeNode(arr[i]))
    }

    for (var i = 0, len = nodes.length; i < len; i++) {
        nodes[i].left = null;
        nodes[i].right= null;
        if(i*2+1 < len){
            nodes[i].left = nodes[i*2+1];
            nodes[i].right= nodes[i*2+2];
        }

    }

    return nodes

}


var treeNode = function(val) {
    this.val = val;
    this.left = this.right = this.next = null;
}

treeNode.prototype.getVal = function() {
    return this.val
}

var printtree = function(root) {
    var left , right;
    left = right = null;
    if (root.left) {
        left = root.left.val;
    }

    if (root.right) {
        right = root.right.val;
    }


    console.log("Node:",root.val,"left:",left,"right:",right);
    if (left) {
        printtree(root.left);
    }

    if (right) {
        printtree(root.right);
    }
}


var findMin = function(nums) {
    if (nums instanceof treeNode){
        return findtreemin_node(nums);
    }else if (nums instanceof Array){
        return findtreemin_arr(nums,0);
    }
}


function findtreemin_node(argument) {
    
}


function findtreemin_arr(nums,i){
    var len = nums.length;
    if (len==0) {
        return null;
    }
    var subtree_min;

    if (i*2+2 <len){
        subtree_min = Math.min(nums[i],findtreemin_arr(nums,i*2+1),findtreemin_arr(nums,i*2+2));
    }else if (i*2+1 <len){
        subtree_min = Math.min(nums[i],findtreemin_arr(nums,i*2+1));
    }else{
        subtree_min = nums[i];
    }

    return subtree_min;
}

//exports.treeNode = treeNode
module.exports = {
    treeNode:treeNode,
    printtree:printtree,
    buildtree:buildtree,
    findMin:findMin
}
    

