package main

import (
	"context"
	"log"
	"net/http"
	"api/controller"
	"api/httputil"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "api/docs"
)

var cfg config
var ctx = context.Background()	

type config struct {
	MongoDbConnStr        string `split_words:"true" required:"true"`
	MongoDbDatabase       string `split_words:"true" required:"true"`
	MongoDbCollectionCustomer string `split_words:"true" required:"true"`
	MongoDBCollCompany 	  string `split_words:"true" required:"true"`
	WebserverPort         string `split_words:"true" default:":8080"`
	EnableAuth            bool   `split_words:"true" default:"false"`
	SiteFilepath          string `split_words:"true" default:"/www/"`
	SwaggerHost           string `split_words:"true" default:"http://v2x.australiaeast.cloudapp.azure.com"`
	MongoDbFindLimit      int64  `split_words:"true" default:"1000"`
}

// @title Notion API
// @version 0.1
// @description Notion API documentation

// @contact.name Tisa

// @host localhost
// @BasePath /api/v1

func main() {
	var cfg config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	docs.SwaggerInfo.Host = cfg.SwaggerHost

	r := gin.Default()
	log.Println("Building API controller...")
	c, err := controller.NewController(
		cfg.MongoDbConnStr,
		cfg.MongoDbDatabase,
		cfg.MongoDbCollectionCustomer,
		cfg.MongoDBCollCompany,
		cfg.MongoDbFindLimit,
	)
	if err != nil {
		log.Fatalln("Unable to create controller with err:", err)
	}
	log.Println("Built API controller.")

	if cfg.EnableAuth {
		log.Println("API Auth is enabled.")
		r.Use(auth())
	} else {
		log.Println("API Auth is not enabled.")
	}


	// RouterGroups
	v1 := r.Group("/api/v1")
	{
		customers := v1.Group("/customers")
		{
			customers.GET(":id", c.ShowCUSTOMER)
			customers.GET("", c.ListCUSTOMERs)
			customers.DELETE(":id", c.DeleteCUSTOMER)
			customers.POST("", c.AddCUSTOMER)
			customers.DELETE("", c.DeleteCUSTOMERs)
		}
		
	}
	log.Println("cfg.SiteFilepath:", cfg.SiteFilepath)

	//Handle API documentation 
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/oauth2-redirect.html", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(static.Serve("/", static.LocalFile(cfg.SiteFilepath, true)))
	r.Run(cfg.WebserverPort)
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
