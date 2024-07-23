// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package errno_test

import (
	"errors"
	"testing"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/DataDog/documentor/internal/errno"
)

const ErrGenericError xerrors.Error = "failure"

func TestError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "With underlying error",
			err:  ErrGenericError,
			want: "failure",
		},
		{
			name: "With nil underlying error",
			err:  nil,
			want: "unknown error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := errno.Error{
				Err: tt.err,
			}

			if got := err.Error(); got != tt.want {
				t.Errorf("Error.Error() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestError_Unwrap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want error
	}{
		{
			name: "With underlying error",
			err:  ErrGenericError,
			want: ErrGenericError,
		},
		{
			name: "With nil underlying error",
			err:  nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := errno.Error{
				Err: tt.err,
			}

			if got := err.Unwrap(); !errors.Is(got, tt.want) {
				t.Errorf("Error.Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Code(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		no   uint8
		want int
	}{
		{
			name: "ExitOK",
			no:   errno.ExitOK,
			want: 0,
		},
		{
			name: "ExitError",
			no:   errno.ExitError,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := errno.Error{
				No: tt.no,
			}

			if got := err.Code(); got != tt.want {
				t.Errorf("Error.Code() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		code     uint8
		err      error
		wantCode int
		wantErr  error
	}{
		{
			name:     "With nil error",
			code:     errno.ExitError,
			err:      nil,
			wantCode: 1,
			wantErr:  nil,
		},
		{
			name:     "With non-nil error",
			code:     errno.ExitError,
			err:      ErrGenericError,
			wantCode: 1,
			wantErr:  ErrGenericError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := errno.New(tt.code, tt.err)

			if got.Code() != tt.wantCode || !errors.Is(got.Unwrap(), tt.wantErr) {
				t.Errorf("New(%d, %v) = {Code: %d, Err: %v}, want {Code: %d, Err: %v}", tt.code, tt.err, got.Code(), got.Unwrap(), tt.wantCode, tt.wantErr)
			}
		})
	}
}
