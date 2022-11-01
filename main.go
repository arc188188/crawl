package main

import (
	"awesomeProject/biz/crawl"
	"awesomeProject/db"
)

var Db = db.Conn()

func main() {
	// Auto Migrate
	// Set table options
	//Db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&User{})
	// Set table cluster options
	//Db.Set("gorm:table_cluster_options", "on cluster default").AutoMigrate(&User{})
	crawl.Actress()
}
