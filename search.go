package gotify

import (
	"encoding/json"
	"strings"

	"github.com/macaulay99/gotify/extensions"
	"github.com/macaulay99/gotify/models"
)

func encodeQuery(q string) string {
	qstring := strings.Replace(q, " ", "%20", -1)
	return qstring
}

// SearchArtists : the method for GET https://api.spotify.com/v1/search
func (t *Tokens) SearchArtists(keywords string) (*models.SearchArtists, error) {
	/**
	https://developer.spotify.com/web-api/search-item/
	*/

	query := encodeQuery(keywords)

	endpoint := "https://api.spotify.com/v1/search?q=" + query + "&type=artist"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	artists := new(models.SearchArtists)

	err = json.Unmarshal(res, artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

// SearchAlbums : the method for GET https://api.spotify.com/v1/search
func (t *Tokens) SearchAlbums(keywords string) (*models.SearchAlbums, error) {
	/**
	https://developer.spotify.com/web-api/search-item/
	*/

	query := encodeQuery(keywords)

	endpoint := "https://api.spotify.com/v1/search?q=" + query + "&type=album"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	albums := new(models.SearchAlbums)

	err = json.Unmarshal(res, albums)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

// SearchPlaylists : the method for GET https://api.spotify.com/v1/search
func (t *Tokens) SearchPlaylists(keywords string) (*models.SearchPlaylists, error) {
	/**
	https://developer.spotify.com/web-api/search-item/
	*/

	query := encodeQuery(keywords)

	endpoint := "https://api.spotify.com/v1/search?q=" + query + "&type=playlist"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	playlists := new(models.SearchPlaylists)

	err = json.Unmarshal(res, playlists)
	if err != nil {
		return nil, err
	}
	return playlists, nil
}

// SearchTracks : the method for GET https://api.spotify.com/v1/search
func (t *Tokens) SearchTracks(keywords string) (*models.SearchTracks, error) {
	/**
	https://developer.spotify.com/web-api/search-item/
	*/

	query := encodeQuery(keywords)

	endpoint := "https://api.spotify.com/v1/search?q=" + query + "&type=track"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	tracks := new(models.SearchTracks)

	err = json.Unmarshal(res, tracks)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}
