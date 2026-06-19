package guard

import (
	"bytes"
	"testing"
)

func TestPrintReportText_WritesRepositoryRoot(t *testing.T) {
	buf := bytes.Buffer{}
	report := DiagnosticReport{
		RepositoryRoot: "H:\\source\\Reze-go",
	}
	err := PrintReportText(&buf, report)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() != "Repository: H:\\source\\Reze-go\n" {
		t.Errorf("got %s, want \"Repository: H:\\source\\Reze-go\"\n", buf.String())
	}
}
