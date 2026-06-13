def removeDuplicates(self, nums):
    """
    :type nums: List[int]
    :rtype: int
    """
    last_elem = None
    i, j = 0, len(nums)
    while 1:
        if i >= j:
            break
        if last_elem != nums[i]:
            last_elem = nums[i]
            i += 1
        else:
            nums.pop(i)
            j -= 1
    return len(nums)
