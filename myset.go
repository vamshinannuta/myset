package myset

type Set map[int]bool

func (s Set) Union(s2 Set) Set {
	for k, _ := range s2 {
		if _, ok := s[k]; !ok {
			s[k] = true
		}
	}
	return s
}

func (s Set) Intersection(s2 Set) Set {
	for k, _ := range s {
		if _, ok := s2[k]; !ok {
			delete(s, k)
		}
	}
	return s
}

func (s Set) Difference(s2 Set) Set {
	for k, _ := range s {
		if _, ok := s2[k]; ok {
			delete(s, k)
		}
	}
	return s
}

