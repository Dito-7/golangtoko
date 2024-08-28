package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Dito-7/golangtoko/app/models"
	"github.com/Dito-7/golangtoko/database/seeders"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (s *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome To " + appConfig.AppName)

	s.initializeDB(dbConfig)
	s.initializeRoutes()
}

func (s *Server) Run(addr string) {
	fmt.Printf("Server is listening on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}

func (s *Server) initializeDB(dbConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	s.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

// migration
func (s *Server) dbMigrate() {
	var err error

	for _, model := range models.RegisterModels() {
		err = s.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database successfully migrated")
}

func (s *Server) InitCommands(appConfig AppConfig, dbConfig DBConfig) {
	s.initializeDB(dbConfig)
	cmdApp := cli.NewApp()
	cmdApp.Commands = []*cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				s.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.RunSeeders(s.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
