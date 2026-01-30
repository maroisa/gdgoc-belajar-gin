package main

import (
	"log"

	"belajar-gin/db"
	"belajar-gin/server/note"

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

	note.NewHandler(r, DB)

	r.Run()
}
