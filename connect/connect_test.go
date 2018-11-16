package connect

import "testing"

func TestConnectionParamsCreating(t *testing.T) {
	t.Log("Try to create new ConnectionParams")
	params := NewEmptyConnectionParams()
	t.Log(params)
}
