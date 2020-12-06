package pmm

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

type rpmManager struct {
}

func (this *rpmManager) List() []PkgInfo {

	var packages []PkgInfo

	output, err := exec.Command("rpm", "-qa", "--qf", "%{NAME}\t%{VERSION}\n").Output()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(output), "\n")
	lines = lines[0 : len(lines)-1]

	for _, pkg := range lines {

		pkgSplit := strings.Split(pkg, "\t")
		packages = append(packages, PkgInfo{
			Name:    pkgSplit[0],
			Version: pkgSplit[1],
			Manager: "rpm",
		})
	}

	return packages
}
