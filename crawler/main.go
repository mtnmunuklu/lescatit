package main

import (
	"Lescatit/crawler/repository"
	"Lescatit/crawler/scraper"
	"Lescatit/crawler/service"
	"Lescatit/db"
	"Lescatit/pb"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// Contains some variables(port, local) for crawler service.
var (
	local bool
	port  int
)

// Init initializes the specify options for categorization service.
func init() {
	flag.IntVar(&port, "port", 9003, "crawler service port")
	flag.BoolVar(&local, "local", true, "run crawler service local")
	flag.Parse()
}

// Main starts the crawler service.
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

	crawlersRepository := repository.NewCrawlersRepository(conn)
	collyScraper := scraper.NewCollyScraper()
	crawlService := service.NewCrawlService(crawlersRepository, collyScraper)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCrawlServiceServer(grpcServer, crawlService)
	log.Printf("Crawler service running on [::]:%d\n", port)

	grpcServer.Serve(lis)
}
