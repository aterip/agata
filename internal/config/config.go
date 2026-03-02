package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	IsDebug *bool
	Listen  struct {
		Protocol string `yaml:"protocol" env:"PROTOCOL"`
		BindIP   string `yaml:"host" env:"HOST"`
		Port     string `yaml:"port" env:"PORT"`
	}
}

type DataBasesList struct {
	DataBase map[string]bool `yaml:"dbList" env:"DBLIST"`
}

type PostgresBasicConfig struct {
	Host   string `yaml:"host" env:"HOST_PG"`
	Port   string `yaml:"port" env:"PORT_PG"`
	NameDB string `yaml:"namedb" env:"NAMEDB_PG"`
	User   string `yaml:"usr" env:"USRDB_PG"`
	Pswd   string `yaml:"pswd" env:"PSWD_PG"`
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
