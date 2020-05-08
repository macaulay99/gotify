# gotify [![Codacy Badge](https://api.codacy.com/project/badge/Grade/0951a711ac0a4f5fa9309cfdf41d8e9d)](https://www.codacy.com/app/macaulay99/go_spotify_util?utm_source=github.com&utm_medium=referral&utm_content=macaulay99/go_spotify_util&utm_campaign=badger)

gotify is the wrapper library for [Spotify API](https://developer.spotify.com/web-api/)

## Requirements

- Go 1.9 or later

## Supported Authentication Flow

gotify supported [Authorization Code Flow](https://developer.spotify.com/web-api/authorization-guide/#authorization_code_flow)

## Supported Endpoint

*Endpoints that are not yet supported for optional parameters are also planned to be in order*

### albums

| Endpoint                                 | Struct                    | Method                       | Optional param support |
|------------------------------------------|---------------------------|------------------------------|------------------------|
| GET /v1/albums?ids={ids}                 | Albums                    | GetAlbums                    | ❌                     |
| GET /v1/albums/{id}/tracks               | AlbumsTracks              | GetAlbumsTracks              | ❌                     |


### artists

| Endpoint                                 | Struct                    | Method                       | Optional param support |
|------------------------------------------|---------------------------|------------------------------|------------------------|
| GET /v1/artists?ids={ids}                | Artists                   | GetArtists                   | no option              |
| GET /v1/artists/{id}/albums              | ArtistsAlbums             | GetArtistsAlbums             | ✅                     |
| GET /v1/artists/{id}/top-tracks          | ArtistsTopTracks          | GetArtistsTopTracks          | no option              |
| GET /v1/artists/{id}/related-artists     | ArtistsRelatedArtists     | GetArtistsRelatedArtists     | no option              |

### browse

| Endpoint                                 | Struct                    | Method                       | Optional param support |
|------------------------------------------|---------------------------|------------------------------|------------------------|
| GET /v1/browse/featured-playlists        | BrowseFeaturedPlaylists   | GetBrowseFeaturedPlaylists   | ❌                     |
| GET /v1/browse/new-releases              | BrowseNewReleases         | GetBrowseNewReleases         | ❌                     |
| GET /v1/browse/categories                | BrowseCategories          | GetBrowseCategories          | ❌                     |
| GET /v1/browse/categories/{id}           | BrowseCategory            | GetBrowseCategory            | ❌                     |
| GET /v1/browse/categories/{id}/playlists | BrowseCategorysPlaylists  | GetBrowseCategorysPlaylists  | ❌                     |
| GET /v1/recommendations                  | Recommendations           | GetRecomendations            | ❌                     |

### following

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| GET /v1/me/following?type=artist         | FollowingArtists              | GetFollowingArtists                | ❌                     |
| PUT /v1/me/following                     | -                             | FollowArtistsOrUsers               | ✅                     |
| DELETE /v1/me/following                  | -                             | UnfollowArtistsOrUsers             | ✅                     |
| GET /v1/me/following/contains            | CurrentFollowsArtistsOrUsers  | GetCurrentFollowsArtistsOrUsers    | ✅                     |
| PUT /v1/users/{owner_id}/playlists/{playlist_id}/followers                  | -                            | FollowPlaylists            | ❌                     |
| DELETE /v1/users/{owner_id}/playlists/{playlist_id}/followers               | -                            | UnFollowPlaylists          | ❌                     |
| GET /v1/users/{owner_id}/playlists/{playlist_id}/followers/contains         | FollowPlaylist               | CheckFollowPlaylist        | ✅                     |


### library

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| PUT /v1/me/tracks                        | -                             | SaveTracks                         | ✅                     |
| GET /v1/me/tracks                        | UsersSavedTracks              | GetUsersSavedTracks                | ❌                     |
| DELETE /v1/me/tracks                     | -                             | RemoveUsersSavedTracks             | ✅                     |
| GET /v1/me/tracks/contains               | FollowTracks                  | CheckUsersSavedTracks              | no option              |
| PUT /v1/me/albums?ids={ids}              | -                             | SaveAlbums                         | ✅                     |
| GET /v1/me/albums                        | UsersSavedAlbums              | GetUsersSavedAlbums                | ❌                     |
| DELETE /v1/me/albums?ids={ids}           | -                             | RemoveAlbumsForCurrentUser         | ✅                     |
| GET /v1/me/albums/contains               | FollowAlbums                  | CheckUsersSavedAlbums              | no option              |

### personalization

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| GET /v1/me/player/recently-played        | RecentlyPlayedTracks          | GetRecentlyPlayedTracks            | ❌                     |

### player

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| GET /v1/me/player/devices                | UsersAvailableDevices         | GetUsersAvailableDevices           | ❌                     |
| PUT /v1/me/player                        | -                             | TransferUsersPlayback              | ❌                     |
| PUT /v1/me/player/play                   | -                             | StartResumeUsersPlayback           | ❌                     |
| PUT /v1/me/player/pause                  | -                             | PauseUsersPlayback                 | ❌                     |
| POST /v1/me/player/next                  | -                             | SkipUsersPlaybackToNext            | ❌                     |
| POST /v1/me/player/previous              | -                             | SkipUsersPlaybackToPrevious        | ❌                     |
| PUT /v1/me/player/seek                   | -                             | SeekToPositionInCurrentlyPlayingTrack        | ❌                     |
| PUT /v1/me/player/repeat                 | -                             | SetRepeatModeUsersPlayback         | ❌                     |
| PUT /v1/me/player/volume                 | -                             | SetVolumeUsersPlayback             | ❌                     |
| PUT /v1/me/player/shuffle                | -                             | ToggleShuffleUsersPlayback         | ❌                     |

### search

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| GET /v1/search?type=artist               | SearchArtists                 | SearchArtists                      | ❌                     |
| GET /v1/search?type=album                | SearchAlbums                  | SearchAlbums                       | ❌                     |
| GET /v1/search?type=playlist             | SearchPlaylists               | SearchPlaylists                    | ❌                     |
| GET /v1/search?type=track                | SearchTracks                  | SearchTracks                       | ❌                     |

### tracks

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| GET /v1/tracks?ids={ids}                 | Tracks                        | GetTracks                          | ❌                     |
| GET /v1/audio-analysis/{id}              | AudioAnalysis                 | GetAudioAnalysis                   | no option              |

### profiles

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| GET /v1/me                               | CurrentUsersProfile           | GetCurrentUsersProfile             | no option              |
| GET /v1/users/{user_id}                  | UsersProfile                  | GetUsersProfile                    | no option              |

### playlists

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| GET /v1/users/{user_id}/playlists        | UsersPlaylists                | GetUsersPlaylists                  | ❌                     |
| GET /v1/me/playlists                     | CurrentUsersPlaylists         | GetCurrentUsersPlaylists           | ❌                     |

## Installation

1. Get the `client ID` and `client secret` of your application

2. Run `go get github.com/macaulay99/go_spotify_util`

## Usage

```go
package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/macaulay99/go_spotify_util"
)

var Auth gotify.OAuth
var Token gotify.Gotify

// Handler : Controller for https://localhost:3000/
func Handler(c echo.Context) error {
	Auth = gotify.Set(clientID, clientSecret, callbackURI) // Set and get the basic data for using Spotify API
	url := Auth.AuthURL() // Get the Redirect URL for authenticate
	return c.Redirect(301, url)
}

// CallbackHandler : Controller for https://localhost:3000/callback/
func CallbackHandler(c echo.Context) error {

	t, err := Auth.GetToken(c.Request()) // Get the token for using Spotify API
	if err != nil {
		return err
	}
	Token = t

	return c.String(http.StatusOK, "Authentication success")
}

// RefreshHandler : Controller for https://localhost:3000/refresh/
func RefreshHandler(c echo.Context) error {

	err := Token.Refresh() // Refreshing token for using Spotify API
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "AccessToken Refreshed")
}
```

## Sample

Please see [here](https://github.com/macaulay99/go_spotify_utilSample) for samples
# go_spotify_util
