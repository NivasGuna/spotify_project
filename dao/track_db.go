package dao

import (
	"errors"

	"github.com/NivasGuna/spotify_project/models"
	"gorm.io/gorm"
)

type TrackRepository struct{
	db *gorm.DB
}

func NewTrackRepository(db *gorm.DB)  *TrackRepository{
	return &TrackRepository{db}
}

func (tr *TrackRepository) GetISRC(isrc string) (*models.Musictrack,error){
	var musictrack models.Musictrack

	result:=tr.db.Where("isrc = ?",isrc).First(&musictrack)
	
	if  result.Error !=nil{
		return &musictrack, result.Error
	}
	
	if &musictrack == nil {
		return nil, errors.New("musictrack is nil")
	}
	 
  return &musictrack,nil
}

func (tr *TrackRepository) GetArtist(ArtistName string) (*[]models.Musictrack,error){
	var musictrack []models.Musictrack

	result:=tr.db.Where("artist_name LIKE ?", "%"+ArtistName+"%").Find(&musictrack)
	if result.Error !=nil{
		return nil, result.Error
	}

	if len(musictrack) <= 0 {
           return nil,nil
	}
	
		return &musictrack,nil
	 
  
}

func (tr *TrackRepository) SaveTrack(track *models.Musictrack) error {
	return tr.db.Create(track).Error
}

func (ur *TrackRepository) UpdateTrack(user *models.Musictrack) error {
	return ur.db.Save(user).Error
}

