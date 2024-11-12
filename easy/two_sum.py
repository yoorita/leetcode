class Solution:
    @staticmethod
    def two_sum(nums: list[int], target: int) -> list[int]:
        saved = dict()
        for index, num in enumerate(nums):
            needed = target - num
            needed_index = saved.get(needed)
            if needed_index is not None:
                return [needed_index, index]
            else:
                saved[num] = index
        return [0, 0]

if __name__ == '__main__':
    print(Solution.two_sum([3, 3], 6))
