class Solution {
public:
    bool removeOnes(vector<vector<int>>& grid) {
        bool issame = true;
        for(int i=1;i<grid.size();i++)
        {
            bool issame=true;
            if(grid[i][0]!=grid[i-1][0])
            {
                issame=false;
            }
            for(int j=1;j<grid[i].size();j++)
            {
                if(grid[i-1][j]!=grid[i][j] && issame==true)
                {
                    return false;
                }
                else if(grid[i][j]==grid[i-1][j] && issame==false)
                {
                    return false;
                }
            }
        }
        return true;
    }
};