package main

import "fmt"
/*
"fmt" imported but not used
Compiler: UnusedImport
*/

func test_func() (int, int) {
    return 3, 4
}

func main() {
    var a float64 = 0.1
    var b float64 = 0.2
    var c float64 = 0.3

    fmt.Printf("%f + %f == %f : %v \n", a, b, c, a + b == c)
}
