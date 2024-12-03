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
	ClientID     string
	ClientSecret string
}

func NewTwitchRepository(id, secret string) *TwitchRepository {
	return &TwitchRepository{ClientID: id, ClientSecret: secret}
}

func (repo *TwitchRepository) GetUsername(userID models.UserID) (string, error) {
	url := fmt.Sprintf("https://api.twitch.tv/helix/users?id=%s", userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	accessToken, err := repo.getAccessToken()
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Client-Id", repo.ClientID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(string(body))
	}

	var data usernameResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	return data.Data[0].Login, nil
}

func (repo *TwitchRepository) getAccessToken() (string, error) {
	resp, err := http.PostForm("https://id.twitch.tv/oauth2/token", url.Values{
		"client_id":     {repo.ClientID},
		"client_secret": {repo.ClientSecret},
		"grant_type":    {"client_credentials"},
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data accessTokenResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	return data.AccessToken, nil
}
