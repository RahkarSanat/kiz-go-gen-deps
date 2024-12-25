package config

import (
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

var Confs = Config{}

type Config struct {
	Base         string
	Path         string            `yaml:"Path"`
	Port         string            `yaml:"Port"`
	DB           DB                `yaml:"DB"`
	ClickhouseDB ClickhouseDB      `yaml:"ClickhouseDB"`
	HealthPort   string            `yaml:"health_port"`
	Kafka        kafka             `yaml:"kafka"`
	TrxTopics    map[string]string `yaml:"trx_topics"`
	Redis        redis             `yaml:"redis"`
	Cdc          cdc               `yaml:"cdc"`
}
type DB struct {
	Name       string `yaml:"Name"`
	Host       string `yaml:"Host"`
	Port       string `yaml:"Port"`
	User       string `yaml:"User"`
	Password   string `yaml:"Password"`
	DBName     string `yaml:"DBName"`
	Collection string `yaml:"Collection"`
	ConString  string `yaml:"ConString"`
}

type ClickhouseDB struct {
	ConString    string `yaml:"ConString"`
	KafkaGroupId string `yaml:"KafkaGroupId"`
}

type kafka struct {
	Host              string `yaml:"host"`
	GatewayPrefix     string `yaml:"gateway_prefix"`
	GroupId           string `yaml:"group_id"`
	NumPartitions     string `yaml:"num_partitions"`
	ReplicationFactor string `yaml:"replication_factor"`
	RetentionMs       string `yaml:"retention_ms"`
	SchemaRegistry    string `yaml:"schema_registry"`
}

type redis struct {
	Addr string `yaml:"Addr"`
	Db   string `yaml:"Db"`
	Pass string `yaml:"Pass"`
}

type cdc struct {
	KafkaTopicPrefix string `yaml:"KafkaTopicPrefix"`
}

func substituteEnv(data interface{}) {
	confVal := reflect.ValueOf(data)
	if confVal.Kind() == reflect.Ptr {
		confVal = confVal.Elem()
	}

	for i := 0; i < confVal.NumField(); i++ {
		field := confVal.Field(i)

		if field.Kind() == reflect.Struct {
			substituteEnv(field.Addr().Interface())
		} else if field.Kind() == reflect.Map {
			for _, e := range field.MapKeys() {
				v := field.MapIndex(e)
				if strings.HasPrefix(v.String(), "__ENV__") {
					env := os.Getenv(strings.Split(v.String(), "__ENV__")[1])
					newVal := reflect.ValueOf(env)
					field.SetMapIndex(e, newVal)
				}

			}
		} else {
			if strings.HasPrefix(field.String(), "__ENV__") {
				env := os.Getenv(strings.Split(field.String(), "__ENV__")[1])
				newVal := reflect.ValueOf(env)
				field.Set(newVal)
			}
		}
	}
}

func (g *Config) Load(path string) error {
	yfile, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yfile, &Confs)
	if err != nil {
		return err
	}
	substituteEnv(&Confs)

	return nil

}
