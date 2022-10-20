package main

import (
	"fmt"
	"github.com/Budi721/todolistskyshi/app/services"
	"github.com/Budi721/todolistskyshi/business/data"
	"github.com/Budi721/todolistskyshi/business/data/activity"
	"github.com/Budi721/todolistskyshi/business/data/todo"
	"github.com/Budi721/todolistskyshi/business/sys/database"
	"github.com/Budi721/todolistskyshi/fondation/web"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/automaxprocs/maxprocs"
	"log"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"strconv"
	"strings"
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

	port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	db, err := database.Open(database.Config{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Name:     os.Getenv("MYSQL_DBNAME"),
		Port:     port,
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
	en := en.New()
	uni := ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")
	err = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} cannot be null", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return ToSnakeCase(t)
	})
	if err != nil {
		return err
	}
	services.NewAppRouter(app, factory, validate, trans)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":3030"); err != nil {
		return err
	}

	return nil
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
