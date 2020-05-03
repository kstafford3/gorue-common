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
	"bytes"
	"errors"
	"io"
	"testing"
)

var errExpected = errors.New("Expected Error")

type FailingReader struct {
	io.Reader
}

func (f FailingReader) Read(p []byte) (n int, err error) {
	return 0, errExpected
}

type FailingStringWriter struct {
	io.Writer
}

func (f FailingStringWriter) Write(p []byte) (n int, err error) {
	return 0, errExpected
}

func TestRead(t *testing.T) {
	testString := "Hello, World!\n"
	in := bytes.NewBufferString(testString)
	reader := bufio.NewReader(in)
	out := bytes.NewBufferString("")
	writer := bufio.NewWriter(out)
	prompter := NewPrompter(reader, writer)
	expected := "Hello, World!"
	actual, err := prompter.read()

	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("Expected '%s', but read '%s'", expected, actual)
	}
}

func TestReadMultiline(t *testing.T) {
	testString := "Hello, World!\nA second line.\n"
	in := bytes.NewBufferString(testString)
	reader := bufio.NewReader(in)
	out := bytes.NewBufferString("")
	writer := bufio.NewWriter(out)
	prompter := NewPrompter(reader, writer)
	expectedLine1 := "Hello, World!"
	expectedLine2 := "A second line."

	actualLine1, errLine1 := prompter.read()
	if errLine1 != nil {
		t.Error(errLine1)
	}
	if actualLine1 != expectedLine1 {
		t.Errorf("Expected '%s', but read '%s'", expectedLine1, actualLine1)
	}

	actualLine2, errLine2 := prompter.read()
	if errLine2 != nil {
		t.Error(errLine2)
	}
	if actualLine2 != expectedLine2 {
		t.Errorf("Expected '%s', but read '%s'", expectedLine1, actualLine2)
	}
}

func TestReadError(t *testing.T) {
	reader := FailingReader{}
	out := bytes.NewBufferString("")
	writer := bufio.NewWriter(out)
	prompter := NewPrompter(reader, writer)
	_, err := prompter.read()

	if err == nil {
		t.Errorf("Expected Read error")
	}
}

func TestWrite(t *testing.T) {
	in := bytes.NewBufferString("")
	reader := bufio.NewReader(in)
	out := bytes.NewBufferString("")
	writer := bufio.NewWriter(out)
	prompter := NewPrompter(reader, writer)

	testString := "Hello, World!"
	expected := bytes.NewBufferString("Hello, World!")
	err := prompter.write(testString)

	if err != nil {
		t.Error(err)
	}

	if out.String() != expected.String() {
		t.Errorf("Expected '%s', but wrote '%s'", expected, out)
	}
}

func TestWriteStringError(t *testing.T) {
	in := bytes.NewBufferString("")
	reader := bufio.NewReader(in)
	writer := FailingStringWriter{}
	bufWriter := bufio.NewWriterSize(writer, 1) // use a tiny buffer to force WriteString to write, triggering an error
	prompter := Prompter{
		reader: reader,
		writer: bufWriter,
	}

	testString := "Hello, World!"
	err := prompter.write(testString)

	if err == nil {
		t.Errorf("Expected WriteString error")
	}
}

func TestWriteFlushError(t *testing.T) {
	in := bytes.NewBufferString("")
	reader := bufio.NewReader(in)
	writer := FailingStringWriter{}
	prompter := NewPrompter(reader, writer)

	testString := "Hello, World!"
	err := prompter.write(testString)

	if err == nil {
		t.Errorf("Expected WriteFlush error")
	}
}

func TestPrompt(t *testing.T) {
	in := bytes.NewBufferString("Hello, World!\n")
	expectedResponse := "Hello, World!"
	reader := bufio.NewReader(in)
	out := bytes.NewBufferString("")
	writer := bufio.NewWriter(out)
	prompter := NewPrompter(reader, writer)
	prompter.prompt = "> "

	testPrompt := "This is the prompt."
	expectedOut := bytes.NewBufferString("This is the prompt.\n> ")
	actualResponse, err := prompter.Prompt(testPrompt)

	if err != nil {
		t.Error(err)
	}

	if expectedOut.String() != out.String() {
		t.Errorf("Expected '%s', but got '%s' as prompt", expectedOut, out)
	}

	if expectedResponse != actualResponse {
		t.Errorf("Expected '%s', but got '%s' for response", expectedResponse, actualResponse)
	}
}

func TestPromptWriteError(t *testing.T) {
	in := bytes.NewBufferString("Hello, World!\n")
	reader := bufio.NewReader(in)
	writer := FailingStringWriter{}
	prompter := NewPrompter(reader, writer)
	prompter.prompt = "> "

	testPrompt := "This is the prompt."
	_, err := prompter.Prompt(testPrompt)

	if err == nil {
		t.Errorf("Expected a write error")
	}
}

func TestPromptReadError(t *testing.T) {
	reader := FailingReader{}
	out := bytes.NewBufferString("")
	writer := bufio.NewWriter(out)
	prompter := NewPrompter(reader, writer)
	prompter.prompt = "> "

	testPrompt := "This is the prompt."
	expectedOut := bytes.NewBufferString("This is the prompt.\n> ")
	_, err := prompter.Prompt(testPrompt)

	if err == nil {
		t.Errorf("Expected a read error")
	}

	if expectedOut.String() != out.String() {
		t.Errorf("Expected '%s', but got '%s' as prompt", expectedOut, out)
	}
}
