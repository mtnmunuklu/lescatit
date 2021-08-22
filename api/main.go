package main

import (
	"Lescatit/api/resthandlers"
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
	port     int
	authAddr string
	catAddr  string
)

func init() {
	flag.IntVar(&port, "port", 9000, "api service port")
	flag.StringVar(&authAddr, "auth_addr", "localhost:9001", "authentication service address")
	flag.StringVar(&catAddr, "cat_addr", "localhost:9002", "categorization service address")
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
	authHandlers := resthandlers.NewAuthHandlers(autSvcClient)
	authRoutes := routes.NewAuthRoutes(authHandlers)

	// for categorization service
	catConn, err := grpc.Dial(catAddr, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer authConn.Close()

	catSvcClient := pb.NewCatServiceClient(catConn)
	catHandlers := resthandlers.NewCatHandlers(catSvcClient)
	catRoutes := routes.NewCatRoutes(catHandlers)

	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, authRoutes)
	routes.Install(router, catRoutes)

	log.Printf("API service running on [::]:%d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), routes.WithCORS(router)))
}
