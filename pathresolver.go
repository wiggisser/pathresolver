//Package pathresolver provides simple access to uses home directory on linux, macOS and Windows systems
package pathresolver

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/mitchellh/go-homedir"

	"path/filepath"
)

var (
	basedirectory string
	isinit        bool = false
)

//Init initializes the pathresolver with a path for the respective system.
//If only one of the parameters is set, it is used of all systems
//If both parameters are set, the basedirectory is set according to the current system
//If the given paths are relative, they are set relative to the users homedir (~  and %HOMEDIR% respectively)
func Init(unix string, windows string) error {
	isinit = false
	path := ""
	if unix == "" && windows == "" {
		return fmt.Errorf("either unixbase or windowsbase must be set")
	} else if unix == "" {
		path = windows
	} else if windows == "" {
		path = unix
	} else if strings.Contains(runtime.GOOS, "windows") {
		path = windows
	} else {
		path = unix
	}

	if filepath.IsAbs(path) {
		basedirectory = path
		isinit = true
		return nil
	}

	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	basedirectory = filepath.Join(home, path)
	isinit = true
	return nil
}

//Path resolves a subpath relative to the basedir
//If subpath is empty, basedir is returned
//Subpath must not be an absolute path
func Path(subpath string) (string, error) {
	if !isinit {
		return "", fmt.Errorf("pathresolver not initialized")
	}

	if subpath == "" {
		return basedirectory, nil
	}

	if filepath.IsAbs(subpath) {
		return "", fmt.Errorf("cannot use absolute path as subpath")
	}

	return filepath.Join(basedirectory, subpath), nil
}
