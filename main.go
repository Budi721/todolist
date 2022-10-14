package main

import (
	"fmt"
	"github.com/Budi721/todolistskyshi/app/services"
	"github.com/Budi721/todolistskyshi/business/data"
	"github.com/Budi721/todolistskyshi/business/data/activity"
	"github.com/Budi721/todolistskyshi/business/data/todo"
	"github.com/Budi721/todolistskyshi/fondation/web"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
	"os/signal"
	"runtime"

	"github.com/Budi721/todolistskyshi/business/sys/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	// perform start and shutdown sequence
	if err := run(); err != nil {
		log.Fatalf("ERROR [startup] - %s", err)
	}
}

func run() error {
	// set correct number of thread for service base cpu
	if _, err := maxprocs.Set(); err != nil {
		return fmt.Errorf("maxprocs: %w", err)
	}
	log.Printf("GOMAXPROCS [startup] %v", runtime.GOMAXPROCS(0))

	// load environment variable
	if err := godotenv.Load(); err != nil {
		return err
	}

	// init database service
	db, err := database.Open(database.Config{
		User:     "root",
		Password: "root",
		Host:     "localhost",
		Name:     "todo4",
		Port:     3306,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}

	if err := db.AutoMigrate(&activity.Activity{}, &todo.Todo{}); err != nil {
		return err
	}

	defer func() {
		log.Println("Stopping database support")
		sql, _ := db.DB()
		_ = sql.Close()
	}()

	// starting service
	app := fiber.New(fiber.Config{
		ErrorHandler: web.InternalServerErrorHandler,
	})
	factory := data.NewFactory(db)
	validate := validator.New()
	services.NewAppRouter(app, factory, validate)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":3000"); err != nil {
		return err
	}

	return nil
}
