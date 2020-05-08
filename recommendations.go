package gotify

import (
	"encoding/json"

	"github.com/macaulay99/gotify/extensions"
	"github.com/macaulay99/gotify/models"
)

// GetRecommendations : the method for GET https://api.spotify.com/v1/recommendations
func (t *Tokens) GetRecommendations() (*models.Recommendations, error) {
	/**
	https://developer.spotify.com/web-api/get-recommendations/
	*/

	endpoint := "https://api.spotify.com/v1/recommendations"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	recommendations := new(models.Recommendations)

	err = json.Unmarshal(res, recommendations)
	if err != nil {
		return nil, err
	}
	return recommendations, nil
}
