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

package mem

import "github.com/kstafford3/gorue"

// InMemoryStorage stores and retrieves state from local memory
type InMemoryStorage struct {
	states map[string]gorue.SerializedState
}

// NewInMemoryStorage returns a new StoreRetriever that stores and retrieves state from local memory
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		states: make(map[string]gorue.SerializedState),
	}
}

// Store stores the given state in memory
func (m *InMemoryStorage) Store(identity gorue.StateIdentity, state gorue.SerializedState) error {
	id := string(identity)
	m.states[id] = state
	return nil
}

// Retrieve retrieves the identified state from memory
func (m *InMemoryStorage) Retrieve(identity gorue.StateIdentity) (gorue.SerializedState, error) {
	id := string(identity)
	state := m.states[id]
	return state, nil
}
