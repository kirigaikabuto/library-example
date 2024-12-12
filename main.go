package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kirigaikabuto/library-example/books"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	version              = "0.0.1"
	postgresUser         = ""
	postgresPassword     = ""
	postgresDatabaseName = ""
	postgresHost         = ""
	postgresPort         = 5432
	postgresParams       = ""
	port                 = "8080"
)

func parseEnvFile() {
	postgresUser = os.Getenv("POSTGRES_USER")
	postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	postgresDatabaseName = os.Getenv("POSTGRES_DB")
	postgresParams = os.Getenv("POSTGRES_PARAM")
	postgresPortStr := os.Getenv("POSTGRES_PORT")
	postgresPort, _ = strconv.Atoi(postgresPortStr)
	postgresHost = os.Getenv("POSTGRES_HOST")
	port = os.Getenv("PORT")
}

func run(c *cli.Context) error {
	parseEnvFile()
	cfg := books.PostgresConfig{
		Host:     postgresHost,
		Port:     postgresPort,
		User:     postgresUser,
		Password: postgresPassword,
		Database: postgresDatabaseName,
		Params:   postgresParams,
	}
	booksPostgresStore, err := books.NewPostgresStore(cfg)
	if err != nil {
		return err
	}

	booksService := books.NewService(booksPostgresStore)
	booksHttpEndpoints := books.NewHttpEndpoints(setdata_common.NewCommandHandler(booksService))

	r := gin.Default()

	booksGroup := r.Group("/books")
	{
		booksGroup.GET("/", booksHttpEndpoints.MakeList())
		booksGroup.POST("/", booksHttpEndpoints.MakeCreate())
		booksGroup.GET("/id/", booksHttpEndpoints.MakeGetById())
	}

	log.Info().Msg("app is running on port:" + port)
	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Server ListenAndServe error")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting.")
	return nil
}

func main() {
	app := &cli.App{
		Name:    "libary api",
		Version: version,
		Usage:   "library api",
		Action:  run,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}
}
