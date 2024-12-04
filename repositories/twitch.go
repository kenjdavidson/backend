package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/streampets/backend/models"
)

type usernameResponse struct {
	Data []struct {
		Login string `json:"login"`
	} `json:"data"`
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type TwitchRepository struct {
	clientID     string
	clientSecret string
	accessToken  string
}

func NewTwitchRepository(id, secret string) (*TwitchRepository, error) {
	repo := &TwitchRepository{clientID: id, clientSecret: secret}
	if err := repo.refreshAccessToken(); err != nil {
		return repo, err
	}
	return repo, nil
}

func (repo *TwitchRepository) GetUsername(userID models.UserID) (string, error) {
	resp, err := repo.getUsernameWithRefresh(userID)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("received incorrect status code from twitch")
	}

	var data usernameResponse
	if err := parseResponse(&data, resp); err != nil {
		return "", err
	}

	return data.Data[0].Login, nil
}

func (repo *TwitchRepository) getUsernameWithRefresh(userID models.UserID) (*http.Response, error) {
	resp, err := getUsername(userID, repo.accessToken, repo.clientID)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 401 {
		return resp, err
	}

	if err := repo.refreshAccessToken(); err != nil {
		return nil, err
	}

	return getUsername(userID, repo.accessToken, repo.clientID)
}

func (repo *TwitchRepository) refreshAccessToken() error {
	resp, err := getAccessToken(repo.clientID, repo.clientSecret)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data accessTokenResponse
	if err := parseResponse(&data, resp); err != nil {
		return err
	}

	repo.accessToken = data.AccessToken
	return nil
}

func parseResponse(data interface{}, resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &data)
}

// Could be in separate TwitchApi file
func getAccessToken(clientID, clientSecret string) (*http.Response, error) {
	return http.PostForm("https://id.twitch.tv/oauth2/token", url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"grant_type":    {"client_credentials"},
	})
}

// Could be in separate TwitchApi file
func getUsername(userID models.UserID, accessToken, clientID string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.twitch.tv/helix/users?id=%s", userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Client-Id", clientID)

	return http.DefaultClient.Do(req)
}
