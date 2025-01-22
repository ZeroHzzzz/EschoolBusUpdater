package constants

import (
	"EBUSU/config/config"
	"time"
)

var EBusHost = config.Config.GetString("eBus.host")
var SystemUid = config.Config.GetString("eBus.uid")
var Interval = config.Config.GetDuration("timer.interval") * time.Hour
var Port = config.Config.GetInt("server.port")
