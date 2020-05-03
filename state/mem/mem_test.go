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

import (
	"testing"

	"github.com/kstafford3/gorue"
)

func TestStoreAndRetrieve(t *testing.T) {
	mem := NewInMemoryStorage()
	identity := gorue.StateIdentity("state-identity")
	expectedState := gorue.SerializedState("serialized-state")

	err := mem.Store(identity, expectedState)
	if err != nil {
		t.Error(err)
	}
	actualState, err := mem.Retrieve(identity)
	if err != nil {
		t.Error(err)
	}

	if string(expectedState) != string(actualState) {
		t.Errorf("Expected '%s' but retrieved state was '%s'", expectedState, actualState)
	}
}

func TestStoreAndRetrieveMultiple(t *testing.T) {
	mem := NewInMemoryStorage()
	identityOne := gorue.StateIdentity("state-identity-one")
	expectedStateOne := gorue.SerializedState("serialized-state-one")
	identityTwo := gorue.StateIdentity("state-identity-two")
	expectedStateTwo := gorue.SerializedState("serialized-state-two")

	var err error
	err = mem.Store(identityOne, expectedStateOne)
	if err != nil {
		t.Error(err)
	}
	err = mem.Store(identityTwo, expectedStateTwo)
	if err != nil {
		t.Error(err)
	}

	actualStateOne, err := mem.Retrieve(identityOne)
	if err != nil {
		t.Error(err)
	}
	actualStateTwo, err := mem.Retrieve(identityTwo)
	if err != nil {
		t.Error(err)
	}

	if string(actualStateOne) != string(expectedStateOne) {
		t.Errorf("Expected '%s' but retrieved state was '%s'", expectedStateOne, actualStateOne)
	}
	if string(actualStateTwo) != string(expectedStateTwo) {
		t.Errorf("Expected '%s' but retrieved state was '%s'", expectedStateTwo, actualStateTwo)
	}
}

func TestRetrieveMissing(t *testing.T) {
	mem := NewInMemoryStorage()
	identity := gorue.StateIdentity("state-identity")
	actualState, err := mem.Retrieve(identity)
	if err != nil {
		t.Error(err)
	}
	if nil != actualState {
		t.Errorf("Expected 'nil' but retrieved state was '%s'", actualState)
	}
}
