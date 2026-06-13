def containsNearbyDuplicate(self, nums, k):
    """
    :type nums: List[int]
    :type k: int
    :rtype: bool
    """
    exist = {}
                                                                                    for i in xrange(0,len(nums)):
        if nums[i] in exist and abs(i-exist[nums[i]])<=k:                                   return True
        exist[nums[i]]=i
    return False
