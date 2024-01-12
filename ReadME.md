# OUTLINE
1. How to Set Up gRPC Project
2. How to Run this project
3. Test Case

# How to Set Up gRPC Project


## Golang Installation

Make sure that Golang is installed on your computer. if not, please [install](https://go.dev/doc/install) first.

check golang version to make sure golang is installed

```bash
go version
```

## Proto Compiler Installation

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

create a folder to store the proto and create a .proto file
```
project
│└───calculator
│   └───calculatorpb
│       │   calculator.proto
│   
└───go.mod
└───go.sum
```


install the following modules.
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

run proto file
```bash
protoc --go_out=. --go-grpc_out=. calculator/calculatorpb/calculator.proto
```

## Set up Client-Server

create a folder to create the main program on the client and server
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
Now, you can start to make your client and server code

# How to Run This Project

## 1. Clone Repository 

clone this repository with following command
```bash
git clone https://github.com/devinareva17/grpc.git
```

## 2. Run Server

run server main.go with this following command
```bash
go run calc_server/main.go
```

## 2. Run Client
run client main.go with adding 2 numbers to calculate
```bash
go run calc_client/main.go 10 2
```

# Test Case

## 1. n1 = 1000; n2 = 23
```bash
go run calc_client/main.go 1000 23
```
client output
```bash
2024/01/12 21:48:47 1000.00 + 23.00 = 1023.00
2024/01/12 21:48:47 1000.00 - 23.00 = 977.00
2024/01/12 21:48:47 1000.00 * 23.00 = 23000.00
2024/01/12 21:48:47 1000.00 / 23.00 = 43.48
```
server output
```bash
listening on port 50053...
2024/01/12 21:48:47 Received request with N1=1000.00 and N2=23.00
```

## 2. n1 = 23; n2 = 1000
```bash
go run calc_client/main.go 23 1000
```

Client Output
```bash
2024/01/12 21:51:51 23.00 + 1000.00 = 1023.00
2024/01/12 21:51:51 23.00 - 1000.00 = -977.00
2024/01/12 21:51:51 23.00 * 1000.00 = 23000.00
2024/01/12 21:51:51 23.00 / 1000.00 = 0.02
```
server output
```bash
listening on port 50053...
2024/01/12 21:48:47 Received request with N1=1000.00 and N2=23.00
2024/01/12 21:51:51 Received request with N1=23.00 and N2=1000.00
```

## 3. n1 = 93.7; n2 = 4.5
```bash
go run calc_client/main.go 93.7 4.5
```

Client Output
```bash
2024/01/12 21:54:00 93.70 + 4.50 = 98.20
2024/01/12 21:54:00 93.70 - 4.50 = 89.20
2024/01/12 21:54:00 93.70 * 4.50 = 421.65
2024/01/12 21:54:00 93.70 / 4.50 = 20.82
```
server output
```bash
listening on port 50053...
2024/01/12 21:48:47 Received request with N1=1000.00 and N2=23.00
2024/01/12 21:51:51 Received request with N1=23.00 and N2=1000.00
2024/01/12 21:54:00 Received request with N1=93.70 and N2=4.50
```

## 4. n1 = 10; n2 = 0
```bash
go run calc_client/main.go 10 0
```
Client Output
```bash
2024/01/12 22:00:06 10.00 + 0.00 = 10.00
2024/01/12 22:00:06 10.00 - 0.00 = 10.00
2024/01/12 22:00:06 10.00 * 0.00 = 0.00
2024/01/12 22:00:06 Dividing error: Cannot divide by zero
```
server output
```bash
listening on port 50053...
2024/01/12 21:48:47 Received request with N1=1000.00 and N2=23.00
2024/01/12 21:51:51 Received request with N1=23.00 and N2=1000.00
2024/01/12 21:54:00 Received request with N1=93.70 and N2=4.50
2024/01/12 22:00:06 Received request with N1=10.00 and N2=0.00
```
