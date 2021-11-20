package main

import (
	"Lescatit/categorization/repository"
	"Lescatit/categorization/service"
	"Lescatit/db"
	"Lescatit/pb"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
)

// Contains some variables(port, local) for categorization service.
var (
	local bool
	port  int
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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Application Layer Transport Security (ALTS) is a mutual authentication and transport encryption system.
	altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
	grpcServer := grpc.NewServer(grpc.Creds(altsTC))
	pb.RegisterCatServiceServer(grpcServer, catService)
	log.Printf("Categorization service running on [::]:%d\n", port)

	grpcServer.Serve(lis)
}
