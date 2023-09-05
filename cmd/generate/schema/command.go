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

package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/henderiw-ai/k8sai/apis"
	docs "github.com/henderiw-ai/k8sai/internal/docs/generated/generatedocs"
	"github.com/henderiw-ai/k8sai/pkg/git"
	"github.com/henderiw-ai/k8sai/pkg/mddoc"
	"github.com/henderiw-ai/k8sai/pkg/util"
	"github.com/pkoukk/tiktoken-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

func NewCommand(ctx context.Context) *cobra.Command {
	return NewRunner(ctx).Command
}

// NewRunner returns a command runner.
func NewRunner(ctx context.Context) *Runner {
	r := &Runner{
		Ctx: ctx,
	}
	c := &cobra.Command{
		Use:     "schema",
		Args:    cobra.MaximumNArgs(0),
		Short:   docs.SchemaShort,
		Long:    docs.SchemaShort + "\n" + docs.SchemaLong,
		Example: docs.SchemaExamples,
		RunE:    r.runE,
	}

	r.Command = c
	return r
}

type Runner struct {
	Command *cobra.Command
	Ctx     context.Context
}

func (r *Runner) runE(c *cobra.Command, args []string) error {

	for _, repo := range viper.GetStringSlice(apis.ViperReposKey) {
		apiRepo := apis.Repo{}
		if err := json.Unmarshal([]byte(repo), &apiRepo); err != nil {
			return err
		}
		parsedRepo, err := git.ParseRepoURL(apiRepo.RepoURL)
		if err != nil {
			return err
		}

		f, err := os.ReadDir(filepath.Join(apis.BaseFilePath, parsedRepo.Owner, parsedRepo.Repo))
		if err != nil {
			panic(err)
		}

		for _, f := range f {
			if !f.IsDir() {
				if err := generate_schema(filepath.Join(apis.BaseFilePath, parsedRepo.Owner, parsedRepo.Repo, f.Name())); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func generate_schema(srcFile string) error {
	util.EnsureDir(apis.BaseMDFilePath)

	b, err := os.ReadFile(srcFile)
	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(srcFile, ".json"):
		x := map[string]any{}
		if err := json.Unmarshal(b, &x); err != nil {
			return err
		}
		for k, v := range x {
			if k == "definitions" {
				switch v := v.(type) {
				case map[string]any:
					for gvk, v := range v {
						description := getDescription(v)

						b, err := yaml.Marshal(v)
						if err != nil {
							return err
						}

						gvk := parseGVK(gvk)
						if gvk != nil {
							if err := createSchema("status", gvk, description, string(b)); err != nil {
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
			return err
		}
		for _, crdVersion := range crd.Spec.Versions {
			gvk := &schema.GroupVersionKind{Group: crd.Spec.Group, Version: crdVersion.Name, Kind: crd.Spec.Names.Kind}

			fmt.Printf("file: %s\n", fmt.Sprintf("%s_%s_%s.md", gvk.Group, gvk.Version, strings.ToLower(gvk.Kind)))

			splitter := New()
			splitter.splitData([]string{}, crdVersion.Schema.OpenAPIV3Schema.Properties)

			for idx, item := range splitter.getItems() {
				fmt.Printf("idx: %d, tokens: %d\n", idx, item.tokens)

				schema := *crdVersion.Schema.OpenAPIV3Schema
				schema.Properties = map[string]extv1.JSONSchemaProps{}
				schema.Properties = item.Properties
				if err := createSchema(
					fmt.Sprintf("part%d", idx),
					gvk,
					crdVersion.Schema.OpenAPIV3Schema.Description,
					schema); err != nil {
					return err
				}
			}
			/*
				if len(splitter.getItems()) > 1 {
					panic(1)
				}
			*/
		}
	}
	return nil
}

func NewItem() Item {
	return Item{tokens: 0, Properties: map[string]extv1.JSONSchemaProps{}}
}

func New() *splitter {
	s := &splitter{
		items: []Item{},
	}
	s.addNewItem()
	return s
}

type splitter struct {
	items []Item
}

func (r *splitter) addNewItem() {
	r.items = append(r.items, NewItem())
}

const (
	maxTokens = 3500
)

func (r *splitter) getItems() []Item {
	return r.items
}

func (r *splitter) IsTokenLimitExceeded(tokens int) bool {
	return r.items[len(r.items)-1].tokens+tokens > maxTokens
}

func (r *splitter) addToken(tokens int) {
	r.items[len(r.items)-1].tokens += tokens
}

func (r *splitter) addProp(propItems []string, propData extv1.JSONSchemaProps) {
	properties := r.items[len(r.items)-1].Properties
	for idx, prop := range propItems {
		fmt.Printf("addProp: idx: %d propItems: %v, keys: %v, items: %d\n", idx, propItems, getSortedKeys(properties), len(r.items))
		if idx == len(propItems)-1 {
			properties[prop] = propData
		}
		properties = properties[prop].Properties
	}
}

func (r *splitter) addSplitProp(propItems []string, propData extv1.JSONSchemaProps, delete bool) {

	fmt.Printf("addSplitProp: propItems: %v, items: %d\n", propItems, len(r.items))

	/*
		if len(r.items) >= 3 {
			fmt.Printf("items latest-1\n")
			fmt.Printf("items latest-1, props: %v\n", getSortedKeys(r.items[len(r.items)-2].Properties))
			for prop, v := range r.items[len(r.items)-2].Properties {
				fmt.Printf("items latest-1, prop: %s, props: %v\n", prop, getSortedKeys(v.Properties))
				for prop, v := range v.Properties {
					fmt.Printf("items latest-1, prop: %s, props: %v\n", prop, getSortedKeys(v.Properties))
					for prop, v := range v.Properties {
						fmt.Printf("items latest-1, prop: %s, props: %v\n", prop, getSortedKeys(v.Properties))
					}
				}
			}

			fmt.Printf("items latest")
			fmt.Printf("items latest, props: %v\n", getSortedKeys(r.items[len(r.items)-1].Properties))
			for prop, v := range r.items[len(r.items)-1].Properties {
				fmt.Printf("items latest, prop: %s, props: %v\n", prop, getSortedKeys(v.Properties))
				for prop, v := range v.Properties {
					fmt.Printf("items latest, prop: %s, props: %v\n", prop, getSortedKeys(v.Properties))
					for prop, v := range v.Properties {
						fmt.Printf("items latest, prop: %s, props: %v\n", prop, getSortedKeys(v.Properties))
					}
				}
			}
			panic(1)
		}
	*/

	r.items[len(r.items)-1].Properties = map[string]extv1.JSONSchemaProps{}
	properties := r.items[len(r.items)-1].Properties
	for idx, prop := range propItems {
		properties[prop] = extv1.JSONSchemaProps{
			Properties: map[string]extv1.JSONSchemaProps{},
		}
		properties = properties[prop].Properties

		if len(propItems)-1 == idx {
			// copy the data
			newProps := propData
			// remove the properties
			if delete {
				newProps.Properties = map[string]extv1.JSONSchemaProps{}
			}
			properties[prop] = newProps
		}
	}
}

func (r *splitter) splitData(propItems []string, props map[string]extv1.JSONSchemaProps) error {
	fmt.Printf("  splitData: propItems: %v, items %d\n", propItems, len(r.items))
	sortedKeys := getSortedKeys(props)
	for _, prop := range sortedKeys {
		b, err := yaml.Marshal(props[prop])
		if err != nil {
			return err
		}
		tokens := getTokens(string(b))

		if !r.IsTokenLimitExceeded(len(tokens)) {
			fmt.Printf("  splitData: token ok %d, prop: %s\n", len(tokens), prop)
			r.addToken(len(tokens))
			r.addProp(append(propItems, prop), props[prop])
		} else {
			// add a new item in the list
			r.addNewItem()
			if !r.IsTokenLimitExceeded(len(tokens)) {
				fmt.Printf("  splitData: token exceeded no split %d, prop: %s\n", len(tokens), prop)
				r.addToken(len(tokens))
				// we need to copy the relevant data from the previous item
				r.addSplitProp(append(propItems, prop), props[prop], false)
				//r.addProp(append(propItems, prop), props[prop])
			} else {
				fmt.Printf("  splitData: token exceeded split %d, prop: %s\n", len(tokens), prop)
				// we need to copy the relevant data from the previous item
				r.addSplitProp(append(propItems, prop), props[prop], true)
				// to be further added
				r.splitData(append(propItems, prop), props[prop].Properties)
			}
		}
	}
	return nil
}

func getSortedKeys(unsortedMap map[string]extv1.JSONSchemaProps) []string {
	// Get the keys from the map
	keys := make([]string, 0, len(unsortedMap))
	for key := range unsortedMap {
		keys = append(keys, key)
	}

	// Sort the keys
	sort.Strings(keys)

	return keys
}

/*
func getTokenPropMap(propMap map[string]PropMap, props map[string]extv1.JSONSchemaProps) error {
	for prop, propData := range props {
		b, err := yaml.Marshal(propData)
		if err != nil {
			return err
		}
		tokens := getTokens(string(b))
		propMap[prop] = PropMap{
			data:    propData,
			tokens:  len(tokens),
			propMap: map[string]PropMap{},
		}
		if len(tokens) > 3500 {
			getTokenPropMap(propMap[prop].propMap, propData.Properties)
		}
	}
	return nil
}


*/

func addData(groupData map[string]extv1.JSONSchemaProps, chainedProps []string, properties extv1.JSONSchemaProps) map[string]extv1.JSONSchemaProps {
	fmt.Printf("addData: groupData %v, chainedprops: %v\n properties: %v\n", len(groupData), chainedProps, properties)
	d := groupData
	if len(d) == 0 {
		d = map[string]extv1.JSONSchemaProps{}
	}
	for idx, prop := range chainedProps {
		d = d[prop].Properties
		if len(d) == 0 {
			d = map[string]extv1.JSONSchemaProps{}
		}
		if idx == len(chainedProps)-1 {
			d[prop] = properties
			return d
		}
	}
	/*
		fmt.Printf("addData: chainedprops: %v\n properties: %v\n", chainedProps, properties)
		if len(chainedProps) == 1 {
			if len(groupData) == 0 {
				groupData = map[string]extv1.JSONSchemaProps{}
			}
			groupData[chainedProps[0]] = properties
		} else {
			addData(groupData[chainedProps[0]].Properties, chainedProps[1:], properties)
		}
	*/
	return d
}

/*
	func buildOutputMap(groups []Item, propMap map[string]PropMap, chainedProps []string) ([]Item, map[string]extv1.JSONSchemaProps) {
		fmt.Printf("buildOutputMap: groups %d, chainedprops: %v\n", len(groups), chainedProps)
		sortedKeys := getSortedKeys(propMap)
		for _, prop := range sortedKeys {
			// idx is the last entry in the group -> we always append to the last entry of the group
			idx := len(groups) - 1
			if groups[idx].tokens+propMap[prop].tokens < 3500 {
				groups[idx].tokens += propMap[prop].tokens
				groups[idx].Properties = addData(
					groups[idx].Properties,
					append(chainedProps, prop),
					propMap[prop].data,
				)
			} else {
				// if we exceed the tokens we add a new item to the group
				groups = append(groups, Item{tokens: 0, Properties: map[string]extv1.JSONSchemaProps{}}) // append the group

				idx := len(groups) - 1
				if groups[idx].tokens+propMap[prop].tokens < 3500 {
					groups[idx].tokens += propMap[prop].tokens
					groups[idx].Properties = addData(
						groups[idx].Properties,
						append(chainedProps, prop),
						propMap[prop].data,
					)
				} else {
					// if a single item exceeds the token limit we need to split
					// add the initial element to the props w/o the properties
					groups[idx].Properties[prop] = propMap[prop].data
					for k := range groups[idx].Properties[prop].Properties {
						delete(groups[idx].Properties[prop].Properties, k)
					}
					groups, groups[idx].Properties = buildOutputMap(
						groups,
						propMap[prop].propMap,
						append(chainedProps, prop), // append the prop
					)
				}
			}
		}
		idx := len(groups) - 1
		return groups, groups[idx].Properties
	}
*/
type Item struct {
	tokens     int
	Properties map[string]extv1.JSONSchemaProps
}

/*
func printPropMap(level int, propMap map[string]PropMap) {
	sortedKeys := getSortedKeys(propMap)
	for _, prop := range sortedKeys {
		indent := strings.Repeat("  ", level)
		b, _ := yaml.Marshal(propMap[prop].data)
		fmt.Printf("%sprop: %s, tokens: %d, data: %d\n", indent, prop, propMap[prop].tokens, len(string(b)))
		if len(propMap[prop].propMap) > 0 {
			level++
			printPropMap(level, propMap[prop].propMap)
		}
	}
}
*/

type PropMap struct {
	tokens  int
	data    extv1.JSONSchemaProps
	propMap map[string]PropMap
}

func createSchema(info string, gvk *schema.GroupVersionKind, description string, data any) error {
	//fmt.Println("gvkString", gvk.String())

	mddoc := mddoc.New()
	mddoc.WriteTitle(fmt.Sprintf("openAPI schema %s for %s.%s.%s", info, gvk.Group, gvk.Version, gvk.Kind), 1)
	mddoc.Writeln()
	if description != "" {
		mddoc.WriteTitle("description", 2)
		mddoc.Writeln()
		mddoc.Write(description)
		mddoc.Writeln()
	}
	mddoc.WriteTitle("schema", 2)
	mddoc.Writeln()
	b, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	mddoc.WriteMultiCode(string(b), "yaml")

	fmt.Printf("file: %s\n", fmt.Sprintf("%s_%s_%s_%s.md", gvk.Group, gvk.Version, strings.ToLower(gvk.Kind), info))

	if err := mddoc.Export(filepath.Join(
		apis.BaseMDFilePath,
		fmt.Sprintf("%s_%s_%s_%s.md", gvk.Group, gvk.Version, strings.ToLower(gvk.Kind), info)),
	); err != nil {
		return err
	}
	return nil
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

func getTokens(s string) []int {
	tkm, err := tiktoken.EncodingForModel("text-embedding-ada-002")
	if err != nil {
		err = fmt.Errorf("EncodingForModel: %v", err)
		panic(err)
	}
	return tkm.Encode(s, nil, nil)
}
