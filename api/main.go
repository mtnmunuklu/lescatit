package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mtnmunuklu/lescatit/api/handlers"
	"github.com/mtnmunuklu/lescatit/api/routes"
	"github.com/mtnmunuklu/lescatit/pb"
	"github.com/mtnmunuklu/lescatit/security"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port      int
	local     bool
	authAddr  string
	catAddr   string
	crawlAddr string
	catzeAddr string
)

func init() {
	flag.BoolVar(&local, "local", true, "run api service local")
	flag.IntVar(&port, "port", 9000, "api service port")
	flag.StringVar(&authAddr, "auth_addr", "localhost:9001", "authentication service address")
	flag.StringVar(&catAddr, "cat_addr", "localhost:9002", "categorization service address")
	flag.StringVar(&crawlAddr, "crawl_addr", "localhost:9003", "crawler service address")
	flag.StringVar(&catzeAddr, "catze_addr", "localhost:9004", "categorizer service address")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}

	// for secure communication
	cert_path := os.Getenv("CERT_PATH")
	tlsCredentials, err := security.LoadCATLSCredentials(cert_path)
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// for authentication service
	authConn, err := grpc.Dial(authAddr, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Panicln(err)
	}
	defer authConn.Close()

	autSvcClient := pb.NewAuthServiceClient(authConn)
	authHandlers := handlers.NewAuthHandlers(autSvcClient)
	authRoutes := routes.NewAuthRoutes(authHandlers)

	// for crawler service
	crawlConn, err := grpc.Dial(crawlAddr, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Panicln(err)
	}
	defer authConn.Close()

	crawlSvcClient := pb.NewCrawlServiceClient(crawlConn)
	crawlHandlers := handlers.NewCrawlHandlers(crawlSvcClient)
	crawlRoutes := routes.NewCrawlRoutes(crawlHandlers)

	// for categorizer service
	catzeConn, err := grpc.Dial(catzeAddr, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Panicln(err)
	}
	defer authConn.Close()

	catzeSvcClient := pb.NewCatzeServiceClient(catzeConn)
	catzeHandlers := handlers.NewCatzeHandlers(autSvcClient, catzeSvcClient)
	catzeRoutes := routes.NewCatzeRoutes(catzeHandlers)

	// for categorization service
	catConn, err := grpc.Dial(catAddr, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Panicln(err)
	}
	defer catConn.Close()

	catSvcClient := pb.NewCatServiceClient(catConn)
	catHandlers := handlers.NewCatHandlers(autSvcClient, crawlSvcClient, catzeSvcClient, catSvcClient)
	catRoutes := routes.NewCatRoutes(catHandlers)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	routes.Install(app, authRoutes)
	routes.Install(app, crawlRoutes)
	routes.Install(app, catzeRoutes)
	routes.Install(app, catRoutes)

	log.Printf("API service running on [::]:%d\n", port)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
