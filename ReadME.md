# gRPC Project

cara set-up gRPC-golang project pada Linux

## Installation

GOLANG

Make sure that Golang is installed on your computer. if not, please [install](https://go.dev/doc/install) first.

check golang version

```bash
go version
```

PROTOBUF

Install Protobuf compiler on Linux

```bash
apt install -y protobuf-compiler
protoc --version  

```
Ensure that your GOBIN or GOPATH/bin directory is added to your system’s PATH environment variable. This allows the protoc compiler to find and use the protoc-gen-go-grpc plugin.

add these comments on yout bashrc in order to add the GOPATH and GOROOT

```bash
sudo gedit .bashrc
```

```bash
export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin
```


## Set up Protoc

Run go mod init {module name} to create a go.mod file for the current directory.


```bash
go mod init github.com/devinareva17/grpc
```
```bash
go mod tidy
```

buat folder untuk menyimpan proto dan buat file .proto
```
project
│└───calculator
│   └───calculatorpb
│       │   calculator.proto
│   
└───go.mod
└───go.sum
```


install modul-modul berikut.
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

run file proto
```bash
protoc --go_out=. --go-grpc_out=. calculator/calculatorpb/calculator.proto
```

## Set up Client-Server

buat folder untuk membuat main program pada client dan server
```
project
│└───calc_client
│   └───main.go
│ 
│└───calc_server
│   └───main.go
│ 
│└───calculator
│   └───calculatorpb
│       │   calculator.proto
│   
└───go.mod
└───go.sum
```

## How to Run
run server main.go
```bash
go run calc_server/main.go
```
run client main.go with adding 2 numbers to calculate
```bash
go run calc_client/main.go 10 2
```
