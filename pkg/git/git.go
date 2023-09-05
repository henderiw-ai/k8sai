package git

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type Git interface {
	GetContent(ctx context.Context, owner, repo, path string) error
}

type Repo struct {
	Host  string
	Owner string
	Repo  string
}

func ParseRepoURL(u string) (*Repo, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	splittedRepo := strings.FieldsFunc(parsedURL.Path, func(c rune) bool { return c == '/' }) // trim empty fields from returned slice
	if len(splittedRepo) < 2 {
		return nil, fmt.Errorf("expecting <owner>/<repo> in url path, received: '%s'", parsedURL.Path)
	}

	return &Repo{
		Host:  parsedURL.Host,
		Owner: splittedRepo[0],
		Repo:  strings.TrimSuffix(splittedRepo[1], ".git"),
	}, nil

}
