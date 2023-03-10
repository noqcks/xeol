package match

import "sort"

var _ sort.Interface = (*ByElements)(nil)

type ByElements []Match

// Len is the number of elements in the collection.
func (m ByElements) Len() int {
	return len(m)
}

// Less reports whether the element with index i should sort before the element with index j.
func (m ByElements) Less(i, j int) bool {
	if m[i].Cycle.ProductName == m[j].Cycle.ProductName {
		if m[i].Cycle.ReleaseCycle == m[j].Cycle.ReleaseCycle {
			if m[i].Package.Name == m[j].Package.Name {
				if m[i].Package.Version == m[j].Package.Version {
					return m[i].Package.Type < m[j].Package.Type
				}
				return m[i].Package.Version < m[j].Package.Version
			}
			return m[i].Package.Name < m[j].Package.Name
		}
		return m[i].Cycle.ReleaseCycle < m[j].Cycle.ReleaseCycle
	}
	return m[i].Cycle.ProductName < m[j].Cycle.ProductName
}

// Swap swaps the elements with indexes i and j.
func (m ByElements) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
