package _type

var EmptyRelation *Relation = NewRelation(nil, nil)

type IRelation interface {
	GetGraph() *ECMASet
	GetX() *ECMASet
	GetY() *ECMASet
	GetDomain() *ECMASet
	GetCoDomain() *ECMASet
	IsHomogeneous() bool
	IsEmpty() bool
	IsUniversal() bool
	IsReflexive() bool
	IsSymmetric() bool
	IsAntiSymmetric() bool
	IsAsymmetric() bool
	IsTransitive() bool
	IsDense() bool
	IsConnected() bool
	IsTotal() bool
	IsStronglyConnected() bool
	IsStronglyTotal() bool
	IsTrichotomous() bool
	IsSubsetOf(s *Relation) bool
	I() *Relation
	R(x interface{}, y interface{}) bool
	Union(s *Relation) *Relation
	Intersection(s *Relation) *Relation
}

type Relation struct {
	graph       *ECMASet
	x           *ECMASet
	y           *ECMASet
	homogeneous bool
}

func NewRelation(x *ECMASet, y *ECMASet) *Relation {
	var set ECMASet = ECMASet{}
	if x != nil || y != nil {
		for _, v1 := range *x {
			for _, v2 := range *y {
				set.Add(NewPair(v1, v2))
			}
		}
	}
	return &Relation{
		graph:       &set,
		x:           x,
		y:           y,
		homogeneous: x == y,
	}
}

// GetGraph :
// Gets the graph pertaining to this relation.
func (r *Relation) GetGraph() *ECMASet {
	return r.graph
}

// GetX :
// Get the first set of a relation.
func (r *Relation) GetX() *ECMASet {
	return r.x
}

// GetY :
// Get the second set of a relation.
func (r *Relation) GetY() *ECMASet {
	return r.y
}

// GetDomain :
// Get the first set of a relation.
func (r *Relation) GetDomain() *ECMASet {
	return r.x
}

// GetCoDomain :
// Get the second set of a relation.
func (r *Relation) GetCoDomain() *ECMASet {
	return r.y
}

// IsHomogeneous :
// Checks if this relation is over two sets that are equal.
func (r *Relation) IsHomogeneous() bool {
	return r.homogeneous
}

// IsEmpty :
// Checks if a homogeneous relation is empty
func (r *Relation) IsEmpty() bool {
	return r.GetGraph() == nil || (r.IsHomogeneous() && r.graph.Size() <= 0)
}

// IsUniversal :
// Determines if this relation is a universal relation.
func (r *Relation) IsUniversal() bool {
	return r.IsHomogeneous()
}

// IsReflexive :
// Checks if this homogeneous relation is reflexive.
func (r *Relation) IsReflexive() bool {
	if !r.IsHomogeneous() {
		return false
	}

	return r.I().IsSubsetOf(r)
}

// IsSymmetric :
// Checks if this homogeneous relation is symmetric.
func (r *Relation) IsSymmetric() bool {
	if !r.IsHomogeneous() {
		return false
	}

	for _, x := range *r.GetX() {
		for _, y := range *r.GetX() {
			if r.R(x, y) && !r.R(y, x) {
				return false
			}
		}
	}
	return true
}

// IsAntiSymmetric :
// Checks if this homogeneous relation is antisymmetric.
func (r *Relation) IsAntiSymmetric() bool {
	if !r.IsHomogeneous() {
		return false
	}

	for _, x := range *r.GetX() {
		for _, y := range *r.GetX() {
			if r.R(x, y) && r.R(y, x) && x != y {
				return false
			}
		}
	}
	return true
}

// IsAsymmetric :
// Checks if this homogeneous relation is asymmetric.
func (r *Relation) IsAsymmetric() bool {
	return r.IsAntiSymmetric() && !r.IsReflexive()
}

// IsTransitive :
// Checks if this homogeneous relation is transitive.
func (r *Relation) IsTransitive() bool {
	if !r.IsHomogeneous() {
		return false
	}

	for _, x := range *r.GetX() {
		for _, y := range *r.GetX() {
			for _, z := range *r.GetX() {
				if r.R(x, y) && r.R(y, z) && !r.R(x, z) {
					return false
				}
			}
		}
	}
	return true
}

// IsDense :
// Checks if this homogeneous relation is dense.
func (r *Relation) IsDense() bool {
	if !r.IsHomogeneous() {
		return false
	}

	for _, x := range *r.GetX() {
		for _, y := range *r.GetX() {
			if r.R(x, y) {
				// find a z
				found := false
				for _, z := range *r.GetX() {
					if r.R(x, z) && r.R(z, y) {
						found = true
						break
					}
				}

				if !found {
					return false
				}
			}
		}
	}
	return true
}

// IsConnected :
// Checks if this homogeneous relation is connected (total).
func (r *Relation) IsConnected() bool {
	if !r.IsHomogeneous() {
		return false
	}

	for _, x := range *r.GetX() {
		for _, y := range *r.GetX() {
			if x != y && !r.R(x, y) && !r.R(y, x) {
				return false
			}
		}
	}
	return true
}

// IsTotal :
// Checks if this homogeneous relation is total (connected).
func (r *Relation) IsTotal() bool {
	return r.IsConnected()
}

// IsStronglyConnected :
// Checks if this homogeneous relation is strongly connected (total).
func (r *Relation) IsStronglyConnected() bool {
	if !r.IsHomogeneous() {
		return false
	}

	for _, x := range *r.GetX() {
		for _, y := range *r.GetX() {
			if !r.R(x, y) && !r.R(y, x) {
				return false
			}
		}
	}
	return true
}

// IsStronglyTotal :
// Checks if this homogeneous relation is strongly total (connected).
func (r *Relation) IsStronglyTotal() bool {
	return r.IsStronglyConnected()
}

// IsTrichotomous :
// Checks if this homogeneous relation is trichotomous.
func (r *Relation) IsTrichotomous() bool {
	if !r.IsHomogeneous() {
		return false
	}

	for _, x := range *r.GetX() {
		for _, y := range *r.GetX() {
			if !r.R(x, y) && !r.R(y, x) && x != y {
				return false
			}
		}
	}
	return true
}

// IsSubsetOf :
// Checks if this relation is a subset of another.
func (r *Relation) IsSubsetOf(s *Relation) bool {
	for _, v := range *r.GetGraph() {
		if !s.GetGraph().Has(v) {
			return false
		}
	}
	return true
}

// I :
// Gets the identity relation for this relation.
func (r *Relation) I() *Relation {
	var set ECMASet = ECMASet{}
	if !r.IsHomogeneous() {
		return nil
	}

	for _, v := range *r.x {
		set.Add(NewPair(v, v))
	}

	return &Relation{
		graph:       &set,
		x:           r.x,
		y:           r.y,
		homogeneous: r.homogeneous,
	}
}

// R :
// Checks if two elements are related to one another per this relation.
func (r *Relation) R(x interface{}, y interface{}) bool {
	p := NewPair(x, y)
	return r.graph.Has(p)
}

// Union :
// Gets the Union of two relations.
func (r *Relation) Union(s *Relation) *Relation {
	// Ensure s is also a relation over sets X and Y
	if r.GetX() != s.GetX() || r.GetY() != s.GetY() {
		return nil
	}

	var set ECMASet = ECMASet{}
	for _, v1 := range *(r.GetGraph()) {
		set.Add(v1)
	}

	for _, v2 := range *(s.GetGraph()) {
		set.Add(v2)
	}

	return &Relation{
		graph:       &set,
		x:           r.x,
		y:           r.y,
		homogeneous: r == s,
	}
}

// Intersection :
// Gets the Intersection of two relations.
func (r *Relation) Intersection(s *Relation) *Relation {
	// Ensure s is also a relation over sets X and Y
	if r.GetX() != s.GetX() || r.GetY() != s.GetY() {
		return nil
	}

	var set ECMASet = ECMASet{}
	for _, v1 := range *(r.GetGraph()) {
		if s.GetGraph().Has(v1) {
			set.Add(v1)
		}
	}

	for _, v2 := range *(s.GetGraph()) {
		if r.GetGraph().Has(v2) {
			set.Add(v2)
		}
	}

	return &Relation{
		graph:       &set,
		x:           r.x,
		y:           r.y,
		homogeneous: r == s,
	}
}

// Composition :
// Gets the Composition of two relations, of order s ° r.
func (r *Relation) Composition(s *Relation) *Relation {
	// Ensure r is a relation over X and Y and s is a relation over Y and Z
	if r.GetY() != s.GetX() {
		return nil
	}
	return nil
}
