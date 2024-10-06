// cmd/golang-pointer/main.go
package main

import "fmt"

type Data struct {
    value int
    data [200]int
}

func main() {
    var data Data
    var p *Data = &data

    fmt.Printf("p의 값: %p\n", p)

    var p1 *Data = &Data{}
    var p2 *Data = p1
    var p3 *Data = p1

    fmt.Printf("p1의 값: %p\n", p1)
    fmt.Printf("p2의 값: %p\n", p2)
    fmt.Printf("p3의 값: %p\n", p3)

    var data1 Data
    var data2 Data = data1
    var data3 Data = data1

    fmt.Printf("data1의 값: %p\n", data1)
    fmt.Printf("data2의 값: %p\n", data2)
    fmt.Printf("data3의 값: %p\n", data3)
}