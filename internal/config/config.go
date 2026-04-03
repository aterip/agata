package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	IsDebug *bool
	Listen  struct {
		Type     string `yaml:"type" env-default:"tcp"`
		Protocol string `yaml:"protocol" env-default:"tcp"`
		BindIP   string `yaml:"host" env-default:"127.0.0.1"`
		Port     string `yaml:"port" env-default:"8080"`
	}
}

type DataBasesList struct {
	DataBase map[string]bool `yaml:"dbList" env-default:""`
}

type PostgresBasicConfig struct {
	Host   string `yaml:"host" env-default:"127.0.0.1"`
	Port   string `yaml:"port" env-default:"5432"`
	NameDB string `yaml:"namedb" env-default:"postgres_db"`
	User   string `yaml:"usr" env-default:"postgres"`
	Pswd   string `yaml:"pswd" env-default:"postgres"`
}

var databasesList *DataBasesList
var postgresInstance *PostgresBasicConfig
var instance *ServerConfig
var once sync.Once

func GetServerConfig() *ServerConfig {

	once.Do(func() {
		instance = &ServerConfig{}
		if err := cleanenv.ReadConfig("srvconf.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Fatal(help)
		}
	})
	return instance
}

func GetDataBasesList() *DataBasesList {
	once.Do(func() {
		if err := cleanenv.ReadConfig("srvconf.yml", databasesList); err != nil {
			help, _ := cleanenv.GetDescription(databasesList, nil)
			log.Fatal(help)
		}
	})
	return databasesList
}

func GetPostgresConfig() *PostgresBasicConfig {
	once.Do(func() {
		if err := cleanenv.ReadConfig("postgres.yml", postgresInstance); err != nil {
			help, _ := cleanenv.GetDescription(postgresInstance, nil)
			log.Fatal(help)
		}
	})
	return postgresInstance
}
