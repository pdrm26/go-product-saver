package migrate

import (
	"fmt"
	"log"
	"my-module/initializers"
	"my-module/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.Product{}, &models.Brand{}, &models.Category{})
	fmt.Println("👍 Migration complete")
}
