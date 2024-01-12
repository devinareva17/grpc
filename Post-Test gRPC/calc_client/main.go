package main

import (
	"context"
	"log"
	"os"
	"strings"
<<<<<<< HEAD
	"strconv"
=======
>>>>>>> f8b3a71fd43ec00324de7b4b80251a4c846e682b
	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
<<<<<<< HEAD

)


var address = flag.String("server", "localhost:50053", "server address")	

func argParser(n1 string, n2 string) (float32, float32) {
	N1, err := strconv.ParseFloat(os.Args[1], 32)
=======
	pb "github.com/devinareva17/grpc/calculator/calculatorpb"

)

var address = flag.String("server", "localhost:50053", "server address")

//Parsing input n1 dan n2 
func argParser(n1, n2 string) (int32, int32) {
	N1, err := strconv.Atoi(n1)
>>>>>>> f8b3a71fd43ec00324de7b4b80251a4c846e682b
	if err != nil {
		log.Fatalf("Cannot parse arg[1]: arguments should be in float32 and in range")
	}
	N2, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		log.Fatalf("Cannot parse arg[2]: arguments should be float32 and in range")
	}
	return float32(N1), float32(N2)
}

func main() {
	//Menghubungkan ke server gRPC
	flag.Parse()

	//Cek input hanya 2 angka (angka1 dan angka2)
	if len(os.Args) != 3 {
		log.Fatalf("only 2 numbers expected: n1 n2")
	}
<<<<<<< HEAD

	n1,n2 := argParser(os.Args[1], os.Args[2])
=======
	//Parsing n1 dan n2
	n1, n2 := argParser(os.Args[1], os.Args[2])
>>>>>>> f8b3a71fd43ec00324de7b4b80251a4c846e682b
	
	// Set up koneksi ke server
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect: %s", err)
	}
	defer conn.Close()

	// Membuat instance dari Calculator Service Client
	client := pb.NewCalculatorClient(conn)

	//Memanggil fungsi Add dari server
<<<<<<< HEAD
	addResponse, err := client.Add(context.Background(), &pb.Request{N1:float32(n1), N2:float32(n2)})
	if err != nil {
		log.Fatalf("Adding error: %s", err)
	}
	log.Printf("%.2f + %.2f = %.2f", n1, n2, addResponse.N1)
	
	//Memanggil fungsi Substract dari server
	subtractResponse, err := client.Subtract(context.Background(), &pb.Request{N1:float32(n1), N2:float32(n2)})
=======
	addResult, err := client.Add(context.Background(), &pb.AddRequest{N1: int32(n1), N2: int32(n2)})
	if err != nil {
		log.Fatalf("Adding error: %s", err)
	}
	log.Printf("%d + %d = %d", n1, n2, addResult.N1)
	
	//Memanggil fungsi Substract dari server
	subtractResult, err := client.Subtract(context.Background(), &pb.SubtractRequest{N1: n1, N2: n2})
>>>>>>> f8b3a71fd43ec00324de7b4b80251a4c846e682b
	if err != nil {
		log.Fatalf("Subtracting error: %s", err)
	}
	log.Printf("%.2f + %.2f = %.2f", n1, n2, subtractResponse.N1)

	//Memanggil fungsi Multiply dari server
<<<<<<< HEAD
	multiplyResponse, err := client.Multiply(context.Background(), &pb.Request{N1:float32(n1), N2:float32(n2)})
	if err != nil {
		log.Fatalf("Multiplying error: %s", err)
	}
	log.Printf("%.2f + %.2f = %.2f", n1, n2, multiplyResponse.N1)
	
	//Memanggil fungsi Divide dari server
	divideResponse, err := client.Divide(context.Background(), &pb.Request{N1:float32(n1), N2:float32(n2)})
	if err != nil {
		// Cek apakah N2 == 0 (tidak bisa pembagian dengan 0)
=======
	multiplyResult, err := client.Multiply(context.Background(), &pb.MultiplyRequest{N1: n1, N2: n2})
	if err != nil {
		log.Fatalf("Multiplying error: %s", err)
	}
	log.Printf("%d * %d = %d", n1, n2, multiplyResult.N1)
	
	//Memanggil fungsi Divide dari server
	divideResult, err := client.Divide(context.Background(), &pb.DivideRequest{N1: n1, N2: n2})
	if err != nil {
		// Cek apakah N2 == 0
>>>>>>> f8b3a71fd43ec00324de7b4b80251a4c846e682b
		if strings.Contains(err.Error(), "Cannot divide by zero") {
			log.Printf("Dividing error: Cannot divide by zero")
		} else {
			log.Fatalf("Dividing error: %s", err)
		}
	} else {
		log.Printf("%.2f + %.2f = %.2f", n1, n2, divideResponse.N1)
	}

}
