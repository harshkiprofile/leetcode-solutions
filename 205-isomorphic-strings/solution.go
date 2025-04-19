// Difficult to read; but easy to debug

func isIsomorphic(s string, t string) bool {
    var smap = make(map[string]string)
    var tmap = make(map[string]string)

    for i := range s {
        c1, c2 := s[i], t[i]

        if _,ok := smap[string(c1)]; ok && smap[string(c1)] != string(c2) {
            return false
        }
        if _,ok := tmap[string(c2)]; ok && tmap[string(c2)] != string(c1) {
            return false
        }

        smap[string(c1)] = string(c2)
        tmap[string(c2)] = string(c1)
    }
    return true
    
}

// easy to read; difficult to debug

func isIsomorphic(s string, t string) bool {
    var smap = make(map[uint8]uint8)
    var tmap = make(map[uint8]uint8)

    for i := range s {
        c1, c2 := s[i], t[i]

        if _,ok := smap[c1]; ok && smap[c1] != c2 {
            return false
        }
        if _,ok := tmap[c2]; ok && tmap[c2] != c1 {
            return false
        }

        smap[c1] = c2
        tmap[c2] = c1
    }
    return true
    
}