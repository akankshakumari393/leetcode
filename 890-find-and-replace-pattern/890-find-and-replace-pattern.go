// Encoding Decoding, Two map Problem
func findAndReplacePattern(words []string, pattern string) []string {
    result := []string{}
    
    for _, word := range words {
        if check(word, pattern) {
            result = append(result, word)            
        }        
    }
    
    return result
}

func check(word string, pattern string) bool {
    encode, decode := map[byte]byte{}, map[byte]byte{}
    for i := range word {
        if c, found := encode[word[i]]; found {
            if pattern[i] != c {
                return false
            }

            continue
        }

        if _, found := decode[pattern[i]]; found {
            return false
        }

        decode[pattern[i]] = word[i]
        encode[word[i]] = pattern[i]
    }
    return true
}