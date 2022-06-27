package embedcontents

import (
	"embed"
	"errors"
	"io/ioutil"
	"os"
	"path"
)

//go:embed assets
var Assets embed.FS

func CreateAssetsFolder(dirName string) error {
	workingDir, wdErr := os.Getwd()
	if wdErr != nil {
		return wdErr
	}
	BuildPath := path.Join(workingDir, dirName)
	if _, err := os.Stat(BuildPath); errors.Is(err, os.ErrNotExist) {

		//create apidoc Build Folder
		createErr := os.Mkdir(BuildPath, 0777)
		if createErr != nil {
			return createErr
		}
	}

	//write all files from embed to apidoc build folder
	entries, embedReadDirErr := Assets.ReadDir("assets")
	if embedReadDirErr != nil {
		return embedReadDirErr
	}
	for _, entery := range entries {
		data, readFileErr := Assets.ReadFile(path.Join("assets", entery.Name()))
		if readFileErr != nil {
			return readFileErr
		}
		ioutil.WriteFile(path.Join(BuildPath, entery.Name()), data, 0664)
	}

	return nil
}
