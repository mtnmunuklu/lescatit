package main

import (
	"Lescatit/api/handlers"
	"Lescatit/api/routes"
	"Lescatit/pb"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	port      int
	authAddr  string
	catAddr   string
	crawlAddr string
	catzeAddr string
)

func init() {
	flag.IntVar(&port, "port", 9000, "api service port")
	flag.StringVar(&authAddr, "auth_addr", "localhost:9001", "authentication service address")
	flag.StringVar(&catAddr, "cat_addr", "localhost:9002", "categorization service address")
	flag.StringVar(&crawlAddr, "crawl_addr", "localhost:9003", "crawler service address")
	flag.StringVar(&catzeAddr, "catze_addr", "localhost:9004", "categorizer service address")
	flag.Parse()
}

func main() {
	// for authentication service
	authConn, err := grpc.Dial(authAddr, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer authConn.Close()

	autSvcClient := pb.NewAuthServiceClient(authConn)
	authHandlers := handlers.NewAuthHandlers(autSvcClient)
	authRoutes := routes.NewAuthRoutes(authHandlers)

	// for crawler service
	crawlConn, err := grpc.Dial(crawlAddr, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer authConn.Close()

	crawlSvcClient := pb.NewCrawlServiceClient(crawlConn)
	crawlHandlers := handlers.NewCrawlHandlers(crawlSvcClient)
	crawlRoutes := routes.NewCrawlRoutes(crawlHandlers)

	// for categorizer service
	catzeConn, err := grpc.Dial(catzeAddr, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer authConn.Close()

	catzeSvcClient := pb.NewCatzeServiceClient(catzeConn)
	catzeHandlers := handlers.NewCatzeHandlers(autSvcClient, catzeSvcClient)
	catzeRoutes := routes.NewCatzeRoutes(catzeHandlers)

	// for categorization service
	catConn, err := grpc.Dial(catAddr, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer catConn.Close()

	catSvcClient := pb.NewCatServiceClient(catConn)
	catHandlers := handlers.NewCatHandlers(crawlSvcClient, catzeSvcClient, catSvcClient)
	catRoutes := routes.NewCatRoutes(catHandlers)

	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, authRoutes)
	routes.Install(router, crawlRoutes)
	routes.Install(router, catzeRoutes)
	routes.Install(router, catRoutes)

	log.Printf("API service running on [::]:%d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), routes.WithCORS(router)))
}
