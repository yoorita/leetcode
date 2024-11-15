class Solution:
    @staticmethod
    def longestPalindrome(s: str) -> str:
        res = ""
        length = 0
        for i in range(len(s)):
            temp = len(s)
            for j in range(temp,i,-1):
                sub = s[i:j]
                if length < len(sub) and sub == sub[::-1]:
                    res = sub
                    length = len(sub)
        return res

if __name__ == '__main__':
    print(Solution.longestPalindrome("babad"))
    print(Solution.longestPalindrome("bb"))
    print(Solution.longestPalindrome("cbbd"))
