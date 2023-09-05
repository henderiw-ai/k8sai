/*
Copyright 2023 Nokia.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package github

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/henderiw-ai/k8sai/apis"
	"github.com/henderiw-ai/k8sai/pkg/git"
	"github.com/henderiw-ai/k8sai/pkg/util"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

const (
	GitHubProvider = "github.com"
)

// Register registers the provider.
func Register(r git.Registry) {
	r.Register(GitHubProvider, func(filePath string) git.Git {
		return &gh{
			filePath: filePath,
			client:   github.NewClient(nil),
		}
	})
}

func New(filePath string) git.Git {
	return &gh{
		filePath: filePath,
	}
}

type gh struct {
	filePath string
	client   *github.Client
}

func (r *gh) GetContent(ctx context.Context, owner, repo, path string) error {
	// check if the owner directory exists
	util.EnsureDir(r.filePath, owner, repo)

	if viper.GetString(apis.ViperGithubTokenKey) != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: viper.GetString(apis.ViperGithubTokenKey)},
		)
		tc := oauth2.NewClient(ctx, ts)

		r.client = github.NewClient(tc)
	}

	c, directoryContent, _, err := r.client.Repositories.GetContents(ctx, owner, repo, path, &github.RepositoryContentGetOptions{})
	if err != nil {
		return err
	}

	for _, c := range directoryContent {
		if err := r.getContent(ctx, owner, repo, path, c); err != nil {
			return err
		}
	}
	if c != nil {
		if err := r.getContent(ctx, owner, repo, path, c); err != nil {
			return err
		}
	}
	return nil
}

func (r *gh) getContent(ctx context.Context, owner, repo, path string, c *github.RepositoryContent) error {
	fmt.Println(*c.Type, *c.Path, *c.Size, *c.SHA)

	local := filepath.Join(apis.BaseFilePath, owner, repo, filepath.Base(*c.Path))
	fmt.Println("local:", local)

	switch *c.Type {
	case "file":
		_, err := os.Stat(local)
		if err == nil {
			b, err1 := ioutil.ReadFile(local)
			if err1 == nil {
				sha := calculateGitSHA1(b)
				if *c.SHA == hex.EncodeToString(sha) {
					fmt.Println("no need to update this file, the SHA is the same")
					return nil
				}
			}
		}
		if err := r.downloadContents(ctx, owner, repo, local, c); err != nil {
			return err
		}
	case "dir":
		if err := r.GetContent(ctx, owner, repo, filepath.Join(path, *c.Path)); err != nil {
			return err
		}
	}
	return nil
}

func (r *gh) downloadContents(ctx context.Context, owner, repo, local string, content *github.RepositoryContent) error {
	if content.Content != nil {
		fmt.Println("content:", *content.Content)
	}

	rc, err := r.client.Repositories.DownloadContents(ctx, owner, repo, *content.Path, &github.RepositoryContentGetOptions{})
	if err != nil {
		return err
	}
	defer rc.Close()

	b, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(local), 0666)
	if err != nil {
		return err
	}

	fmt.Println("Writing the file:", local)
	f, err := os.Create(local)
	if err != nil {
		return err
	}
	defer f.Close()
	n, err := f.Write(b)
	if err != nil {
		fmt.Println(err)
	}
	if n != *content.Size {
		fmt.Printf("number of bytes differ, %d vs %d\n", n, *content.Size)
	}
	return nil
}

// calculateGitSHA1 computes the github sha1 from a slice of bytes.
// The bytes are prepended with: "blob " + filesize + "\0" before runing through sha1.
func calculateGitSHA1(contents []byte) []byte {
	contentLen := len(contents)
	blobSlice := []byte("blob " + strconv.Itoa(contentLen))
	blobSlice = append(blobSlice, '\x00')
	blobSlice = append(blobSlice, contents...)
	h := sha1.New()
	h.Write(blobSlice)
	bs := h.Sum(nil)
	return bs
}
