package pmm

import "github.com/arduino/go-apt-client"

type aptManager struct {
}

func (this *aptManager) List() []pkgInfo {

	var packages []pkgInfo

	allPackages, _ := apt.List()

	for _, pkg := range allPackages {

		if pkg.Status == "installed" {

			packages = append(packages, pkgInfo{
				Name:    pkg.Name,
				Version: pkg.Version,
				Manager: "apt",
			})
		}
	}

	return packages
}
