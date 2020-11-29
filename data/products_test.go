package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProducts_TestCheckValidations(t *testing.T) {
	type args struct {
		product Product
	}
	type test struct {
		name    string
		args    args
		wantErr int
	}

	tests := []test{
		{
			name: "A proper product",
			args: args{
				Product{
					Name:  "Tidal wave",
					Price: 1,
					SKU:   "abc-abc-abc",
				},
			},
			wantErr: 0,
		},
		{
			name: "Missing name",
			args: args{
				Product{
					Description: "bla bla bla",
					Price:       1,
					SKU:         "abc-abc-abc",
				},
			},
			wantErr: 1,
		},
		{
			name: "Missing price",
			args: args{
				Product{
					Name:        "Tidal wave",
					Description: "bla bla bla",
					SKU:         "abc-abc-abc",
				},
			},
			wantErr: 1,
		},
		{
			name: "Wrong price range",
			args: args{
				Product{
					Name:        "Tidal wave",
					Description: "bla bla bla",
					Price:       -1,
					SKU:         "abc-abc-abc",
				},
			},
			wantErr: 1,
		},
		{
			name: "Invalid SKU",
			args: args{
				Product{
					Name:        "Tidal wave",
					Description: "bla bla bla",
					Price:       1,
					SKU:         "abc-123",
				},
			},
			wantErr: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidation()
			errs := validator.Validate(tt.args.product)
			if tt.wantErr != 0 {
				assert.Len(t, errs, tt.wantErr)
				return
			}

			assert.Len(t, errs, 0)
		})
	}
}

func TestJson_ToJSON(t *testing.T) {
	type args struct {
		product Product
	}
	type test struct {
		name    string
		args    args
		wantErr error
	}

	tests := []test{
		{
			name: "Sucessfully convert to JSON",
			args: args{
				Product{
					Name: "Tidal wave",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := bytes.NewBufferString("")
			err := ToJSON(tt.args.product, writer)
			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
			}

			assert.NoError(t, err)
		})
	}
}
