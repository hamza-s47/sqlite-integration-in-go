package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hamza-s47/sqlite-integration-in-go/database"
	"github.com/hamza-s47/sqlite-integration-in-go/lead"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *gin.Engine) {
	app.GET("/api/v1/lead", lead.GetLeads)
	app.GET("/api/v1/lead/:id", lead.GetLead)
	app.POST("/api/v1/lead", lead.NewLead)
	app.DELETE("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "./leads.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	log.Println("Connection opened to Database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	log.Println("Database Migrated")
}

func main() {
	app := gin.Default()
	initDatabase()
	setupRoutes(app)
	app.Run(":8080")
	defer database.DBConn.Close()
}
