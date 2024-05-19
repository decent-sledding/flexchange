package config

import (
    "log"
	"os"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)



var conf = AppConfiguration{}
var logger = log.New(os.Stdout, "Config: ", log.LstdFlags)




type AppConfiguration struct {
	Database DatabaseConfig
	Server 	 ServerConfig
	Mail 	 MailConfig
}

type DatabaseConfig struct {
	User 		 string `env:"DB_USER"`
	Password 	 string `env:"DB_PASSWORD"`
	DatabaseName string `env:"DB_NAME"`
	Host 		 string `env:"DB_HOST"`
	Port 		 uint 	`env:"DB_PORT"`
}

type ServerConfig struct {
	Host string `env:"SERVER_HOST"`
	Port uint 	`env:"SERVER_PORT"`
}

type MailConfig struct {
	SmtpHost 	 string `env:"SMTP_HOST"`
	SmtpPort 	 uint  	`env:"SMTP_PORT"`
	SenderEmail  string `env:"SMTP_SENDER"`
	SmtpPassword string `env:"SMTP_PASSWORD"`
}


func GetConfig() *AppConfiguration {
	if (AppConfiguration{}) != conf {
		return &conf
	}

	logger.Print("Parsing config")

	if err := godotenv.Load(); err != nil {
        logger.Print("No .env file found")
    }

	setupDbConfig(&conf)
	setupServerConfig(&conf)
	setupMailConfig(&conf)

	return &conf
}


func setupDbConfig(c *AppConfiguration) {
	dbconf := &c.Database

	dbconf.User = "unkex"
	dbconf.Password = "kex123"
	dbconf.Host = "localhost"
	dbconf.Port = 5432
	
	if err := env.Parse(dbconf); err != nil {
		logger.Print("Could not parse Database config from .env")
	}
}

func setupServerConfig(c *AppConfiguration) {
	sconf := &c.Server

	sconf.Host = "localhost"
	sconf.Port = 9090

	if err := env.Parse(sconf); err != nil {
		logger.Print("Could not parse server config from .env")
	}
}


func setupMailConfig(c *AppConfiguration) {
	mconf := &c.Mail

	mconf.SmtpHost = "smtp.gmail.com"
	mconf.SmtpPort = 587
	mconf.SmtpPassword = ""

	if err := env.Parse(mconf); err != nil {
		logger.Print("Could not parse mail config from .env")
	}
}