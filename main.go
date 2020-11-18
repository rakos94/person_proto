package main

import (
	"person_proto/config"
	"person_proto/dao"
	"person_proto/server"

	"gorm.io/gorm"
)

func main() {
	//init db and inject to dao and service
	g := initDb()
	dao.SetDao(g)

	//init proto
	server.SetProto()
}

func initDb() *gorm.DB {
	g, err := config.Conn()
	if err == nil {
		// config.MigrateSchema(g)
		return g
	}
	panic(err)
}
