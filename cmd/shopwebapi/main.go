package main

import (
	"fiber-shop/internal/config"
	"fiber-shop/internal/routes"
	"os"

	"fiber-shop/pkg/pgsql"
	"fiber-shop/pkg/utils"
	"log"
	"time"

	_ "fiber-shop/docs"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title Fiber Shop API
// @version 1.0
// @description This is a simple (yet) Shop api build using Fiber and Go.
func main() {
	//TODO make config load from env once before app started

	log.Println("Initializing logger")

	l := log.New(os.Stdout, "go-fiber-store: ", log.LstdFlags)

	l.Println("Logger initialized")

	l.Println("Defining fiber app")

	a := fiber.New()

	l.Println("App defined")

	pgCfg := pgsql.NewConfig("postgres", "Qweasdzxc4", "192.168.1.150", "5432", "store")

	pgsql.NewGlClient(1, 1*time.Second, pgCfg, l)

	defer pgsql.ConnPool.Close()

	testConn, err := pgsql.ConnPool.Acquire()

	if err != nil {
		l.Fatal(err)
	}

	l.Println(testConn.Query("SELECT * FROM public.user"))


	l.Println("Middleware initializing")

	a.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	
	a.Use(cors.New())

	//TODO middlewares
	//TODO auth middleware

	l.Println("Middleware initialized")

	l.Println("Routes initializing")

	routes.Swagger(a)
	routes.Metrics(a)
	routes.User(a)

	// TODO routes
	l.Println("Routes initialized")

	l.Println("Starting server")

	ip := config.GetConfigValue("BIND_IP")

	port := config.GetConfigValue("PORT")

	url := ip + ":" + port

	utils.StartServer(a, url, l)

}

