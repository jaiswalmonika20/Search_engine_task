package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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
	ulrs            string = ":8080"
	mongo_uri       string = "mongodb://mongo-container:27017"
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI(mongo_uri)
	pagec := DatabaseConnectionSetup(mongoconn)
	pageservice = services.NewPageService(pagec, ctx)
	pagecontroller = controllers.New(pageservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	//versioning api
	basepath := server.Group("/v1")
	pagecontroller.Routes(basepath)

	log.Fatal(server.Run(ulrs))

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

	fmt.Println("mongo connection established")

	pagescollection = mongoclient.Database("new").Collection("page")
	return pagescollection
}
