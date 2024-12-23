package utils

import (
	"reflect"
	"testing"
)

func TestParseIntLines(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    [][]int
		wantErr bool
	}{
		{
			name:    "single line single digit",
			input:   "7",
			want:    [][]int{{7}},
			wantErr: false,
		},
		{
			name:    "single line multiple digits",
			input:   "1234",
			want:    [][]int{{1, 2, 3, 4}},
			wantErr: false,
		},
		{
			name:    "multiple lines with digits",
			input:   "12\n34",
			want:    [][]int{{1, 2}, {3, 4}},
			wantErr: false,
		},
		{
			name:    "lines with empty line in between",
			input:   "12\n\n34",
			want:    [][]int{{1, 2}, {3, 4}},
			wantErr: false,
		},
		{
			name:    "trailing newline",
			input:   "12\n34\n",
			want:    [][]int{{1, 2}, {3, 4}},
			wantErr: false,
		},
		{
			name:    "large input",
			input:   "9876543210\n1234567890",
			want:    [][]int{{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseIntLines(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIntLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseIntLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}
