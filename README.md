아래는 실습 순서에 맞춰 다시 작성한 `README.md`입니다.

---

# golang-pointer

`golang-pointer`는 Golang으로 작성된 간단한 애플리케이션으로, 포인터의 사용을 익히기 위한 실습입니다.

## 포인터 변수 선언

포인터 변수는 가리키는 데이터 타입 앞에 *를 붙여서 선언합니다.
```go
var p *int
```
```go
var a int
var p *int
p = &a	// a의 메모리 주소를 포인터 변수 p에 대입합니다.
```

## 프로젝트 구조

```plaintext
golang-pointer/
│
├── cmd/
│   └── golang-pointer/
│       └── main.go
│
├── go.mod
├── Makefile
└── README.md
```

## 실습 순서

### 1. 패키지 구조를 위한 디렉토리 생성

먼저 프로젝트 디렉터리를 설정하고 필요한 디렉터리들을 생성합니다.

```bash
mkdir golang-pointer
cd golang-pointer
go mod init golang-pointer

mkdir -p cmd/golang-pointer
```

### 2. 포인터 변수 선언

`cmd/golang-pointer/` 디렉터리 아래에 `main.go` 파일을 생성하고,
포인터 정의 및 활용하는 코드를 작성합니다.

```go
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
```

이 코드를 통해 프로그램을 실행하면 포인터의 값이 출력됩니다.

### 3. `Makefile` 작성

이제 프로젝트의 빌드 및 실행을 자동화하기 위한 `Makefile`을 프로젝트 루트에 작성합니다.

```makefile
# Go 관련 변수 설정
APP_NAME := golang-pointer
CMD_DIR := ./cmd/golang-pointer
BUILD_DIR := ./build

.PHONY: all clean build run test fmt vet install

all: build

# 빌드 명령어
build:
	@echo "==> Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)

# 실행 명령어
run: build
	@echo "==> Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

# 코드 포맷팅
fmt:
	@echo "==> Formatting code..."
	go fmt ./...

# 코드 분석
vet:
	@echo "==> Running go vet..."
	go vet ./...

# 의존성 설치
install:
	@echo "==> Installing dependencies..."
	go mod tidy

# 테스트 실행
test:
	@echo "==> Running tests..."
	go test -v ./...

# 빌드 정리
clean:
	@echo "==> Cleaning build directory..."
	rm -rf $(BUILD_DIR)
```

`Makefile`을 이용하여 코드를 빌드하고 실행할 수 있습니다.

```bash
make run
```

이 명령어를 통해 `main.go`에서 작성한 struct 정보를 확인할 수 있습니다.

### 4. 포인터 변숫값 비교하기

포인터의 값을 비교하는 방법에 대해 알아보겠습니다.

```go
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
```

이제 `make run` 명령을 사용하면 각 조건의 결과가 출력됩니다.

```bash
make run
```

### 5. 포인터의 사용이유

변수 대입이나 함수 인수 전달은 항상 값을 복사하기 때문에 많은 메모리 공간을 사용하는 문제와
큰 메모리 공간을 복사할 때 발생하는 성능 문제를 안고 있습니다.
또한 다른 공간으로 복사되기 때문에 변경 사항이 적용되지도 않습니다.

```go
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
```

이제 `make run` 명령을 사용하면 data의 값이 출력됩니다.

```bash
make run
```

### 6. 인스턴스

인스턴스란 메모리에 할당된 데이터의 실체를 말합니다.

```go
// cmd/golang-pointer/main.go
package main

import "fmt"

type Data struct {
    value int
    data [200]int
}

func main() {
    var data Data
    var p *Data = &data // 기존에 있는 data 인스턴스를 가리킵니다.

    fmt.Printf("p의 값: %p\n", p)

    var p1 *Data = &Data{}  // Data 인스턴스를 생성합니다.
    var p2 *Data = p1       // 인스턴스를 가리킵니다.
    var p3 *Data = p1       // 인스턴스를 가리킵니다.

    fmt.Printf("p1의 값: %p\n", p1)
    fmt.Printf("p2의 값: %p\n", p2)
    fmt.Printf("p3의 값: %p\n", p3)

    var data1 Data
    var data2 Data = data1  // data1과 값만 같고, 인스턴스가 생성됩니다.
    var data3 Data = data1  // data1과 값만 같고, 인스턴스가 생성됩니다.

    fmt.Printf("data1의 값: %p\n", &data1)
    fmt.Printf("data2의 값: %p\n", &data2)
    fmt.Printf("data3의 값: %p\n", &data3)
}
```

이제 `make run` 명령을 사용하면 pointer의 인스턴스를 확인할 수 있습니다.

```bash
make run
```

### 7. new() 내장함수 와 인스턴스 소멸 시점

포인터값을 별도의 변수를 선언하지 않고 초기화 하는 방법을 봤습니다.
new 내장함수를 사용하면 더 간단히 표현가능합니다.

```go
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
    mt.Println("u: ", &u)
}   // 내부 변수 u는 사라집니다. 더불어 인스턴스도 사라집니다.

func main() {
    p1 := &Data{}        // &를 사용하는 초기화
    var p2 = new(Data)  // new()를 사용하는 초기화

    fmt.Printf("p1의 값: %p\n", p1)
    fmt.Printf("p2의 값: %p\n", p2)

    TestFunc()
}
```

이제 `make run` 명령을 사용하면 new() 함수와 인스턴스 소멸을 확인할 수 있습니다.

```bash
make run
```

### 8. 스택 메모리와 힙 메모리

이론상 스택 메모리 영역이 힙 메모리 영역보다 훨씬 효율적이기 때문에 스택 메모리 영역에 메모리를 할당하는 게 더 좋지만, 스택 메모리는 함수 내부에서만 사용 가능한 영역입니다. 그래서 함수 외부로 공개되는 메모리 공간은 힙 메모리 영역에서 할당합니다.

Go 언어는 탈출 검사(escape analysis)를 해서 어느 메모리에 할당할지 결정합니다.

```go
// cmd/golang-pointer/main.go
package main

import "fmt"

type User struct {
    Name    string
    Age int
}

func NewUser(name string, age int) *User{
    var u = User{name, age}
    fmt.Printf("u의 주소값: %p\n", &u)
    return &u   // 탈출 분석으로 u 메모리가 사라지지 않음
}

func main() {
    userPointer := NewUser("AAA", 23)

    fmt.Println(userPointer)
    fmt.Printf("userPointer의 주소값: %p\n", userPointer)
}
```

이제 `make run` 명령을 사용하면 User의 메모리 값을 확인할 수 있습니다.

```bash
make run
```

Go 언어에서는 탈출 검사를 통해서 u 변수의 인스턴스가 함수 외부로 공개되는 것을 분석해내서 u를 스택 메모리가 아닌 힙 메모리에서 할당하게 됩니다. 메모리 공간이 함수 외부로 공개되는지 여부를 자동으로 검사해서 스택 메모리에 할당할지 힙 메모리에 할당할지 결정합니다.
