package path

import "testing"

func Test_Ext(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get extension",
			args: args{
				path: "/test/path/to/sample.txt",
			},
			want: ".txt",
		},
		{
			name: "get extension: file is hidden one with extension",
			args: args{
				path: "/test/path/to/.sample.txt",
			},
			want: ".txt",
		},
		{
			name: "not get extension: no extension in path",
			args: args{
				path: "/test/path/to/sample",
			},
			want: "",
		},
		{
			name: "not get extension: file is hidden one without extension",
			args: args{
				path: "/test/path/to/.sample",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ext(tt.args.path); got != tt.want {
				t.Errorf("Ext() = %v, want %v", got, tt.want)
			}
		})
	}
}
