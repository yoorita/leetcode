class Solution:
    @staticmethod
    def convert(s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        zigzag = [[] for _ in range(numRows)]
        count = 0
        flag = True
        for letter in s:
            if count < numRows:
                zigzag[count].append(letter)
                if flag:
                    count += 1
                else:
                    count -= 1

            if count == numRows:
                count -= 2
                flag = False

            if count == 0:
                flag = True

        return "".join(["".join(i) for i in zigzag])


if __name__ == "__main__":
    print(Solution.convert("PAYPALISHIRING", 3))
    print(Solution.convert("PAYPALISHIRING", 4))
    print(Solution.convert("A", 1))
    print(Solution.convert("ABC", 1))
