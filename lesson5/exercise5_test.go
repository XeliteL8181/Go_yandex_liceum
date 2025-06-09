package lesson5

import "testing"

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		name string
		input []byte
		err error
		expected int
	} {
		{
			name: "Empty",
			input: []byte(""),
			err: nil,
			expected: 0,
		},
		{
			name: "Simple ASCII string",
			input: []byte("hello"),
			expected: 5,
			err: nil,
		},
		{
			name: "UTF-8 string with multi-byte characters",
			input: []byte("привет"),
			expected: 6,
			err: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			length, err := GetUTFLength(tt.input)
			if length != tt.expected {
				t.Errorf(`GetUTFLength(%q)=%d, want %d`, tt.input, length, tt.expected)
			}
			if err != tt.err {
				t.Errorf(`GetUTFLength(%q)=%v, want %v`, tt.input, err, tt.err)
			} 
		})
	}
}