import "reflect"
func isAnagram(s string, t string) bool {
    var smap = make(map[byte]int)
    var tmap = make(map[byte]int)

    for c := range s{
        if _,ok := smap[s[c]]; ok{
            smap[s[c]] += 1
        } else {
            smap[s[c]] = 1
        }
    }
    for c := range t{
        if _,ok := tmap[t[c]]; ok {
            tmap[t[c]] += 1
        } else {
            tmap[t[c]] = 1
        }
    }
    return reflect.DeepEqual(smap, tmap)
}