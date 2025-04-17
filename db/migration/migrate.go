package main

import(
	"fmt"
	"ai-pdf-chat/db/model"
	"ai-pdf-chat/config"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	config.DB.AutoMigrate(&model.User{})
	config.DB.AutoMigrate(&model.File{})
	config.DB.AutoMigrate(&model.Chat{})

	fmt.Println("Migration completed")
}