package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Bhargav16exd/nginxctl/constants"
)

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
			fmt.Println(path, "does not exist")
			fmt.Println("Creating Dir : ", path)

			if err := os.Mkdir(path, os.ModePerm); err != nil {
				log.Fatal("Package Error", err)
			}

		}
	}

	return sitesAvailableDir, sitesEnabledDir
}

func ResetConfigration() {

}
