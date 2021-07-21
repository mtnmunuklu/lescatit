package main

import (
	"CWS/categorization/repository"
	"CWS/categorization/service"
	"CWS/db"
	"CWS/pb"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	local bool
	port  int
)

func init() {
	flag.IntVar(&port, "port", 9002, "categorization service port")
	flag.BoolVar(&local, "local", true, "run categorization service local")
	flag.Parse()
}

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
	categoryService := service.NewCatSevice(categoriesRepository)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCatServiceServer(grpcServer, categoryService)
	log.Printf("Categorization service running on [::]:%d\n", port)

	grpcServer.Serve(lis)
}
