package filex

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dabao-zhao/ddltoall/version"
)

const (
	NL            = "\n"
	ddltomodelDir = "./ddltomodel"
)

var thisToolHome string

// LoadTemplate gets template content by the specified file.
func LoadTemplate(file, builtin string) (string, error) {
	dir, err := GetTemplateDir()
	if err != nil {
		return "", err
	}

	file = filepath.Join(dir, file)
	if !FileExists(file) {
		return builtin, nil
	}

	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// GetTemplateDir returns the path value in ToolHome where could get it by GetThisToolHome.
func GetTemplateDir() (string, error) {
	home, err := GetThisToolHome()
	if err != nil {
		return "", err
	}
	if home == thisToolHome {
		// backward compatible, it will be removed in the feature
		// backward compatible start.
		beforeTemplateDir := filepath.Join(home, version.GetThisToolVersion())
		fs, _ := os.ReadDir(beforeTemplateDir)
		var hasContent bool
		for _, e := range fs {
			info, err := e.Info()
			if err != nil {
				return "", err
			}
			if info.Size() > 0 {
				hasContent = true
			}
		}
		if hasContent {
			return beforeTemplateDir, nil
		}
		// backward compatible end.

		return filepath.Join(home), nil
	}

	return filepath.Join(home, version.GetThisToolVersion()), nil
}

func GetThisToolHome() (home string, err error) {
	defer func() {
		if err != nil {
			return
		}
		info, err := os.Stat(home)
		if err == nil && !info.IsDir() {
			_ = os.Rename(home, home+".old")
			_ = MkdirIfNotExist(home)
		}
	}()
	if len(thisToolHome) != 0 {
		home = thisToolHome
		return
	}
	home, err = GetDefaultToolHome()
	return
}

func GetDefaultToolHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ddltomodelDir), nil
}

// FileExists returns true if the specified file is exists.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// Clean deletes all templates and removes the parent directory.
func Clean() error {
	dir, err := GetTemplateDir()
	if err != nil {
		return err
	}
	return os.RemoveAll(dir)
}

// MustTempDir creates a temporary directory.
func MustTempDir() string {
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatalln(err)
	}

	return dir
}
