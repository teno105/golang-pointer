// cmd/golang-pointer/main.go
package main

import "fmt"

type User struct {
    Name    string
    Age int
}

func NewUser(name string, age int) *User{
    var u = User{name, age}
    fmt.Printf("u의 주소값: %p\n", u)
    return &u   // 탈출 분석으로 u 메모리가 사라지지 않음
}

func main() {
    userPointer := NewUser("AAA", 23)

    fmt.Println(userPointer)
    fmt.Printf("userPointer의 주소값: %p\n", userPointer)
}