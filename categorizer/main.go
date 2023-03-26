package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mtnmunuklu/lescatit/categorizer/classifiers"
	"github.com/mtnmunuklu/lescatit/categorizer/repositories/categorizersrps"
	"github.com/mtnmunuklu/lescatit/categorizer/repositories/classifiersrps"
	"github.com/mtnmunuklu/lescatit/categorizer/service"
	"github.com/mtnmunuklu/lescatit/categorizer/tokenizer"
	"github.com/mtnmunuklu/lescatit/db"
	"github.com/mtnmunuklu/lescatit/pb"
	"github.com/mtnmunuklu/lescatit/security"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// Contains some variables(port, local) for categorizer service.
var (
	port  int
	local bool
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

	categorizersRepository := categorizersrps.NewCategorizersRepository(conn)
	classifiersRepository := classifiersrps.NewClassifiersRepository(conn)
	tokenizer := tokenizer.NewTokenizer()
	nbCategorizer := classifiers.NewNaiveBayesianClassifier()
	catzeService := service.NewCatzeService(categorizersRepository, classifiersRepository, tokenizer, nbCategorizer)

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
	pb.RegisterCatzeServiceServer(grpcServer, catzeService)

	log.Printf("Categorizer service running on [::]:%d\n", port)

	grpcServer.Serve(listen)
}
