package database

import (
	"context"
	"log"

	"github.com/NivasGuna/spotify_project/models"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initializers() (*gorm.DB,*spotify.Client,error){
	DB,err:= gorm.Open(postgres.Open("host=otto.db.elephantsql.com user=fpxmkeia password=4La7MjYeEuTDPttPLpxaj-NTbAfBrQxr dbname=fpxmkeia  port=5432 sslmode=disable "), &gorm.Config{});
	if err != nil {
		log.Fatal("connection error")
	}


	if err := DB.AutoMigrate(&models.Musictrack{}); err != nil {
		return nil,nil,err
	}
	datas:=authenticateSpotify()
	return DB,datas,nil
}
var SpotifyCredentials = struct {
	ClientID     string
	ClientSecret string
}{
	ClientID:     "d0a672cffb8149f685b4002ec1bc9da3",
	ClientSecret: "0795acb1e47345138c00c462374d2158",
}

func authenticateSpotify() *spotify.Client {
	config := &clientcredentials.Config{
		ClientID:     SpotifyCredentials.ClientID,
		ClientSecret: SpotifyCredentials.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		log.Printf("Failed to get Spotify token: %v", err)
		return nil
	}

	client := spotify.Authenticator{}.NewClient(token)
	return &client
}