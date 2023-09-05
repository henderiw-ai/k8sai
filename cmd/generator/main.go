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

package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/google/go-github/github"
	"github.com/henderiw-ai/k8sai/pkg/mddoc"
	"github.com/pkoukk/tiktoken-go"
	"golang.org/x/oauth2"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

const (
	filepath_data = "data"
	filepath_md   = "data/md"
	swagger_file  = "data/swagger.json"
	k8s_swagger   = "https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/openapi-spec/swagger.json"

	k8sipam_owner     = "nokia"
	k8sipam_repo      = "k8s-ipam"
	k8sipam_repo_path = "config/crd/bases"

	kcc_owner     = "GoogleCloudPlatform"
	kcc_repo      = "k8s-config-connector"
	kcc_repo_path = "crds"
)

var client *github.Client

func main() {
	if err := downloadFile(swagger_file, k8s_swagger); err != nil {
		panic(err)
	}

	if err := getGitHubContent(k8sipam_owner, k8sipam_repo, k8sipam_repo_path); err != nil {
		panic(err)
	}
	if err := getGitHubContent(kcc_owner, kcc_repo, kcc_repo_path); err != nil {
		panic(err)
	}

	// generate resources
	if err := generate_schema(filepath_md, swagger_file); err != nil {
		panic(err)
	}

	if err := generateCRDSchema(k8sipam_owner, k8sipam_repo); err != nil {
		panic(err)
	}
	if err := generateCRDSchema(kcc_owner, kcc_repo); err != nil {
		panic(err)
	}
}

func generateCRDSchema(owner, repo string) error {
	f, err := os.ReadDir(filepath.Join(filepath_data, owner, repo))
	if err != nil {
		panic(err)
	}
	for _, f := range f {
		if !f.IsDir() {
			fmt.Println("fileName: ", f.Name())
			if err := generate_schema(filepath_md, filepath.Join(filepath_data, owner, repo, f.Name())); err != nil {
				panic(err)
			}
		}
	}
	return nil
}

func getGitHubContent(owner, repo, path string) error {
	if err := os.Mkdir(filepath.Join(filepath_data, owner), os.ModePerm); err != nil {
		if !strings.Contains(err.Error(), "exists") {
			return err
		}
	}
	if err := os.Mkdir(filepath.Join(filepath_data, owner, repo), os.ModePerm); err != nil {
		if !strings.Contains(err.Error(), "exists") {
			return err
		}
	}

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)

	_, directoryContent, _, err := client.Repositories.GetContents(ctx, owner, repo, path, &github.RepositoryContentGetOptions{})
	if err != nil {
		return err
	}

	for _, c := range directoryContent {
		fmt.Println(*c.Type, *c.Path, *c.Size, *c.SHA)

		local := filepath.Join(filepath_data, owner, repo, filepath.Base(*c.Path))
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
						continue
					}
				}
			}
			if err := downloadContents(owner, repo, local, c); err != nil {
				return err
			}
		case "dir":
			if err := getGitHubContent(owner, repo, filepath.Join(path, *c.Path)); err != nil {
				return err
			}
		}
	}
	return nil
}

func downloadContents(owner, repo, local string, content *github.RepositoryContent) error {
	if content.Content != nil {
		fmt.Println("content:", *content.Content)
	}

	rc, err := client.Repositories.DownloadContents(context.Background(), owner, repo, *content.Path, &github.RepositoryContentGetOptions{})
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

func getDescription(v any) string {
	switch v := v.(type) {
	case map[string]any:
		for k, v := range v {
			if k == "description" {
				switch v := v.(type) {
				case string:
					return v
				}
			}
		}
	}
	return ""
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

func generate_schema(filePath, srcFile string) error {
	if err := os.Mkdir(filePath, os.ModePerm); err != nil {
		if !strings.Contains(err.Error(), "exists") {
			return err
		}
	}
	b, err := os.ReadFile(srcFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("generate schema: ", srcFile)

	switch {
	case strings.HasSuffix(srcFile, ".json"):
		x := map[string]any{}
		if err := json.Unmarshal(b, &x); err != nil {
			panic(err)
		}
		for k, v := range x {
			if k == "definitions" {
				switch v := v.(type) {
				case map[string]any:
					for gvk, v := range v {

						description := getDescription(v)

						gvk := parseGVK(gvk)
						if gvk != nil {
							fmt.Println(gvk.String())

							mddoc := mddoc.New()
							mddoc.WriteTitle(fmt.Sprintf("openAPI schema for %s.%s.%s", gvk.Group, gvk.Version, gvk.Kind), 1)
							mddoc.Writeln()
							mddoc.WriteTitle("description", 2)
							mddoc.Writeln()
							mddoc.Write(description)
							mddoc.Writeln()
							mddoc.WriteTitle("schema", 2)
							mddoc.Writeln()
							b, err := yaml.Marshal(v)
							if err != nil {
								return err
							}
							mddoc.WriteMultiCode(string(b), "yaml")

							tokens := getTokens(string(b))
							//fmt.Println(tokens)

							if len(tokens) > 4000 {
								extensions := parser.AutoHeadingIDs | parser.CommonExtensions
								p := parser.NewWithExtensions(extensions)
								doc := p.Parse(b)

								ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
									switch x := node.(type) {
									default:
										fmt.Printf("type: %v\n", x)
									}
									return ast.GoToNext
								})
								panic(fmt.Sprintf("tokens exceeeded: %d", len(tokens)))
							}

							fmt.Printf("file: %s tokens: %d\n", fmt.Sprintf("%s_%s_%s.md", gvk.Group, gvk.Version, strings.ToLower(gvk.Kind)), mddoc.GetTokens())

							if err := mddoc.Export(filepath.Join(
								filePath,
								fmt.Sprintf("%s_%s_%s.md", gvk.Group, gvk.Version, strings.ToLower(gvk.Kind))),
							); err != nil {
								return err
							}
						}
					}
				}
			}
		}
	case strings.HasSuffix(srcFile, ".yaml"):
		crd := extv1.CustomResourceDefinition{}
		if err := yaml.Unmarshal(b, &crd); err != nil {
			panic(err)
		}
		for _, crdVersion := range crd.Spec.Versions {
			gvk := schema.GroupVersionKind{Group: crd.Spec.Group, Version: crdVersion.Name, Kind: crd.Spec.Names.Kind}
			fmt.Println(gvk.String())

			if _, ok := crdVersion.Schema.OpenAPIV3Schema.Properties["spec"]; ok {
				spec := *crdVersion.Schema.OpenAPIV3Schema
				delete(spec.Properties, "status")

				mddoc := mddoc.New()
				mddoc.WriteTitle(fmt.Sprintf("openAPI spec schema for %s.%s.%s", gvk.Group, gvk.Version, gvk.Kind), 1)
				mddoc.Writeln()
				if crdVersion.Schema.OpenAPIV3Schema.Description != "" {
					mddoc.WriteTitle("description", 2)
					mddoc.Writeln()
					mddoc.Write(crdVersion.Schema.OpenAPIV3Schema.Description)
					mddoc.Writeln()
				}
				mddoc.WriteTitle("schema", 2)
				mddoc.Writeln()

				b, err := yaml.Marshal(crdVersion.Schema.OpenAPIV3Schema)
				if err != nil {
					return err
				}
				tokens := getTokens(string(b))

				mddoc.WriteMultiCode(string(b), "yaml")

				if len(tokens) > 4000 {
					fmt.Printf("tokens exceeded: %d, file: %s\n", len(tokens), fmt.Sprintf("%s_%s_%s_spec.md", gvk.Group, gvk.Version, strings.ToLower(gvk.Kind)))
				}

				// write md file
				if err := mddoc.Export(filepath.Join(
					filePath,
					fmt.Sprintf("%s_%s_%s_spec.md", gvk.Group, gvk.Version, strings.ToLower(gvk.Kind)))); err != nil {
					return err
				}

			}
			if _, ok := crdVersion.Schema.OpenAPIV3Schema.Properties["status"]; ok {
				status := *crdVersion.Schema.OpenAPIV3Schema
				delete(status.Properties, "spec")

				mddoc := mddoc.New()
				mddoc.WriteTitle(fmt.Sprintf("openAPI status schema for %s.%s.%s", gvk.Group, gvk.Version, gvk.Kind), 1)
				mddoc.Writeln()
				if crdVersion.Schema.OpenAPIV3Schema.Description != "" {
					mddoc.WriteTitle("description", 2)
					mddoc.Writeln()
					mddoc.Write(crdVersion.Schema.OpenAPIV3Schema.Description)
					mddoc.Writeln()
				}
				mddoc.WriteTitle("schema", 2)
				mddoc.Writeln()

				b, err := yaml.Marshal(crdVersion.Schema.OpenAPIV3Schema)
				if err != nil {
					return err
				}
				tokens := getTokens(string(b))

				mddoc.WriteMultiCode(string(b), "yaml")

				if len(tokens) > 4000 {
					fmt.Printf("tokens exceeded: %d, file: %s\n", len(tokens), fmt.Sprintf("%s_%s_%s_status.md", gvk.Group, gvk.Version, strings.ToLower(gvk.Kind)))
				}

				// write md file
				if err := mddoc.Export(filepath.Join(
					filePath,
					fmt.Sprintf("%s_%s_%s_status.md", gvk.Group, gvk.Version, strings.ToLower(gvk.Kind)))); err != nil {
					return err
				}

			}
		}

	default:
		return fmt.Errorf("unsupported file extension")
	}

	return nil
}

func getTokens(s string) []int {
	tkm, err := tiktoken.EncodingForModel("text-embedding-ada-002")
	if err != nil {
		err = fmt.Errorf("EncodingForModel: %v", err)
		panic(err)
	}
	return tkm.Encode(s, nil, nil)
}

func parseGVK(gvk string) *schema.GroupVersionKind {
	split := strings.Split(gvk, ".")
	if len(split) > 2 {
		return &schema.GroupVersionKind{
			Group:   strings.Join(split[:len(split)-2], "."),
			Version: split[len(split)-2],
			Kind:    split[len(split)-1],
		}
	}
	return nil
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
