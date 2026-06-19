package guard

import (
	"context"
	"errors"
	"testing"
)

type fakeRepositoryReader struct {
	repo     Repository
	err      error
	startDir string
}

func (f *fakeRepositoryReader) Discover(ctx context.Context, startDir string) (Repository, error) {
	f.startDir = startDir
	if f.err != nil {
		return Repository{}, f.err
	}
	return f.repo, nil
}

func TestDoctor_Run_ReturnsRepositoryRoot(t *testing.T) {
	fakeReader := &fakeRepositoryReader{
		repo: Repository{
			Root: "H:\\source\\Reze-go",
		},
	}
	doctor := NewDoctor(fakeReader)
	opts := DoctorOptions{StartDir: "."}
	report, err := doctor.Run(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	if report.RepositoryRoot != fakeReader.repo.Root {
		t.Errorf("got %q, want %q", report.RepositoryRoot, fakeReader.repo.Root)
	}
	if fakeReader.startDir != opts.StartDir {
		t.Errorf("got %q, want %q", fakeReader.startDir, opts.StartDir)
	}
}

func TestDoctor_Run_WrapsDiscoverError(t *testing.T) {
	discoverErr := errors.New("discover error")
	fakeReader := &fakeRepositoryReader{
		repo: Repository{
			Root: "H:\\source\\Reze-go",
		},
		err: discoverErr,
	}
	doctor := NewDoctor(fakeReader)
	opts := DoctorOptions{StartDir: "."}
	_, err := doctor.Run(context.Background(), opts)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, discoverErr) {
		t.Errorf("got error %v, want %v", err, discoverErr)
	}
}
