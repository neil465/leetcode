public class Solution {
    public int IslandPerimeter(int[][] grid) {
        var row = grid.GetLength(0);// grid.Max(r=>r.Length);
        var col = grid.Min(r=>r.Length);
        int peri = 0;
        
        for (int r = 0; r<row; r++)
        {
            int left = 0;
            int top = 0;
            int bottom = 0;
            int right = 0;
           
            for (int c=0; c<col;c++)
            {
                
                
               if ((c==0)  &&  grid[r][c] ==1)
                   left++;
               else if (c!=0)
               {
                   if (((grid[r][c-1] ==0) && (grid[r][c] ==1)))
                       left++;
               }
               
               if ((c==col-1) && grid[r][c] ==1)
                   right++;
               else if (c<col-1) 
                   {
                       if (grid[r][c] ==1 && grid[r][c+1] ==0)
                           right++;
                   }
             
             if ((r==0)  &&  grid[r][c] ==1)
                   top++;
              else if (r!=0) 
              {
                  if ((grid[r-1][c]== 0) && (grid[r][c] ==1))
                      top++;
               }
             
                if ((r==row-1) && grid[r][c] ==1)
                   bottom++;
                else if (r<row-1)
                {
                    if (grid[r][c] ==1 && grid[r+1][c] ==0)
                        bottom++;
                 }
                    
              
            }
            peri += left + top + right + bottom;
        
           
                
            }
           
            
        
        //
         
        
        
        return peri;
    }
    
}