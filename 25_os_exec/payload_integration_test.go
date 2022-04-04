package os_exec_test

import (
	"testing"

	payload "github.com/ichang0301/learn-golang/25_os_exec"
)

func TestGetDataIntegration(t *testing.T) {
	got := payload.GetData(payload.GetXMLFromCommand("cat", "msg.xml"))
	want := "HAPPY NEW YEAR!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
