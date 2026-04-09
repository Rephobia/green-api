package validatorx_test

import (
	"testing"

	"github.com/Rephobia/green-api-test-task/internal/validatorx"
	"github.com/stretchr/testify/require"
)

type TestDTO struct {
	Name string `validate:"required"`
}

func TestValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   TestDTO
		wantErr bool
	}{
		{
			name:    "valid dto",
			input:   TestDTO{Name: "ok"},
			wantErr: false,
		},
		{
			name: "missing name",
			//nolint:exhaustruct // проверка на пустое поле
			input:   TestDTO{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			errs := validatorx.Validate(tt.input)

			if tt.wantErr {
				require.NotEmpty(t, errs)
			} else {
				require.Empty(t, errs)
			}
		})
	}
}
