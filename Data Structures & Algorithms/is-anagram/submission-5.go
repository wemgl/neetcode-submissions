import "reflect"

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sletters, tletters := make(map[byte]int, len(s)), make(map[byte]int, len(s))
    for i := 0; i < len(s); i++ {
        sletters[s[i]] += 1
        tletters[t[i]] += 1
    }

    // for _, c := range t {
    //     if count, ok := letters[c]; ok {
    //         if count == 0 {
    //             delete(letters, c)
    //         } else {
    //             letters[c] -= 1
    //         }
    //     }
    // }

    return reflect.DeepEqual(sletters, tletters)
}
