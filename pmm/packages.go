package pmm

type pkgInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Manager string `json:"manager"`
}

func PkgMissingFromA(listA, listB []pkgInfo) []pkgInfo {

	listAMap := make(map[string]pkgInfo, len(listA))

	bOnlyPkgs := []pkgInfo{}

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
