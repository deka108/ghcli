package ghutil

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	github "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type MalformedFileError string

func (e MalformedFileError) Error() string { return string(e) }

// NewAuthorizedClient creates authorized GitHub client
func NewAuthorizedClient(accessToken string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

	return client
}

// NewAuthorizedClientFromEnv creates authorized GitHub client from GITHUB_TOKEN environment variable
func NewAuthorizedClientFromEnv() (*github.Client, error) {
	accessToken := os.Getenv("GITHUB_TOKEN")
	if accessToken == "" {
		return nil, errors.New("the environment variable GITHUB_TOKEN must be set")
	}
	return NewAuthorizedClient(accessToken), nil
}

// NewAuthorizedClientFromSecretFile creates authorized GitHub client from a secret file
func NewAuthorizedClientFromSecretFile(secretFile string) (*github.Client, error) {
	bytes, err := ioutil.ReadFile(secretFile)

	if err != nil {
		return nil, err
	}
	var data interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	v, ok := data.(map[string]interface{})
	if !ok {
		return nil, MalformedFileError("malformed key value json")
	}
	if v["token"] == nil {
		return nil, MalformedFileError("\"token\" field is not found in the secret file")
	}
	accessToken := v["token"].(string)

	return NewAuthorizedClient(accessToken), nil
}
