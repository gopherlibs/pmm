package pmm

import (
	"encoding/json"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type pipManager struct {
	commands []string
}

func (this *pipManager) List() []pkgInfo {

	var packages []pkgInfo
	var pipJSON []map[string]string

	for _, pipCmd := range []string{"pip", "pip3"} {

		if !commandExists(pipCmd) {
			continue
		}

		output, err := exec.Command(pipCmd, "list", "--format=json").Output()
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(output, &pipJSON)
		if err != nil {
			log.Fatal(err)
		}

		for _, pkg := range pipJSON {
			packages = append(packages, pkgInfo{
				Name:    pkg["name"],
				Version: pkg["version"],
				Manager: "pip",
			})
		}
	}

	return packages
}
