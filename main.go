package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/NivasGuna/spotify_project/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/NivasGuna/spotify_project/dao"
	"github.com/NivasGuna/spotify_project/database"
	"github.com/NivasGuna/spotify_project/service"
)


// @title Musictrack API
// @version 1.0
// @description API for managing music tracks
// @host localhost:8080
// @BasePath /

func main(){
	router:=gin.Default()

	docs.SwaggerInfo.BasePath ="/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	ConnectDB,SpotifyAction,err:=database.Initializers()
	if err!=nil{
		log.Fatal("error in  db connection")
	}
	fmt.Print(ConnectDB)
	
	trackRepository:=dao.NewTrackRepository(ConnectDB)
 
    spotifyRepository:=dao.SpotifyRepository(SpotifyAction)

	 trackService:=service.NewTrackService(trackRepository,spotifyRepository)

	 router.GET("/getisrc/:isrc",trackService.GetISRCHandlers)
	 router.GET("/getartist/:artistname",trackService.GetArtistHandlers)
	 router.POST("/postisrc",trackService.PostISRCHandlers)
	 router.PUT("/updatetrack/:isrc",trackService.UpdateTrackHandlers)

	router.Run(":8080")
	
}