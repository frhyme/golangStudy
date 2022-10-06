package main

import "fmt"

func main() {
    var t[10] int

    for i := 0; i < 10; i++ {
        t[i] = i * 11
    }
    // len 을 사용하여 해당 배열의 길이를 확인할 수 있습니다.
    // 다만 정적 배열이므로 현재 코드 내에서 len(t) 는 항상 10으로 고정됩니다.
    for i := 0; i < len(t); i++ {
        fmt.Println(i, t[i])
    }

    fmt.Println("== iteration by range")
    for i, v := range t {
        fmt.Println(i, v)
    }

    // multiple array

    fmt.Println("== multiple array")
    // array length must be constant, not variable
    var mArr[4][3] int

    for i := 0; i < len(mArr); i++ {
        for j := 0; j < len(mArr[i]); j++ {
            mArr[i][j] = i * j
        }
    }

    for i := 0; i < len(mArr); i++ {
        for j := 0; j < len(mArr[i]); j++ {
            fmt.Printf("%d ", mArr[i][j])
        }
        fmt.Println()
    }
}

