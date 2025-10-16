package core

type Mapping[E comparable] struct {
	New E
	Old E
}

type MappingList [Mapping[E]]struct {
	Mappings []Mapping[E]
	Index    map[E]E
}
