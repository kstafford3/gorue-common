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

package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kstafford3/gorue"
)

var dirMode os.FileMode = 0755
var fileMode os.FileMode = 0644

// FileStorage stores and retrieves state in a provided directory
type FileStorage struct {
	path string
}

// NewFileStorage returns a new StoreRetriever that stores and retrieves state in the directory specified by path
func NewFileStorage(path string) *FileStorage {
	os.MkdirAll(path, dirMode)
	return &FileStorage{
		path: path,
	}
}

// Store stores the given state in a file
func (s *FileStorage) Store(identity gorue.StateIdentity, state gorue.SerializedState) error {
	id := string(identity)
	filename := filepath.Join(s.path, id)
	err := ioutil.WriteFile(filename, state, fileMode)
	return err
}

// Retrieve retrieves the identified state from memory
func (s *FileStorage) Retrieve(identity gorue.StateIdentity) (gorue.SerializedState, error) {
	id := string(identity)
	filename := filepath.Join(s.path, id)
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, nil
	}
	state, err := ioutil.ReadFile(filename)
	return gorue.SerializedState(state), nil
}
