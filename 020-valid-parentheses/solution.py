class Solution:
    def isValid(self, s: str) -> bool:
        brackets = {
            ')':'(',
            '}':'{',
            ']':'['
        }
        visit = []
        str = list(s)

        for c in str:
            if c not in brackets:
                visit.append(c)
            elif c in brackets:
                if len(visit) != 0:
                    removed_ele = visit[-1]
                    if removed_ele == brackets[c]:
                        visit.pop()
                        continue
                    else:
                        break
                else:
                    visit.append(c)
                    
        return len(visit)==0
        