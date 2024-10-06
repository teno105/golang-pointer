// cmd/golang-pointer/main.go
package main

import "fmt"

func main() {
    var a int = 500
    var p *int

    p = &a

    fmt.Printf("p의 값: %p\n", p)   // 메모리 주솟값 출력
    fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p) // p가 가리키는 메모리의 값 출력

    *p = 100
    fmt.Printf("a의 값: %d\n", a)   // a의 값 변화 확인
}