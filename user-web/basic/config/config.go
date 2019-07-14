package config

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	err error
)

var (
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
	consulConfig            defaultConsulConfig
	profiles                defaultProfiles
	m                       sync.RWMutex
	inited                  bool
)

func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[Init] conf has been init")
		return
	}

	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("./", string(filepath.Separator))))

	pt := filepath.Join(appPath, "conf")
	err := os.Chdir(appPath)
	if err != nil {
		log.Logf("[Init] can not cd the %s", appPath)
	}

	if err = config.Load(file.NewSource(file.WithPath(pt + "/application.yml"))); err != nil {
		panic(err)
	}

	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Logf("[Init] loading config file: path: %s,%+v\n", pt+"application.yml", profiles)

	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")
		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"
			log.Logf("[Init] loading config file: path: %s\n", filePath)
			sources[i] = file.NewSource(file.WithPath(filePath))
		}
		if err:=config.Load(sources...);err!=nil{
			panic(err)
		}
	}
	err = config.Get(defaultRootPath, "consul").Scan(&consulConfig)
	if err!=nil{
		panic(err)
	}

	inited = true
}

func GetConsulConfig() (ret ConsulConfig) {
	return consulConfig
}