package middleware

import (
	"fmt"
	"os"

	"github.com/EusRique/frajola/infrastructure/db"
	"github.com/gin-gonic/gin"
)

func Database() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("[middleware] Database()")
		database := db.ConnectDB(os.Getenv("env"))
		sqlDB := database.DB()
		defer sqlDB.Close()
		c.Set("DB", database)
		c.Next()
	}
}
