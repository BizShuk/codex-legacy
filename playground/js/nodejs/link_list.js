// TODO:
//  link list prototype pop()
//  link list prototype remove( node )

(function(root){
    
    root.linknode = function(value,next,prev){
        this.value = value || 0;
        this.next = next || null;
        this.prev = prev || null;
    };


    root.linklist = function(){
        this.head = null;
        this.point = this.head;
        this.tail = null;
        this.list = [];
        this.length = this.list.length;
    }
 
    root.linklist.prototype.start = function(){
        return this.head;
    }

    root.linklist.prototype.point = function(){
        return this.point;    
    }

    root.linklist.prototype.next = function(){
        this.point = this.point.next;
        return this.point;    
    }
    
    root.linklist.prototype.prev = function(){
        this.point = this.point.prev;
        return this.point;    
    }

    root.linklist.prototype.push = function( node ){
        if( this.list.length == 0 ) this.head = node;
        this.tail = node;
        
        Array.prototype.push.apply( this.list , [node] );
        this.length = this.list.length;
        return this;
    }


    root.linklist.prototype.pop = function(){}
    root.linklist.prototype.remove = function( node ){}

    root.linklist.prototype.reverse = function(){
        var prev = temp = null;
        var current = this.head;
        this.tail = this.head;
        while(current != null){
            var temp = current.next;
            current.next = prev;
            prev = current;
            current.prev = temp;
            current = temp;
        }
        this.head = prev;
        return this;    
    }
    
    // start from head and pass each node to callback function
    root.linklist.prototype.foreach = function(callback){
        console.log(callback);
        var point_tmp = this.head;
        while( point_tmp != null ){
            callback.apply(null,[point_tmp]);
            point_tmp = point_tmp.next
        }
        
    }

})(global);


var link_list_sample = new linklist();

var i = 0;
var prev = null;
for( i=0 ; i <=5 ; i++ ){
    var node = new linknode( i , null , prev );
    
    if( link_list_sample.length > 0 ){
        node.prev.next = node;
    }
    prev = node;
    link_list_sample.push( node );
}

console.log("\x1b[32m link list sample:\x1b[m\n");

var tmp = link_list_sample.start();
while( tmp != null ){
    console.log( tmp );
    tmp = tmp.next;
}


link_list_sample.reverse();
console.log("\x1b[32m reverse link list:\x1b[m\n");

link_list_sample.foreach(function(node){ 
    console.log(node);
});





