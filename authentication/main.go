package main

import (
	"Lescatit/authentication/repository"
	"Lescatit/authentication/service"
	"Lescatit/db"
	"Lescatit/pb"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// Contains some variables(port, local) for authentication service.
var (
	local bool
	port  int
)

// Init initializes the specify options for authentication service.
func init() {
	flag.IntVar(&port, "port", 9001, "authentication service port")
	flag.BoolVar(&local, "local", true, "run authentication service local")
	flag.Parse()
}

// Main starts the authentication service.
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

	usersRepository := repository.NewUsersRepository(conn)
	authService := service.NewAuthService(usersRepository)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)

	log.Printf("Authentication service running on [::]:%d\n", port)

	grpcServer.Serve(lis)
}
