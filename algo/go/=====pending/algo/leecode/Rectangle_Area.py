ass Solution(object):
    def computeArea(self, A, B, C, D, E, F, G, H):
        """
        :type A: int
        :type B: int
        :type C: int
        :type D: int
        :type E: int
        :type F: int
        :type G: int
        :type H: int
        :rtype: int
        """
        cover_Ax = A 
        cover_By = B
        cover_Cx = C
        cover_Dy = D
        cross_area = 0 
        
        if cover_Ax < E:
            cover_Ax = E
            
        if cover_Cx > G :
            cover_Cx = G
            
        if cover_By < F:
            cover_By = F
            
        if cover_Dy > H:
            cover_Dy = H
            
        if cover_Dy > cover_By and cover_Cx > cover_Ax:
            cross_area = ( cover_Cx - cover_Ax ) * ( cover_Dy - cover_By )
        
        return (C-A)*(D-B) + (G-E)*(H-F) - cross_area
            
            
            
        
