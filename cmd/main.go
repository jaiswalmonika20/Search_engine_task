package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/module_page/cmd/config"
	"github.com/module_page/pkg/controllers"

	"github.com/module_page/pkg/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server          *gin.Engine
	pageservice     services.PageService
	pagecontroller  controllers.PageController
	ctx             context.Context
	pagescollection *mongo.Collection
	mongoclient     *mongo.Client
	err             error

	//ulrs            string = ":8070"
	//mongo_uri string = "mongodb://localhost:27017"

)

func init() {
	ctx = context.TODO()
	config.LoadConfig()
	mongo_uri := config.Config.Database.Protocol + "://" + config.Config.Database.Host + ":" + fmt.Sprint(config.Config.Database.Port)
	mongoconn := options.Client().ApplyURI(mongo_uri)
	pagec := DatabaseConnectionSetup(mongoconn)
	pageservice = services.NewPageService(pagec, ctx)
	pagecontroller = controllers.New(pageservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	//versioning api
	basepath := server.Group(config.Config.Server.Version)
	pagecontroller.Routes(basepath)
	log.Fatal(server.Run(":" + fmt.Sprint(config.Config.Server.Port)))

}

func DatabaseConnectionSetup(mongoconnection *options.ClientOptions) *mongo.Collection {
	mongoclient, err = mongo.Connect(ctx, mongoconnection)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	//fmt.Println("mongo connection established")

	pagescollection = mongoclient.Database(config.Config.Database.DBName).Collection(config.Config.Database.Collection)
	return pagescollection
}
