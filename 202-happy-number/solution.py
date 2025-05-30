class Solution:
    def isHappy(self, n: int) -> bool:
        svisit = set()

        while n not in svisit:
            svisit.add(n)
            n = self.sumOfSquares(n)

            if n == 1:
                return True
        return False

    def sumOfSquares(self,n):
        output = 0

        while n:
            digit = n % 10
            digit = digit ** 2
            output += digit
            n = n // 10
        return output