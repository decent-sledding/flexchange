package config

import (
	_ "os"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)


type SetOpt func(*AppConfiguration)


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



func defaultConfig() *AppConfiguration {
	conf := AppConfiguration {}

	setupDbConfig(&conf)
	setupServerConfig(&conf)
	setupMailConfig(&conf)

	return &conf
}


func NewConfig(opts ...SetOpt) *AppConfiguration {
	godotenv.Load()
	conf := defaultConfig()

	for _, setopt := range opts {
		setopt(conf)
	}

	return conf
}


func setupDbConfig(c *AppConfiguration) {
	dbconf := &c.Database

	dbconf.User = "unkex"
	dbconf.Password = "kex123"
	dbconf.Host = "localhost"
	dbconf.Port = 5432
	
	_ = env.Parse(dbconf)
}


func setupServerConfig(c *AppConfiguration) {
	sconf := &c.Server

	sconf.Host = "localhost"
	sconf.Port = 9090

	_ = env.Parse(sconf)
}


func setupMailConfig(c *AppConfiguration) {
	mconf := &c.Mail

	mconf.SmtpHost = "smtp.gmail.com"
	mconf.SmtpPort = 587
	mconf.SmtpPassword = ""

	_ = env.Parse(mconf)
}