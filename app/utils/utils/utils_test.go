package utils

import (
	"net/url"
	"testing"
)

func TestGetPaginationParams(t *testing.T) {
	tests := []struct {
		name string
		args url.Values
		want []int
	}{
		{
			name: "page=3,limit=25",
			args: map[string][]string{
				"page":  []string{"3"},
				"limit": []string{"25"},
			},
			want: []int{25, 50},
		},
		{
			name: "page=0,limit=0",
			args: map[string][]string{
				"page":  []string{"0"},
				"limit": []string{"0"},
			},
			want: []int{20, 0},
		},
		{
			name: "no params",
			args: map[string][]string{},
			want: []int{20, 0},
		},
		{
			name: "malformed:page=-1,limit=1000",
			args: map[string][]string{
				"page":  []string{"-1"},
				"limit": []string{"1000"},
			},
			want: []int{20, 0},
		},
		{
			name: "malformed page:page=-1,limit=200",
			args: map[string][]string{
				"page":  []string{"-1"},
				"limit": []string{"200"},
			},
			want: []int{200, 0},
		},
		{
			name: "malformed limit:page=5,limit=500",
			args: map[string][]string{
				"page":  []string{"5"},
				"limit": []string{"500"},
			},
			want: []int{20, 80},
		},
		{
			name: "page=5,limit=5",
			args: map[string][]string{
				"page":  []string{"5"},
				"limit": []string{"5"},
			},
			want: []int{5, 20},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			limit, offset := GetPaginationParams(test.args)
			if limit != test.want[0] || offset != test.want[1] {
				t.Errorf(`
					Wanted: %v,
					Got: %v
				`, test.want, []int{limit, offset})
			}
		})
	}
}
