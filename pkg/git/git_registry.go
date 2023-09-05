package git

import (
	"fmt"
	"sort"
	"strings"
)

type Initializer func(filePath string) Git

type Registry interface {
	Register(provider string, i Initializer)
	NewGitOfProvider(provider string, filePath string) (Git, error)
}

func NewRegistry() Registry {
	return &registry{
		index: map[string]Initializer{},
	}
}

// registry implementation for fast lookup
type registry struct {
	index map[string]Initializer
}

func (r *registry) Register(provider string, i Initializer) {
	r.index[provider] = i
}

func (r *registry) NewGitOfProvider(provider string, filePath string) (Git, error) {
	i, ok := r.index[provider]
	if !ok {
		return nil, fmt.Errorf("provider %q is not supported. supported providers are %q", provider, strings.Join(r.getRegisteredProviders(), ", "))
	}
	// return a new instance of the requested provider node
	return i(filePath), nil
}

func (r *registry) getRegisteredProviders() []string {
	var result []string
	for k := range r.index {
		result = append(result, k)
	}
	// sort and return
	sort.Strings(result)

	return result
}
