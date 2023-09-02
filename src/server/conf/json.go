package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cjzswang5/leaf/log"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

func init() {
	execpath, err := os.Executable()
	if err != nil {
		log.Fatal("%v", err)
	}
	execdir := filepath.Dir(execpath)
	fmt.Printf("execdir\t\t= \"%s\"\n", execdir)
	configPath := "bin/conf/server.json"
	fmt.Printf("configPath\t= \"%s\"\n", filepath.Join(execdir, configPath))
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
