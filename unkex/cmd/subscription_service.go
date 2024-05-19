package main

import (
	"fmt"

	"unkex/config"
	"unkex/internal/mail_api"
)



func main() {
	conf := config.GetConfig()

	cr := cron.New()
	cr.AddFunc("0 0 0 * * *", main_api.handleSubscriptionNewsletter(&conf))
	cr.Start()
}