package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Bhargav16exd/nginxctl/constants"

	_ "embed"
)

//go:embed default.conf
var defaultConf []byte

func CheckNginxInstallation() {

	_, err := exec.LookPath(constants.NGINX)

	if err != nil {
		log.Fatal(err)
	}

}

func FetchNginxConfPath() (string, error) {

	var arg = constants.CONF_PATH_ARG
	var path string = ""

	output, err := exec.Command(constants.NGINX, "-V").CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fields := strings.FieldsSeq(string(output))

	for value := range fields {
		if strings.HasPrefix(value, arg) {
			path = strings.TrimPrefix(value, arg)
		}
	}

	if path == "" {
		return "", errors.New("Path Not Found")
	}

	return path, nil
}

func CheckCreateSitesDir(path string) (string, string) {

	trimmedPath, exist := strings.CutSuffix(path, constants.CONF_FILE_NAME)

	if exist == false {
		fmt.Println("Package Error Contact Administrator")
	}

	sitesAvailableDir := trimmedPath + constants.DIR_SITES_AVAILALBE
	sitesEnabledDir := trimmedPath + constants.DIR_SITES_ENABLED

	var configPaths = []string{sitesAvailableDir, sitesEnabledDir}

	for _, path := range configPaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.Mkdir(path, os.ModePerm); err != nil {
				log.Fatal("Package Error", err)
			}
		}
	}

	return sitesAvailableDir, sitesEnabledDir
}

func ResetConfigration(configPath string, sitesAvailablePath string, sitesEnabledPath string) bool {

	//delete old file
	if _, err := os.Stat(configPath); err == nil {
		err := os.Remove(configPath)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}

	err := os.WriteFile(configPath, defaultConf, 0644)

	if err != nil {
		fmt.Println(err)
		return false
	}

	//delete
	var configPaths = []string{sitesAvailablePath, sitesEnabledPath}

	for _, path := range configPaths {

		//if file does not exist
		if _, err := os.Stat(path); os.IsNotExist(err) {

			if err := os.Mkdir(path, os.ModePerm); err != nil {
				log.Fatal("Package Error", err)
				return false
			}

		} else {
			err := os.RemoveAll(path)
			if err != nil {
				log.Fatal("Package Error", err)
				return false
			}
			if err := os.Mkdir(path, os.ModePerm); err != nil {
				log.Fatal("Package Error", err)
				return false
			}
		}
	}

	return true
}
