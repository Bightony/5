package numutil
func IsPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
func Factorial(n int) int {
    if n <= 0 {
        return 1
    }
    result := 1
    for i := 1; i <= n; i++ {
        result = result * i
    }
    return result
}