package main

import (
    "fmt"
)

func minSwap(a, b int) (int, int) {
    if a<b {
        return a, b
    }
    return b, a
}

func main GCD(a, b int) int {
    a, b = minSwap(a, b)
    for a > 1 {
        temp := b%a
        if temp == 0 {
            return a
        }
        a, b = temp, a
    }
    return a
}

func main(){
    a, b := 10, 5
    fmt.Println(GCD(a, b))
    // fmt.Println(LCM(a, b))
}