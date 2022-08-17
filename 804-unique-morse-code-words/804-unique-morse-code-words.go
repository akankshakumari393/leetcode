func uniqueMorseRepresentations(words []string) int {
    trans := make(map[string]struct{})
    for _, word := range words {
        morse := GetMorse(word)
        trans[morse] = struct{}{}
    }
    return len(trans)
}


func GetMorse(word string) string {
    maps := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    out := make([]string, len(word))
    for i, c := range word {
        out[i] = maps[c - 'a']
    }
    return strings.Join(out, "")
}