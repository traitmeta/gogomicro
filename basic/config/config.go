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
	mysqlConfig             defaultMysqlConfig
	redisConfig             defaultRedisConfig
	jwtConfig               defaultJwtConfig
	profiles                defaultProfiles
	m                       sync.RWMutex
	inited                  bool
	sp                      = string(filepath.Separator)
)

func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[InitConfig] 配置已经初始化过")
		return
	}

	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))

	pt := filepath.Join(appPath, "conf")
	_ = os.Chdir(appPath)

	if err = config.Load(file.NewSource(file.WithPath(pt + sp + "application.yml"))); err != nil {
		panic(err)
	}

	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Logf("[Init] 加载配置文件：path: %s, +%v\n", pt+sp+"application.yml", profiles)

	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")
		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"
			log.Logf("[Init] 加载配置文件：path :%s\n", filePath)
			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}
	_ = config.Get(defaultRootPath, "consul").Scan(&consulConfig)
	_ = config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)
	_ = config.Get(defaultRootPath, "redis").Scan(&redisConfig)
	_ = config.Get(defaultRootPath, "redis").Scan(&jwtConfig)
	inited = true
}

func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

func GetConsulConfig() (ret ConsulConfig) {
	return consulConfig
}

func GetReidsConfig() (ret RedisConfig) {
	return redisConfig
}

func GetJwtConfig() (ret JwtConfig) {
	return jwtConfig
}
