package globalVars

import "errors"

var projectRootPath = struct {
	val   string
	isSet bool
}{
	val:   "",
	isSet: false,
}

func SetProjectRootPath(newVal string) error {
	if projectRootPath.isSet {
		return errors.New("projectRootPath is already set")
	}
	projectRootPath.val = newVal
	projectRootPath.isSet = true
	return nil
}

func GetProjectRootPath() (string, error) {
	if !projectRootPath.isSet {
		return "", errors.New("projectRootPath is not set")
	}
	return projectRootPath.val, nil
}
