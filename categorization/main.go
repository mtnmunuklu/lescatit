package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mtnmunuklu/lescatit/categorization/repository"
	"github.com/mtnmunuklu/lescatit/categorization/service"
	"github.com/mtnmunuklu/lescatit/db"
	"github.com/mtnmunuklu/lescatit/pb"
	"github.com/mtnmunuklu/lescatit/security"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// Contains some variables(port, local) for categorization service.
var (
	port  int
	local bool
)

// Init initializes the specify options for categorization service.
func init() {
	flag.IntVar(&port, "port", 9002, "categorization service port")
	flag.BoolVar(&local, "local", true, "run categorization service local")
	flag.Parse()
}

// Main starts the categorization service.
func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	categoriesRepository := repository.NewCategoriesRepository(conn)
	catService := service.NewCatSevice(categoriesRepository)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cert_path := os.Getenv("CERT_PATH")
	tlsCredentials, err := security.LoadServerTLSCredentials(cert_path)
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	pb.RegisterCatServiceServer(grpcServer, catService)

	log.Printf("Categorization service running on [::]:%d\n", port)

	grpcServer.Serve(listen)
}
