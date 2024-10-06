// cmd/golang-pointer/main.go
package main

import "fmt"

type Data struct {
    value int
    data [200]int
}

func ChangeData(arg Data) {
    arg.value = 99
    arg.data[100] = 999
}

func ChangeData2(arg *Data) {
    // arg는 포인터 변수이기 때문에 (*arg).value = 999라고 써야 하지만
    // Go 언어에서는 arg.value 라고만 써도 동작합니다.
    arg.value = 999
    arg.data[100] = 999
}

func main() {
    var data Data

    ChangeData(data)    // 인수로 data를 넣습니다.
    fmt.Printf("value = %d\n", data.value)
    fmt.Printf("data[100] = %d\n", data.data[100])

    ChangeData2(&data)  // 인수로 data의 주소를 넘깁니다.
    fmt.Printf("value = %d\n", data.value)
    fmt.Printf("data[100] = %d\n", data.data[100])
}