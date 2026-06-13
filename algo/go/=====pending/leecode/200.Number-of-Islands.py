class Solution(object):

    def numIslands(self, grid):
        """
        :type grid: List[List[str]]
        :rtype: int
        """

        def bfs(grid,i,j,count):

            grid[i][j] = str(count)
            if i-1>=0 and grid[i-1][j]==u'1':
                grid[i-1][j] = count
                bfs(grid,i-1,j,count)
            if i+1<len(grid) and grid[i+1][j]==u'1':
                grid[i+1][j] = count
                bfs(grid,i+1,j,count)
            if j-1>=0 and grid[i][j-1]==u'1':
                grid[i][j-1] = count
                bfs(grid,i,j-1,count)
            if j+1<len(grid[i]) and grid[i][j+1]==u'1':
                grid[i][j+1] = count
                bfs(grid,i,j+1,count)

        count = 1

        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[i][j] == u'1':
                    count+=1
                    bfs(grid,i,j,count)
        return count-1




