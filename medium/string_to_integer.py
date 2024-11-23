class Solution:
    @staticmethod
    def myAtoi(s: str) -> int:
        s = s.lstrip()
        if len(s) == 0:
            return 0

        answer = 0
        sign = None
        signed = ['+', '-']
        valid = [str(i) for i in range(10)]

        if s[0] in signed:
            sign = s[0]
            s = s[1:]

        if len(s) == 0:
            return 0

        for char in s:
            if char in valid:
                answer = answer * 10 + int(char)
            else:
                break

        if sign == '-':
            answer = answer * -1

        if answer > 2 ** 31 - 1:
            return 2 ** 31 - 1
        if answer < -2 ** 31:
            return -2 ** 31

        return answer


if __name__ == "__main__":
    print(Solution.myAtoi("42"))
    print(Solution.myAtoi(" -42"))
    print(Solution.myAtoi("1337c0d3"))
    print(Solution.myAtoi("0-1"))
    print(Solution.myAtoi("words and 987"))
    print(Solution.myAtoi("-91283472332"))
    print(Solution.myAtoi("3.14159"))
    print(Solution.myAtoi(".1"))
    print(Solution.myAtoi("+-12"))
    print(Solution.myAtoi("-+12"))
