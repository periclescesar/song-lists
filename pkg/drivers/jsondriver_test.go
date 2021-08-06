package drivers

import (
	"GoSOLID/pkg/domain"
	"testing"
)

func Test_open(t *testing.T) {
	tests := []struct {
		name     string
		json_str interface{}
		valid    bool
	}{
		{"write on json", domain.List{Name: "exemple", Songs: []int{1, 2, 3}, Type: "custom"}, true},
	}
	for _, tt := range tests {
		jd := JsonDriver{filepath: "test.json"}

		t.Run(tt.name, func(t *testing.T) {
			err := jd.Write(tt.json_str)
			if err != nil && tt.valid {
				t.Errorf("Error writing json file")
			}
		})

		_ = jd.DeleteFile()
	}
}
