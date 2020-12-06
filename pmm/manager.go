package pmm

import (
	"errors"
	"fmt"
	"os/exec"
)

type ManagerType uint

const (
	TypeAPT ManagerType = iota
	TypeRPM
	TypePIP
	TypeAPK
)

func (i ManagerType) String() string {
	return [...]string{"Apt", "RPM", "Pip", "Apk"}[i]
}

type manager interface {
	//Available() []pkgInfo
	List() []PkgInfo
}

type managerNotAvailableError struct {
	mType   ManagerType
	command string
}

func (this *managerNotAvailableError) Error() string {
	return fmt.Sprintf("Error: Can't init %s due to the following commands not being available on the local machine: %v", this.mType, this.command)
}

func New(mType ManagerType) (manager, error) {

	// make sure these binaries are available
	// this list is what is used by github.com/arduino/go-apt-client
	mCommands := map[ManagerType][]string{
		TypeAPK: []string{
			"apk",
		},
		TypeAPT: []string{
			"dpkg-query",
			"apt",
			"apt-get",
		},
		TypePIP: []string{
			"pip",
			"pip3",
		},
		TypeRPM: []string{
			"rpm",
		},
	}

	for _, command := range mCommands[mType] {
		if !commandExists(command) {
			return nil, &managerNotAvailableError{mType, command}
		}
	}

	switch mType {
	case TypeAPK:
		return &apkManager{}, nil
	case TypeAPT:
		return &aptManager{}, nil
	case TypePIP:
		return &pipManager{mCommands[mType]}, nil
	case TypeRPM:
		return &rpmManager{}, nil
	}

	return nil, errors.New("Error: invalid package manager type.")
}

func commandExists(command string) bool {

	_, err := exec.LookPath(command)

	return err == nil
}
