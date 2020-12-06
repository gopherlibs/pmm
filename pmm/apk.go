package pmm

import (
	"bufio"
	"os"
	"strings"
)

type apkManager struct {
}

func (this *apkManager) List() []PkgInfo {

	var packages []PkgInfo

	file, err := os.Open("/lib/apk/db/installed")
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		if strings.HasPrefix(line, "P:") {

			pkgName := line[2:len(line)]
			scanner.Scan()
			line := scanner.Text()
			pkgVersion := line[2:len(line)]

			packages = append(packages, PkgInfo{
				Name:    pkgName,
				Version: pkgVersion,
				Manager: "apk",
			})
		}

	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return packages
}
