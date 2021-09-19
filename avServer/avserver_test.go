package avServer

import "testing"

func TestNew(t *testing.T) {
	ob := New(nil, nil)

	if ob == nil {
		t.Error("Erro no teste de criação do objeto server")
	}
}

func TestValidarRequest(t *testing.T) {
	ValidarRequest("/api/v1/feiralivre/1")
}
