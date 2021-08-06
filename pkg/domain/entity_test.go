package domain

import "testing"

func TestSong_isValid(t *testing.T) {
	type fields struct {
		Name  string
		Songs []int
		Type  string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		//{"valid", fields{"list", []int{1,2,3}, "custom"}, true},
		{"empty name", fields{"", []int{1, 2, 3}, "custom"}, false},
		//{"invalid type", fields{"list", []int{1,2,3}, "favorites"}, false},
		//{"invalid max size", fields{"list", []int{1,2,3,4,5,6}, "custom"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &List{
				Name:  tt.fields.Name,
				Songs: tt.fields.Songs,
				Type:  tt.fields.Type,
			}
			if got := s.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
