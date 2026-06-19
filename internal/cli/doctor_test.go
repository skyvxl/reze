package cli

import (
	"context"
	"errors"
	"testing"

	"github.com/skyvxl/reze/internal/gitx"
)

func TestHumanizeGitError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "git not found",
			err:  gitx.ErrGitNotFound,
			want: "git not found: install Git and make sure it is available in PATH",
		},
		{
			name: "not a git repository",
			err:  gitx.ErrNotGitRepository,
			want: "current directory is not a git repository",
		},
		{
			name: "operation canceled",
			err:  context.Canceled,
			want: "operation canceled",
		},
		{
			name: "other error",
			err:  errors.New("boom"),
			want: "boom",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := HumanError(tt.err)
			if got == nil {
				t.Fatal("expected error, got nil")
			}
			if got.Error() != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got.Error())
			}
		})
	}
}
