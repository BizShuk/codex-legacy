// 將一個 MxN的矩陣 從一點到其他點的步數 with walkable retriction


// ori: 2d matrix
// (x, y): start point
function mapPace(ori, numRows, numColumns, x, y) {
    if (null == ori
        && numRows <= x
        && numColumns <= y ) return false;

    // init current tile pace
    if (! (ori[x][y] >= 0)) ori[x][y] = 0;

    var bfs_stack = [[x,y]];




    function walkTo(x,nx,y,ny) {
        function inboundary(nx,ny){

            if (0 <= nx && nx < numRows &&
                0 <= ny && ny < numColumns &&
                !(ori[nx][ny]<0)) return true;
            return false;
        };

        function lessSteps(x,nx,y,ny) {
            // console.log(nx,ny);
            if (ori[nx][ny] == null || ori[nx][ny] > ori[x][y]) return true;
            return false;
        };

        var current_steps = ori[x][y];

        // console.log(bfs_stack);
        if (inboundary(nx,ny) && lessSteps(x,nx,y,ny)) {
            ori[nx][ny] = current_steps+1;
            bfs_stack.push([nx,ny]);
        }

    }

    while(bfs_stack.length>0) {
        var startpoint = bfs_stack.pop();

        walkTo(startpoint[0],startpoint[0],startpoint[1],startpoint[1]+1);
        walkTo(startpoint[0],startpoint[0],startpoint[1],startpoint[1]-1);
        walkTo(startpoint[0],startpoint[0]+1,startpoint[1],startpoint[1]);
        walkTo(startpoint[0],startpoint[0]-1,startpoint[1],startpoint[1]);

    }


    console.log(ori);
    return ori;
}

var matrix = [
    [,,,,,-5],
    [,,,-1,,],
    [,,-1,,,],
    [,,,-4,,],
    [,-4,,,-3,]
]


mapPace(matrix,matrix.length,matrix[0].length,3,5);