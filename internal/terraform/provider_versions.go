package terraform

type ProviderVersions []*ProviderVersion

func (p ProviderVersions) Last() *ProviderVersion {
	return p[len(p)-1]
}

// Implementing sort.Interface based on the Version field

func (p ProviderVersions) Len() int {
	return len(p)
}

func (p ProviderVersions) Less(i, j int) bool {
	pi, _ := p[i].SemanticVersion()
	pj, _ := p[j].SemanticVersion()
	return pi.LessThan(pj)
}

func (p ProviderVersions) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
