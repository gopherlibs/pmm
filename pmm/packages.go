package pmm

type PkgInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Manager string `json:"manager"`
}

func PkgMissingFromA(listA, listB []PkgInfo) []PkgInfo {

	listAMap := make(map[string]PkgInfo, len(listA))

	bOnlyPkgs := []PkgInfo{}

	for _, v := range listA {
		listAMap[v.Name] = v
	}

	for _, v := range listB {

		if _, found := listAMap[v.Name]; !found {
			bOnlyPkgs = append(bOnlyPkgs, v)
		}
	}

	return bOnlyPkgs
}
