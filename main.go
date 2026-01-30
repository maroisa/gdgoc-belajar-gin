package main

import (
	"log"

	"belajar-gin/db"
	"belajar-gin/server/note"
	"belajar-gin/db/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("gagal memuat .env")
	}

	DB := db.NewConnection()
	r := gin.Default()

	DB.AutoMigrate(&model.Note{})

	note.NewHandler(r, DB)

	r.Run()
}
