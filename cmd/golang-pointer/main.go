// cmd/golang-pointer/main.go
package main

import "fmt"

type Data struct {
    value int
    data [200]int
}

func TestFunc() {
    u := &Data{}    // u 포인터 변수를 선언하고 인스턴스를 생성합니다.
    u.value = 30
    fmt.Println(u)
}   // 내부 변수 u는 사라집니다. 더불어 인스턴스도 사라집니다.

func main() {
    p1 := &Data{}        // &를 사용하는 초기화
    var p2 = new(Data)  // new()를 사용하는 초기화

    fmt.Printf("p1의 값: %p\n", p1)
    fmt.Printf("p2의 값: %p\n", p2)

    TestFunc()
}