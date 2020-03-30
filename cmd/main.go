package main

import (
	"log"

	"github.com/zaininfo/example-project/pkg"
)

var dbConn = pkg.NewDB()

func main() {
	//go func() {
	//	for {
	//		select {
	//		case <-time.After(time.Second * 10):
	//			reportHealth()
	//		}
	//	}
	//}()

	pkg.StartMonitoring()
}

func reportHealth() {
	pong, err := dbConn.Ping()
	if err != nil {
		pkg.ServiceStatus.WithLabelValues(pkg.AppName).Set(0)

		log.Printf("DB is down: %s", err)
		return
	}

	pkg.ServiceStatus.WithLabelValues(pkg.AppName).Set(1)

	log.Printf("DB is up: %s", pong)
}
