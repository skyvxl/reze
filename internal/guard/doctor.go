package guard

import (
	"context"
	"fmt"
)

type DoctorOptions struct {
	StartDir string
}

type RepositoryReader interface {
	Discover(ctx context.Context, startDir string) (Repository, error)
}

type Doctor struct {
	RepositoryReader
}

func NewDoctor(repositories RepositoryReader) *Doctor {
	return &Doctor{
		RepositoryReader: repositories,
	}
}

func (d *Doctor) Run(ctx context.Context, opts DoctorOptions) (DiagnosticReport, error) {
	repo, err := d.Discover(ctx, opts.StartDir)
	if err != nil {
		return DiagnosticReport{}, fmt.Errorf("discover repository: %w", err)
	}
	return DiagnosticReport{
		RepositoryRoot: repo.Root,
	}, nil
}
