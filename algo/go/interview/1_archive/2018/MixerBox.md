# 20180806



# Question 1
Play with Building Blocks Vera is a great mom, she usually spends a great amount of time playing with her son with various types of toys.

Vera and her son were playing cubic shaped building blocksearlier today. The blocks can be used to build towers. A tower is defined as a single stack of blocks. All blocks are identical.

They built these blocks into a row with ntowers, and the i-th tower from the left was built with biblocks.

Vera wants to train her son's mathematical ability, so she gave her son a difficult problem: In order to make the i-th tower taller than every other tower on its right, how many moreblocks does he need to add to the i-th tower?


Input Format
The first line would have a positive integer n (1 ≤n ≤ 105), representing how many towers Vera and her son built in total.
The second line contains n spaced-separated positive integers bi (1 ≤ bi ≤ 109), biis the height of the i-th tower from the left.


Output Format
Print nnumbers in a line. The i-th number represents the minimum number of blocks required for the i-th tower. If the i-th toweris already higher than all other towers on its right hand side, please print 0 asthe i-th number.



##### Sample Input 1
6
3 7 11 9 14 11

##### Sample Output 1
12 8 4 6 0 0

    For the 1sttower (3), it needs to be 15 in order to be taller than all other towers on its right hand side (7 11 9 14 11),so it needs 15-3=12 more blocks.

    For the 2nd tower (7),it needs to be 15 in order to be taller than all other towers on its right hand side (11 9 14 11),so it needs 15-7=8 more blocks.

    For the 3rd tower (11), it needs to be 15 in order to be taller than all other towers on its right hand side (9 14 11),so it needs 15-11=4more blocks.

    For the 4th tower (9), it needs to be 15 in order to be taller than all other towers on its right hand side (14 11),so it needs 15-9=6more blocks.

    For the 5th tower (14), it needs 0 block in order to be taller than all other towers on its right hand side (11).

    For the 6th tower (11), it needs 0 block in order to be taller than all other towers on its right hand side.




# Question 2
Given a board with n rows and n columns. Each cell contains either a symbol 'X' or a symbol 'O'. Now, you're asked to change all 'X' into 'O'. The only allowed action is (you can perform the action asmany times as you want):
Choose an 'X' which has exactly fourother 'X's sharing a sidewithit. Then change the chosen 'X' along with the four connected(shared a side) 'X's into 'O'. Is it possible to complete this task?


### Input Format
The first line contains a positive integern,the side length of the board.
The following n line describes the board.
The j-th characterin thei-thline of those n lines denotes the symbol in thei-th row and j-th columnon the board.


### Output Format
If it's possible to complete the task, output `Possible`; Otherwise, output `Impossible`.

### Constraint
    3 ≤ n ≤ 100.
    Each character describing the board is either'O' or 'X'.

##### Input
5
OOOXO
OXXXX
XXXXO
OXOOO
OOOOO

##### Output
Possible

##### Explanation
You can do two actions to change all 'X' into 'O'. For example, you can choose the 'X' symbol in row 3, column 2 in the first action and choose the 'X' symbol in row 2, column 4.


### Sample 1
##### Input
4
OXXO
XXXX
OXXO
OOOO

##### Output
Impossible




# Question 3
Consider a desktop which is a rectangle of unit cells with rows and columns.
Each cell has a diagonally-placed double-sided mirror which looks like

..../
:  /:
: / :
:/  :
/....

or

\....
:\  :
: \ :
:  \:
....\

.

The desktop can be described as aby array using only `/` and `\`. 
For example:
//
\/

is a 2 by 2 desktop with 3 top right to bottom left mirror and a top left to bottom right one.



Eddy got a laser beam, he shoots the laser from some boundary of the desktop.
He wants to know whether the beam passes some cell he is interested in.

For example, suppose Eddy shoots the beam from thetop ofthe top right cell as i the previous example, the beam passesand flies away.



### Input Format
The first line contains two space-separated integers representingandrespectively.
The next n lines describe the desktop, where the i+1th line of input represents the ith row of the desktop by a string of length m.
The n+2th line contains a integerrepresenting the number of queries.
Each of the nextlines describes the query in formatsx sy side tx tymeaning Eddy shoots the beam from thesideside of the boundary cell (sx, sy), and is asking whether the beam passes the cell (tx, ty).



Note: The coordinates of thetop leftcell is (0, 0) while the bottom left cell is (n-1, 0) and thebottom right cell is (n-1, m-1).

### Constraint
    1 &lt;= n, m &lt;= 50
    1 &lt;= q&lt;= 100
    0 &lt;= sx, tx &lt;n
    0 &lt;= sy, ty &lt;m
    sideis eitherU(top),L(left),D(down) orR(right).
    It is guaranteed thatsideside of the cellmust be on the boundary of the desktop.


### Output Format
For each query, outputYesin a single line if the beam does pass the cell and ,Nootherwise.

##### Input
3 4
///\
\/\/
\\\\
2
0 0 U 1 2
0 1 U 1 1

##### Output
No
Yes

##### Explanation

In the first query, the beam passes only the cell (0, 0).
In the second query, the beam passes(0, 1), (0, 0), (1,0), (1,1), (0,1), (0,2)and flies out.




