package main

import (
	"CWS/api/resthandlers"
	"CWS/api/routes"
	"CWS/pb"
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
)

func init() {
	flag.IntVar(&port, "port", 9000, "api service port")
	flag.StringVar(&authAddr, "auth_addr", "localhost:9001", "authentication service address")
	flag.Parse()
}

func main() {

	conn, err := grpc.Dial(authAddr, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	autSvcClient := pb.NewAuthServiceClient(conn)
	authHandlers := resthandlers.NewAuthHandlers(autSvcClient)
	authRoutes := routes.NewAuthRoutes(authHandlers)

	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, authRoutes)

	log.Printf("API service running on [::]:%d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), routes.WithCORS(router)))
}
