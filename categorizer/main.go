package main

import (
	"Lescatit/categorizer/repository"
	"Lescatit/categorizer/service"
	"Lescatit/db"
	"Lescatit/pb"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// Contains some variables(port, local) for categorizer service.
var (
	local bool
	port  int
)

// Init initializes the specify options for categorizer service.
func init() {
	flag.IntVar(&port, "port", 9004, "categorizer service port")
	flag.BoolVar(&local, "local", true, "run categorizer service local")
	flag.Parse()
}

// Main starts the categorizer service
func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Println(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	categorizersRepository := repository.NewCategorizersRepository(conn)
	catzeService := service.NewCatzeService(categorizersRepository)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCatzeServiceServer(grpcServer, catzeService)
	log.Printf("Categorizer service running on [::]:%d\n", port)

	grpcServer.Serve(lis)
}
