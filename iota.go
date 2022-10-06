// iota.go
package main

import "fmt"

func main() {
    /*
    iota 는 0부터 시작하여 1씩 증가하는 incremental constant입니다.
    */
    const (
        JAN int = iota + 1
        FEB
        MAR
        APR
        MAY
        JUN
        JUL
        AUG
        SEP
        OCT
        NOV
        DEC
    )
    fmt.Println(JAN) // 1
    fmt.Println(DEC) // 12
}

