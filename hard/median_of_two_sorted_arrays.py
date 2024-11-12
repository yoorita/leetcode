class Solution:
    @staticmethod
    def findMedianSortedArrays(nums1: list[int], nums2: list[int]) -> float:
        nums = sorted(nums1 + nums2)
        center = len(nums) // 2
        if len(nums) % 2 != 0:
            return float(nums[center])
        return (nums[center-1] + nums[center]) / 2

if __name__ == '__main__':
    print(Solution.findMedianSortedArrays([1, 2], [3]))
