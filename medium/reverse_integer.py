class Solution:
    @staticmethod
    def reverse(x: int) -> int:
        is_negative = False

        if x < 0:
            is_negative = True
            x *= -1

        res = 0
        while x > 0:
            res = res * 10 + (x % 10)
            x //= 10

        if res.bit_length() >= 32:
            return 0
        return res * -1 if is_negative else res

if __name__ == "__main__":
    print(Solution.reverse(123))
    print(Solution.reverse(-123))
    print(Solution.reverse(1563847412))
