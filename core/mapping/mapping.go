package core

type Mapping[T comparable] struct {
	New T
	Old T
}

type MappingList[T comparable] struct {
	Mappings []Mapping[T] `json:"mappings,omitempty"`
	Index    map[T]T      `json:"index,omitempty"`
}

func NewMappingList[T comparable](mappings []Mapping[T]) *MappingList[T] {
	index := make(map[T]T, len(mappings))
	for _, mapping := range mappings {
		index[mapping.Old] = mapping.New
	}
	return &MappingList[T]{Mappings: mappings, Index: index}
}

func (mappingList *MappingList[T]) Apply(value T) T {
	if newValue, ok := mappingList.Index[value]; ok {
		return newValue
	}
	return value
}
