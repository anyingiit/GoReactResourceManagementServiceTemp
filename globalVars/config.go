package globalVars

import (
	"errors"

	"github.com/anyingiit/GoReactResourceManagement/structs"
)

var projectConfig = struct {
	val   *structs.Config
	isSet bool
}{
	val:   nil,
	isSet: false,
}

func SetProjectConfig(newVal *structs.Config) error {
	if projectConfig.isSet {
		return errors.New("projectRootPath is already set")
	}
	projectConfig.val = newVal
	projectConfig.isSet = true
	return nil
}

func GetProjectConfig() (*structs.Config, error) {
	if !projectConfig.isSet {
		return nil, errors.New("projectRootPath is not set")
	}
	return projectConfig.val, nil
}
