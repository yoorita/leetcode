class Solution:
    @staticmethod
    def lengthOfLongestSubstring(s: str) -> int:
        max_len = 0
        for i, v in enumerate(s):
            sub = ""
            for j in range(i, len(s)):
                if s[j] in sub:
                    break
                sub += s[j]
                max_len = len(sub) if len(sub) > max_len else max_len

        return max_len

if __name__ == '__main__':
    print(Solution.lengthOfLongestSubstring("au"))
