package main

import "strings"

func IsPalindrome(s string) bool {
    s = strings.ToLower(s)

    // remover caracteres não alfanuméricos
    limpo := ""
    for _, c := range s {
        if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
            limpo += string(c)
        }
    }

    i, j := 0, len(limpo)-1
    for i < j {
        if limpo[i] != limpo[j] {
            return false
        }
        i++
        j--
    }

    return true
}
