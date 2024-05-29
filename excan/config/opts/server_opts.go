package opts 

import (
	c "excan/config"
)



func SetServerHost(host string) c.SetOpt {
	return func(conf *c.AppConfiguration) {
		conf.Server.Host = host
	}
}