package git

import (
	"errors"
	"net/http"
	"strings"

	"github.com/google/go-github/v62/github"
)

// IsNotFoundError for git.
func IsNotFoundError(err error) bool {
	var e *github.ErrorResponse
	if errors.As(err, &e) {
		return e.Response.StatusCode == http.StatusNotFound
	}

	return strings.Contains(err.Error(), "no file named")
}
