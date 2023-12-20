package dao

import (
	"fmt"

	"github.com/zmb3/spotify"
)

type SpotifyClient struct {
	client *spotify.Client
}

func SpotifyRepository(client *spotify.Client) *SpotifyClient{
	return  &SpotifyClient{client}
 }

 func (s *SpotifyClient) SearchTrackByISRC(isrc string) (*spotify.FullTrack, error) {
	query := fmt.Sprintf("isrc:%s", isrc)
	searchResult, err := s.client.Search(query, spotify.SearchTypeTrack)
	if err != nil || len(searchResult.Tracks.Tracks) == 0 {
		return nil, fmt.Errorf("track not found")
	}
	
	return &searchResult.Tracks.Tracks[0], nil
}