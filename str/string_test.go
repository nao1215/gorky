// Package str implement string manipulations not provided by the golang standard package (strings package)
package str

import (
	"testing"
)

func TestTrimGaps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Before:' Hello,    World  ! ', After:'Hello, World !'",
			args: args{
				" Hello,    World  ! ",
			},
			want: "Hello, World !",
		},
		{
			name: "Before:' Hello,    World  ! ', After:'Hello, World !'",
			args: args{
				" Hello,    World  ! ",
			},
			want: "Hello, World !",
		},
		{
			name: "Before:' \t\n\t Hello, \n\t World \n ! \n\t ', After:'Hello, World !'",
			args: args{
				" \t\n\t Hello, \n\t World \n ! \n\t ",
			},
			want: "Hello, World !",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimGaps(tt.args.s); got != tt.want {
				t.Errorf("TrimGaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		list []string
		s    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "slice has target string",
			args: args{
				list: []string{"a", "bb", "metallica", "abc"},
				s:    "metallica",
			},
			want: true,
		},
		{
			name: "slice does not have target string",
			args: args{
				list: []string{"a", "bbb", "abc"},
				s:    "metallica",
			},
			want: false,
		},
		{
			name: "slice is empty",
			args: args{
				list: []string{},
				s:    "metallica",
			},
			want: false,
		},
		{
			name: "target string is empty string",
			args: args{
				list: []string{"a", "bbb", "abc"},
				s:    "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.list, tt.args.s); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
