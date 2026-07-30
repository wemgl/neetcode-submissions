func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sletters, tletters := make(map[byte]int, len(s)), make(map[byte]int, len(s))
    for i := 0; i < len(s); i++ {
        sletters[s[i]] += 1
        tletters[t[i]] += 1
    }

    for k, v := range sletters {
        if tv, ok := tletters[k]; !ok || tv != v {
            return false
        }
    }

    return true
}
