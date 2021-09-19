package avlog_test

import (
	"avunico/avlog"
	"testing"
)

func TestIniciarLogger(t *testing.T) {
	ob := avlog.AvLog{}

	ob.IniciarLogger()
}

func TestInfo(t *testing.T) {

	ob := avlog.AvLog{}

	ob.IniciarLogger()

	if ob.Info() == nil {
		t.Error("Erro na chamada de Info")
	}
}

func TestError(t *testing.T) {

	ob := avlog.AvLog{}

	ob.IniciarLogger()

	if ob.Error() == nil {
		t.Error("Erro na chamada de Info")
	}
}

func TestWarning(t *testing.T) {

	ob := avlog.AvLog{}

	ob.IniciarLogger()

	if ob.Warning() == nil {
		t.Error("Erro na chamada de Info")
	}
}
