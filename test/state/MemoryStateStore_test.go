package test_state

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-components-go/state"
)

func newStateStoreFixture() *StateStoreFixture {
	stateStore := state.NewEmptyMemoryStateStore()
	fixture := NewStateStoreFixture(stateStore)
	return fixture
}

func TestConsoleLogLevel(t *testing.T) {
	fixture := newStateStoreFixture()
	fixture.TestSaveAndLoad(t)
}

func TestConsoleSimpleLogging(t *testing.T) {
	fixture := newStateStoreFixture()
	fixture.TestDelete(t)
}
