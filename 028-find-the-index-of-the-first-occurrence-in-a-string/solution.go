func strStr(haystack string, needle string) int {
    for i:=0; i<len(haystack);i++{
        if haystack[i] == needle[0]{
            if i+len(needle) <= len(haystack) && haystack[i:i+len(needle)] == needle{
                return i
            }
        }
    }
    return -1
}