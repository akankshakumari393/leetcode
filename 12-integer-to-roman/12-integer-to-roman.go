func intToRoman(num int) string {
    var ans []string
    for ;num > 0; {
        a, n := Category(num)
        ans = append(ans, a)
        num = n
    }
    // slice of string to string using join
    return strings.Join(ans, "")
}

func Category(num int) (string, int) {
    if num / 1000 >= 1 {
        return "M", num-1000
    } else if num >= 900 {
        return "CM", num-900
    }else if num / 500 >= 1 {
        return "D", num-500
    } else if num >= 400 {
        return "CD", num-400
    }else if num / 100 >= 1 {
        return "C", num-100
    } else if num >= 90 {
        return "XC", num-90
    }else if num / 50 >= 1 {
        return "L", num-50
    } else if num >= 40 {
        return "XL", num-40
    }else if num / 10 >= 1 {
        return "X", num-10
    } else if num == 9 {
        return "IX", num-9
    }else if num / 5 >= 1 {
        return "V", num-5
    } else if num == 4 {
        return "IV", num-4
    }else if num <= 3 {
        return "I", num-1
    }
    return "", 0
}