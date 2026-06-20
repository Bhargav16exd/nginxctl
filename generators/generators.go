package generators

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

var apiConfig = `server {
    listen 80;
    server_name {{.Domain}};

    location / {
        proxy_pass http://localhost:{{.Port}};
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
`

type ApiConfig struct {
	Port   string
	Domain string
}

var ApiConfigTemplate = template.Must(template.New("nginxApiConfigTemplate").Parse(apiConfig))

func GenerateApiConfig(projectName string, sitesAvailablePath string, sitesEnabledPath string, port string, domain string) bool {

	payload := ApiConfig{Port: port, Domain: domain}

	var buf bytes.Buffer

	if err := ApiConfigTemplate.Execute(&buf, payload); err != nil {
		fmt.Println(err)
		return false
	}

	path := filepath.Join(sitesAvailablePath, "api-"+projectName)

	err := os.WriteFile(path, buf.Bytes(), 0644)

	if err != nil {
		fmt.Println(err)
		return false
	}

	//create symlinkcj
	_, err = exec.Command("ln", "-s", path, sitesEnabledPath).CombinedOutput()

	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = exec.Command("nginx", "-t").CombinedOutput()

	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = exec.Command("nginx", "-s", "reload").CombinedOutput()

	if err != nil {
		fmt.Println(err)
	}

	return true
}

func GenerateStaticContentConfig() {

}
