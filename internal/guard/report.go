package guard

import "io"

func PrintReportText(w io.Writer, report DiagnosticReport) error {
	_, err := w.Write([]byte("Repository: " + report.RepositoryRoot + "\n"))
	if err != nil {
		return err
	}
	return nil
}
