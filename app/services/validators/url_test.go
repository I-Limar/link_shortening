package validators

import "testing"

func TestIsValidUrl(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "just string",
			args: "qweqweqwe",
			want: false,
		},
		{
			name: "invalid tld",
			args: "vk.comm",
			want: false,
		},
		{
			name: "without protocol",
			args: "vk.com",
			want: false,
		},
		{
			name: "without tld",
			args: "http://vk",
			want: false,
		},
		{
			name: "valid url",
			args: "https://vk.com",
			want: true,
		},
		{
			name: "valid url with path",
			args: "https://vk.com/asdasdasd/asdasdasd/asdasd",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUrl(tt.args); got != tt.want {
				t.Errorf("IsValidUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
