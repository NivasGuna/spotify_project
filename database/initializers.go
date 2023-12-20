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
	DB,err:= gorm.Open(postgres.Open("postgres://your_username:your_password@localhost:5432/your_database"), &gorm.Config{});
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
	ClientID:     "Your_client_ID",
	ClientSecret: "Your_client_Secret",
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
