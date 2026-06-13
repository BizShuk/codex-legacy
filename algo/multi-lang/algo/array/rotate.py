def rotate(self, nums, k):
    """
    :type nums: List[int]
    :type k: int
    :rtype: void Do not return anything, modify nums in-place instead.
    """
    k = len(nums) - k % len(nums)
    for i in xrange(0, k):
        nums.append(nums.pop(0))
