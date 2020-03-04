package info

import (
	"testing"
)

func TestGet(t *testing.T) {
	user := "dslim"
	repo := "docker-slim"

	result, err := Get(user, repo)
	if err != nil {
		t.Errorf("Get() = Error - %v", err)
	}

	if result.User != user || result.Name != repo {
		t.Errorf("Get() = Expected %q/%q got %q/%q", user, repo, result.User, result.Name)
	}
}
