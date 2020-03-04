package info

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrNotFound = errors.New("not found")

type Image struct {
	User            string    `json:"user,omitempty"`
	Name            string    `json:"name,omitempty"`
	Namespace       string    `json:"namespace,omitempty"`
	RepositoryType  string    `json:"repository_type,omitempty"`
	Status          int       `json:"status,omitempty"`
	Description     string    `json:"description,omitempty"`
	FullDescription string    `json:"full_description,omitempty"`
	IsPrivate       bool      `json:"is_private,omitempty"`
	IsAutomated     bool      `json:"is_automated,omitempty"`
	CanEdit         bool      `json:"can_edit,omitempty"`
	IsMigrated      bool      `json:"is_migrated,omitempty"`
	HasStarred      bool      `json:"has_starred,omitempty"`
	StarCount       uint64    `json:"star_count,omitempty"`
	PullCount       uint64    `json:"pull_count,omitempty"`
	LastUpdated     time.Time `json:"last_updated,omitempty"`
}

const (
	clientTimeout     = 13
	imageInfoEndpoint = "https://hub.docker.com/v2/repositories"
)

func Get(user, repo string) (*Image, error) {
	client := http.Client{
		Timeout: clientTimeout * time.Second,
	}

	endpoint := fmt.Sprintf("%s/%s/%s", imageInfoEndpoint, user, repo)
	var result Image

	resp, err := client.Get(endpoint)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == http.StatusNotFound {
				return nil, ErrNotFound
			}

			return nil, fmt.Errorf("bad status - %d", resp.StatusCode)
		}

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&result)
		if err != nil {
			return nil, err
		}

		return &result, nil
	}

	return nil, err
}
