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

package mddoc

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkoukk/tiktoken-go"
)

type MDDoc struct {
	builder *strings.Builder
}

func New() *MDDoc {
	return &MDDoc{
		builder: new(strings.Builder),
	}
}

func (m *MDDoc) WriteTitle(content string, level int) *MDDoc {
	m.write(m.getTitle(content, level))
	m.Writeln()
	return m
}

func (m *MDDoc) Write(content string) *MDDoc {
	m.write(content)
	m.Writeln()
	return m
}

func (m *MDDoc) WriteMultiCode(content, t string) *MDDoc {
	m.write(m.getMultiCode(content, t))
	return m
}

func (m *MDDoc) Writeln() *MDDoc {
	m.write("\n")
	return m
}

func (m *MDDoc) Export(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(m.builder.String()))
	defer f.Close()
	return err
}

func (m *MDDoc) GetTokens() int {
	tkm, err := tiktoken.EncodingForModel("text-embedding-ada-002")
	if err != nil {
		err = fmt.Errorf("EncodingForModel: %v", err)
		panic(err)
	}
	return len(tkm.Encode(m.builder.String(), nil, nil))
}

func (m *MDDoc) getTitle(content string, level int) string {
	return strings.Repeat("#", level) + " " + content
}

func (m *MDDoc) getMultiCode(content, t string) string {
	return fmt.Sprintf("```%s\n%s\n```\n", t, content)
}

func (m *MDDoc) write(content string) {
	m.builder.WriteString(content)
}
