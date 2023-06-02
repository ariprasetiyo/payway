package main

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"payway/app/main/layer/handler"
	httprequest "payway/app/main/layer/http_request"
	"payway/app/main/layer/repository"
	"payway/app/main/layer/server"
)

var db = make(map[string]string)

func main() {

	server.InitConfig()
	server.InitLogrusFormat()

	// running open telemetry
	cleanup := server.InitTracer()
	defer cleanup(context.Background())

	gin.SetMode(server.GIN_MODE)
	engine := gin.New()
	engine.Use(server.RequestResponseLogger())
	engine.Use(otelgin.Middleware(server.APP_NAME))
	engine.Use(cors.New(server.CorsConfig()))

	initPosgresSQL := server.InitPostgreSQL()

	repoDatabase := repository.NewDatabase(initPosgresSQL.DB, initPosgresSQL)
	httpClientCerberus := httprequest.NewCerberusHttpClient()
	httpClientTokoPublish := httprequest.NewTokoPublishttpClient()

	cerberusAuthAPIHandler := handler.NewCerberusAuthorization(httpClientCerberus)
	createReciptHandler := handler.NewCreateReciptHandler(httpClientTokoPublish, &repoDatabase)
	sampleAnotherHandler := handler.NewSampleAnotherHandler()

	engine.GET("/health", handler.HealthCheck)

	authorized := engine.Group("/", cerberusAuthAPIHandler.Execute())
	authorized.POST("/public/api/v1/createReceipt/:merchantId", createReciptHandler.Execute)
	authorized.GET("/public/api/v1/sample/another/handler", sampleAnotherHandler.Execute)

	engine.Run(":" + server.HTTP_SERVER_PORT)
}
