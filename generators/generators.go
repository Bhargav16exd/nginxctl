package generators

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var apiConfig = `server {
    listen 80;
    server_name {{. Domain}};

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

func GenerateApiConfig(projectName string, basePath string, port string, domain string) {

	payload := ApiConfig{Port: port, Domain: domain}

	var buf bytes.Buffer

	if err := ApiConfigTemplate.Execute(&buf, payload); err != nil {
		log.Fatal("Error While Generating Template")
	}

	path := filepath.Join(basePath, "api-"+projectName)

	err := os.WriteFile(path, buf.Bytes(), 0644)

	fmt.Println(err)
}

func GenerateStaticContentConfig() {

}
