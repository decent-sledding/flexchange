package main

import (
	"fmt"

	"excan/config"
	"excan/internal/mail_api"
)



func main() {
	conf := config.GetConfig()

	err = storage.RunMigrations(connectionString)

	if err != nil {
		return err
	}

	cr := cron.New()
	cr.AddFunc("0 0 0 * * *", main_api.handleSubscriptionNewsletter(&conf))
	cr.Start()
}