package strutil
import "strings"
func Reverse(s string) string {
    result := ""
    for i := len(s) - 1; i >= 0; i-- { //с конца
        result += string(s[i])
    }
    return result
}
func ToUpper(s string) string {
    return strings.ToUpper(s)
}