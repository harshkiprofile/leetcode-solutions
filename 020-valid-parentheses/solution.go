func isValid(s string) bool {
    brackets := map[string]string{
        ")": "(",
        "]": "[",
        "}": "{",
    }
    var visit []string
    for i := 0; i < len(s); i++ {
        char := string(s[i])
        if _, ok := brackets[char]; ok {
            if len(visit) > 0 {
                top := visit[len(visit)-1]
                if top == brackets[char] {
                    visit = visit[:len(visit)-1]
                } else {
                    return false
                }
            } else {
                return false
            }
        } else {
            visit = append(visit, char)
        }
    }
    return len(visit) == 0
}
