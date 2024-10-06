// cmd/golang-pointer/main.go
package main

import "fmt"

func main() {
    var a int = 50
    var b int = 20

    var p1 *int = &a
    var p2 *int = &a
    var p3 *int = &b

    fmt.Printf("p1 == p2 : %v\n", p1 == p2)
    fmt.Printf("p2 == p3 : %v\n", p2 == p3)

    var p *int
    
    if p != nil {
        fmt.Printf("p의 값: %p\n", p)
    } else {
        fmt.Println("p 는 nil입니다")
    }
}