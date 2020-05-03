// Copyright 2020 Kyle Stafford

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 		http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package simple

import (
	"bufio"
	"io"
	"strings"
)

// Prompter prompts the user for a response using the provided io.Reader and io.Writer.
type Prompter struct {
	reader *bufio.Reader
	writer *bufio.Writer
	prompt string
}

// NewPrompter returns a new SimplePrompter that reads from the provided reader and writes to the provided writer
func NewPrompter(reader io.Reader, writer io.Writer) *Prompter {
	bufReader := bufio.NewReader(reader)
	bufWriter := bufio.NewWriter(writer)
	return &Prompter{
		reader: bufReader,
		writer: bufWriter,
	}
}

func (p Prompter) write(out string) error {
	_, err := p.writer.WriteString(out)
	if err != nil {
		return err
	}
	err = p.writer.Flush()
	return err
}

func (p Prompter) read() (string, error) {
	raw, err := p.reader.ReadString('\n')
	if err != nil {
		return raw, err
	}
	return strings.TrimSpace(raw), nil
}

// Prompt prints a prompt to the command line and returns the response
func (p Prompter) Prompt(out string) (string, error) {
	var err = p.write(out + "\n" + p.prompt)
	if err != nil {
		return "", err
	}
	response, err := p.read()
	return response, err
}
