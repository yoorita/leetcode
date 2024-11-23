class Solution:
    @staticmethod
    def isPalindrome(x: int) -> bool:
        if x < 0:
            return False

        orig = x
        reverse = 0

        while x > 0:
            digit = x % 10
            reverse = reverse * 10 + digit
            x //= 10

        return True if orig == reverse else False

if __name__ == "__main__":
    print(Solution.isPalindrome(123))
    print(Solution.isPalindrome(121))