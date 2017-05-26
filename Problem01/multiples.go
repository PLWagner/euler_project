package main

import "fmt"

const m1 = 3
const m2 = 5

func check_if_multiple_are_equal(n int) bool {
    if n%m1 == 0 && n%m2 == 0 {
        return true
    }
    return false
}

func main() {
    var num, sum = 1000, 0
    for i := 1; i < num; i++ {
        if check_if_multiple_are_equal(i) {
            sum = sum + i
        } else {
            if i%m1 == 0 {
                sum = sum + i
            } else if i%m2 == 0 {
                sum = sum + i
            }
        }
    }
    fmt.Println(sum)
}