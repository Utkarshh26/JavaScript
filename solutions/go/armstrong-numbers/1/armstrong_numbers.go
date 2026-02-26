package armstrong

func intPow(base, exp int) int {
    result := 1
    for i := 0; i < exp; i++ {
        result *= base
    }
    return result
}

func IsNumber(n int) bool {
    if n < 0 {
        return false
    }

    pow := 0
    for temp := n; temp > 0; temp /= 10 {
        pow++
    }

    sum := 0
    for temp := n; temp > 0; temp /= 10 {
        digit := temp % 10
        sum += intPow(digit, pow)
    }

    return sum == n
}
