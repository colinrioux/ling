package _type

type IECMASet interface {
	Add(e interface{}) bool
	Remove(e interface{}) bool
	Has(e interface{}) bool
	Find(e interface{}) int
	Size() int
	Union(other *ECMASet) *ECMASet
	Intersection(other *ECMASet) *ECMASet
	Difference(other *ECMASet) *ECMASet
}

type ECMASet []interface{}

// Add :
// Adds an element to the graph if it doesn't exist in the graph.
func (s *ECMASet) Add(e interface{}) bool {
	if s.Has(e) {
		return false
	}
	*s = append(*s, e)
	return true
}

// Remove :
// Removes an element from the graph if it exists.
func (s *ECMASet) Remove(e interface{}) bool {
	loc := s.Find(e)
	if loc == -1 {
		return false
	}
	*s = append((*s)[:loc], (*s)[loc+1:]...)
	return true
}

// Has :
// Checks if a graph has an element.
func (s *ECMASet) Has(e interface{}) bool {
	for _, v := range *s {
		if v == e {
			return true
		}
	}
	return false
}

// Find :
// Get the position of an element in a graph (-1 if not exists).
func (s *ECMASet) Find(e interface{}) int {
	for i, v := range *s {
		if v == e {
			return i
		}
	}
	return -1
}

// Size :
// Gets the number of elements in the graph.
func (s *ECMASet) Size() int {
	return len(*s)
}

// Union :
// Return a new graph who is the union of two sets.
func (s *ECMASet) Union(other *ECMASet) *ECMASet {
	var newSet ECMASet = make(ECMASet, s.Size())
	copy(newSet, *s)
	for _, v := range *other {
		newSet.Add(v)
	}
	return &newSet
}

// Intersection :
// Return a new graph who is the intersection of two sets.
func (s *ECMASet) Intersection(other *ECMASet) *ECMASet {
	var newSet ECMASet = ECMASet{}
	for _, v := range *s {
		idx := other.Find(v)
		if idx > -1 {
			newSet = append(newSet, v)
		}
	}
	return &newSet
}

// Difference :
// Return a new graph who is the difference of two sets.
func (s *ECMASet) Difference(other *ECMASet) *ECMASet {
	var newSet ECMASet = make(ECMASet, s.Size())
	copy(newSet, *s)
	for i, v := range newSet {
		if other.Has(v) {
			newSet = append(newSet[:i], newSet[i+1:]...)
		}
	}
	return &newSet
}
