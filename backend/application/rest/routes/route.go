package routes

import (
	rest "github.com/EusRique/frajola/application/rest/handlers"
	"github.com/EusRique/frajola/application/usecase"
	"github.com/EusRique/frajola/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func configRoutes(api *gin.RouterGroup) {
	document := rest.NewUsers(usecase.DocumentUseCase{})

	api.GET("/", rest.Alive)
	api.GET("/alive", rest.Alive)

	api.POST("/create-document", document.CreateDocument)
	api.GET("/list-documents", document.ListAllDocuments)
	api.DELETE("/delete-document/:document_id", document.DeleteDocument)
	api.PUT("/update-document/:document_id", document.UpdateDocument)
}

func Start(port string) {
	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	defaultMiddlewares := []gin.HandlerFunc{middleware.Database()}

	version := "v1"
	api := r.Group(version+"/api", defaultMiddlewares...)

	configRoutes(api)
	r.Run(":" + port)
}
