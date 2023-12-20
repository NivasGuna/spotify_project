package service

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	// "google.golang.org/genproto/googleapis/rpc/status"
	"github.com/NivasGuna/spotify_project/dao"
	"github.com/NivasGuna/spotify_project/models"
)

type TrackService struct {
	trackRepository   *dao.TrackRepository
	spotifyRepository *dao.SpotifyClient
}

func NewTrackService(trackRepository *dao.TrackRepository, spotifyRepository *dao.SpotifyClient) *TrackService {
	return &TrackService{trackRepository, spotifyRepository}
}

// @Summary Get track by ISRC or search and store from Spotify
// @Description Get a track by its ISRC or search and store from Spotify if not found
// @ID get-track-by-isrc
// @Accept json
// @Produce json
// @Param isrc path string true "ISRC of the track"
// @Success 200 {object} models.Musictrack "Successful response"
// @Failure 404 {object} map[string]interface{} "Track not found"
// @Failure 500 {object} map[string]interface{} "Server error"
// @Router /getisrc/{isrc} [get]
func (us *TrackService) GetISRCHandlers(c *gin.Context) {
	isrc := c.Param("isrc")

	searchTrack, err := us.trackRepository.GetISRC(isrc)

	if searchTrack != nil && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": searchTrack,
		},
		)
	}else{
	track, _ := us.spotifyRepository.SearchTrackByISRC(isrc)

	 // fmt.Println(track)
	 if track ==nil{
		c.JSON(400, gin.H{
			"message": "unable to get the track by your isrc",
		})
	 }
	if  track !=nil  {
		Track := &models.Musictrack{
			ISRC:         isrc,
			Title:        track.Name,
			ArtistName:   track.Artists[0].Name,
			SpotifyImage: track.Album.Images[0].URL,
		}
		err = us.trackRepository.SaveTrack(Track)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "unable to store the track",
			})
		} else {
			c.JSON(http.StatusOK, Track)
		}
	}
	}
}

// ErrorResponse represents the structure of an error response
// @typedef ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Get tracks by artist name
// @Description Get tracks by the artist's name
// @ID get-tracks-by-artist
// @Accept json
// @Produce json
// @Param artistname path string true "Name of the artist"
// @Success 200 {array} models.Musictrack "Successful response"
// @Failure 200 {object} ErrorResponse "Error response"
// @Router /getartist/{artistname} [get]
func (us *TrackService) GetArtistHandlers(c *gin.Context) {
	artistname := c.Param("artistname")

	searchTrack, err := us.trackRepository.GetArtist(artistname)
	if err != nil {
		c.JSON(http.StatusOK, err)
	}
	if searchTrack == nil {
		c.JSON(400, gin.H{
			"message": "unable to get the track by your artist name",
		})
	} else {
		c.JSON(http.StatusOK, searchTrack)
	}

}

// @Summary Save or search track by ISRC
// @Description Save a track in the database by its ISRC or search and store from Spotify if not found
// @ID save-or-search-track-by-isrc
// @Accept json
// @Produce json
// @Param body body object true "ISRC of the track"
// @Success 200 {object} models.Musictrack "Successful response"
// @Failure 400 {object} ErrorResponse "Bad request or ISRC already exists"
// @Router /postisrc [post]
func (us *TrackService) PostISRCHandlers(c *gin.Context) {
	var Isrc struct {
		ISRC string `json:"isrc" binding:"required"`
	}

	if err := c.ShouldBindJSON(&Isrc); err != nil {
		fmt.Println("failed to save track in the database")
	}
	Gettingisrc := Isrc.ISRC

	searchTrack, err := us.trackRepository.GetISRC(Gettingisrc)
	
	var datas = models.Musictrack{
		ISRC: searchTrack.ISRC,
	}

	if len(datas.ISRC) > 0 {
		
		if err != nil {
		
			c.JSON(400, gin.H{
				"message": "error to your isrc",
			})
		}else{
			c.JSON(400, gin.H{
				"message": "isrc  already exist or error ocured",
			})
		}
	
	} else {
		result, err := us.spotifyRepository.SearchTrackByISRC(Gettingisrc)
		if result ==nil || err != nil{
			c.JSON(400, gin.H{
				"message": "unable to get the track by your isrc or find your track",
			})
		 }
		if err == nil {
			Track := &models.Musictrack{
				ISRC:         Gettingisrc,
				Title:        result.Name,
				ArtistName:   result.Artists[0].Name,
				SpotifyImage: result.Album.Images[0].URL,
			}

			err = us.trackRepository.SaveTrack(Track)
			if err != nil {
				fmt.Println("failed to save track in the database")
			}

			c.JSON(http.StatusOK, Track)
		}

	}

}

// @Summary Update track details by ISRC
// @Description Update track information by its ISRC
// @ID update-track-by-isrc
// @Accept json
// @Produce json
// @Param isrc path string true "ISRC of the track"
// @Param body body object true "Track details to update"
// @Success 200 {object} models.Musictrack "Updated track details"
// @Failure 400 {object} ErrorResponse "Invalid user data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /updatetrack/{isrc} [put]
func (us *TrackService) UpdateTrackHandlers(c *gin.Context) {
	isrc := c.Param("isrc")
	searchTrack, err := us.trackRepository.GetISRC(isrc)
	if err != nil {
		fmt.Println("failed to save track in the database")
	}

	
   if len(searchTrack.ISRC) == 0 {
	c.JSON(400,gin.H{
		"message":"give valid syntax",
	})
   }else{
	if err := c.ShouldBindJSON(&searchTrack); err != nil {
		c.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}
	if searchTrack.ISRC==isrc{
		err := us.trackRepository.UpdateTrack(searchTrack)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, &searchTrack)
	   }else{
		c.JSON(400,gin.H{
			"message":"give valid syntax",
		})
	}
	}
	
}
