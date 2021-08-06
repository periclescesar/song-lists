package domain

type List struct {
	Name  string `json:"name"`
	Songs []int  `json:"songs"`
	Type  string `json:"type"`
}

const ListMaxSize = 5

func (s *List) IsValid() bool {
	return s.Name != "" && s.Type == "custom" && len(s.Songs) <= ListMaxSize
}
