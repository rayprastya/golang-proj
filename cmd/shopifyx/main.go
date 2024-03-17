package main

import (
	"database/sql"
	"log"

	"project/shopifyx/internal/api"
	"project/shopifyx/internal/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "your_connection_string")
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	albumRepo := repository.NewAlbumRepository(db)
	albumHandler := api.NewAlbumHandler(albumRepo)

	router := gin.Default()
	router.GET("/albums", albumHandler.GetAlbums)
	router.POST("/albums", albumHandler.PostAlbums)

	router.Run("localhost:3001")
}
