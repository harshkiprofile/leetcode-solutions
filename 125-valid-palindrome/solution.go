func isAlphaNumeric(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
	 return true
	}
   
	return false
   }
   
   func isPalindrome(s string) bool {
   
	s = strings.ToLower(s)
   
	l := 0
	r := len(s) - 1
   
	for l < r {
   
	 for l < r && !isAlphaNumeric(s[l]) {
	  l += 1 // Shift left pointer to left + 1
	 }
	 for l < r && !isAlphaNumeric(s[r]) {
	  r -= 1 // Shift right pointer to right + 1
	 }
   
	 if s[l] != s[r] {
	  return false
	 }
	 l += 1
	 r -= 1
	}
   
	return true
   }