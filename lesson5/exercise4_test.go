package lesson5

import "testing"

func TestDeleteTowels(t *testing.T) {
	cases := []struct {
		name string
		value string
		want string
	}{
		{
			name: "not small",
			value: "destructor ai",
			want: "dstrctr ",
		},
		{
			name: "not big",
			value: "DESTRUCTOR AI",
			want: "DSTRCTR ",
		},
		{
			name: "mixed",
			value: "DESTRuctor Ai",
			want: "DSTRctr ",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := DeleteVowels(tc.value)
			if got != tc.want {
				t.Errorf(`DeleteVowels(%s)=%s, want %s`, tc.value, got, tc.want)
			}
		})
	}
}