function Node(n){
    this.val = n;
    this.left = this.right = null;

}

function treeConstructor(binaryTreeArr, i){
    if (!(binaryTreeArr instanceof Array) || binaryTreeArr == null) return null;
    if (binaryTreeArr[i] == null) return null;

    var root = new Node(binaryTreeArr[i]);

    let curNode = root;
    if (binaryTreeArr[i*2+1]!=null) curNode.left  = treeConstructor(binaryTreeArr,i*2+1)
    if (binaryTreeArr[i*2+2]!=null) curNode.right = treeConstructor(binaryTreeArr,i*2+2)

    return root;
}




function outputTree(root) {
    if (root == null) return null;
    var outputGraph = { maxLength:0, graph:[""]};
    let leftGraph = outputTree(root.left) || {maxLength:0,graph:[]};
    let rightGraph = outputTree(root.right) || {maxLength:0,graph:[]};

    appendSpace(outputGraph.graph,0,leftGraph.maxLength," ")
    appendSpace(outputGraph.graph,0,1,root.val)
    outputGraph.graph.push(...leftGraph.graph);
    appendRangeSpace(outputGraph.graph,1,outputGraph.graph.length-1, 1, " ")
    appendRightGraph(outputGraph.graph, rightGraph)
    appendSpace(outputGraph.graph,0,rightGraph.maxLength," ")

    outputGraph.maxLength = outputGraph.graph[0].length
    return outputGraph
}

function appendSpace(arr, index, spaceLength, c){
    if (c == null) c = " ";
    let str = new Array(spaceLength+1).join(c.toString())
    arr[index] += str;
    return arr;
}

function appendRangeSpace(arr, from, to, spaceLength, c) {
    if (c == null) c = " ";
    let spaces = new Array(spaceLength+1).join(c);
    for (var i=from;i<=to; i++){
        arr[i] = arr[i].concat(spaces);
    }
    return arr;
}

function appendRightGraph(arr, rightGraph) {
    let leftGraphSize = arr[0].length, leftGraphRow=arr.length;
    for (var i = 0; i < rightGraph.graph.length ; i++) {
        if (i+1>=leftGraphRow) {
            arr.push("")
            appendSpace(arr, i+1, leftGraphSize , " ")
        }
        arr[i+1] = arr[i+1].concat(rightGraph.graph[i]);
    }
    return arr;
}


let tree1 = ["(1)","(2)","(3)","(4)","(5)","(6)","(7)"],
    tree2 = ["(1)",null,"(3)",null,null,"(6)",null,null,null,null,null,"(12)","(13)"],
    tree3 = ["(1)","(2)","(3)"],
    tree4 = ["(1)","(2)",null,"(4)",null,null,null,"(8)",null,null,null,null,null,null,null,"(16)"],
    tree5 = ["(1)","(2)","(3)",null,null,"(6)","(7)",null,null,null,null,"(12)","(13)",null,null,null,"(17)","(18)","(19)","(20)","(21)","(22)","(23)","(24)","(25)"],
    tree6 = ["(1)","(2)","(3)"];

let root1 = treeConstructor(tree1,0),
    root2 = treeConstructor(tree2,0),
    root3 = treeConstructor(tree3,0),
    root4 = treeConstructor(tree4,0),
    root5 = treeConstructor(tree5,0);

let clog = (root) => {
    let g = outputTree(root);
    for (var i = 0; i < g.graph.length; i++) {
        console.log(g.graph[i])
    }
}

clog(root5)