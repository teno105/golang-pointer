아래는 실습 순서에 맞춰 다시 작성한 `README.md`입니다.

---

# golang-pointer

`golang-pointer`는 Golang으로 작성된 간단한 애플리케이션으로, 포인터의 사용을 익히기 위한 실습입니다.


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

구조체 변수를 초기화 하는 방법에 대해 알아보겠습니다.

```go
// cmd/golang-pointer/main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

이제 `make run` 명령을 사용하면 각 house의 정보가 출력됩니다.

```bash
make run
```

### 5. 포인터의 사용이유


### 6. 인스턴스


### 7. new() 내장함수


### 8. 스택 메모리와 힙 메모리

