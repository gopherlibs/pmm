package pmm

import "github.com/arduino/go-apt-client"

type aptManager struct {
}

func (this *aptManager) List() []PkgInfo {

	var packages []PkgInfo

	allPackages, _ := apt.List()

	for _, pkg := range allPackages {

		if pkg.Status == "installed" {

			packages = append(packages, PkgInfo{
				Name:    pkg.Name,
				Version: pkg.Version,
				Manager: "apt",
			})
		}
	}

	return packages
}
