package repositories

import (
	"GoSOLID/pkg/domain"
	"reflect"
	"testing"
)

func TestJsonListRepositories_CreateList(t *testing.T) {
	tests := []struct {
		name    string
		list    *domain.List
		wantErr bool
	}{
		{"write on json", &domain.List{Name: "exemple", Songs: []int{1, 2, 3}, Type: "custom"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JsonListRepository{}
			if err := j.CreateList(tt.list); (err != nil) != tt.wantErr {
				t.Errorf("CreateList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJsonListRepositories_FetchLists(t *testing.T) {
	tests := []struct {
		name string
		want []domain.List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JsonListRepository{}
			if got := j.FetchLists(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestNewJsonListRepository(t *testing.T) {
//	tests := []struct {
//		name string
//		want *JsonListRepository
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewJsonListRepository(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewJsonListRepository() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
