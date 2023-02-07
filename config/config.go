package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	redisClient "go-template/client/redis"
	"go-template/logger"

	redis "github.com/go-redis/redis/v8"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"

	"github.com/newrelic/go-agent/_integrations/nrlogrus"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	AppName        = "GO_TEMPLATE"
	consulEndpoint = "127.0.0.1:8500"
	consulPath     = "GO_TEMPLATE"
	Env            string
	MasterDB       *sqlx.DB
	SlaveDB        *sqlx.DB
)

type Config struct {
	Env                string         `mapstructure:"env"`
	Port               int            `mapstructure:"port"`
	LogLevel           string         `mapstructure:"logLevel"`
	LogFormat          string         `mapstructure:"logFormat"`
	AppTimeout         int            `mapstructure:"appTimeout"`
	AppCorsDomain      string         `mapstructure:"appCorsDomain"`
	APITokenKey        string         `mapstructure:"apiTokenKey"`
	Postgres           PostgresConfig `mapstructure:"postgres"`
	Redis              Redis          `mapstructure:"redis"`
	Swagger            Swagger        `mapstructure:"swagger"`
	NewRelic           NewRelic       `mapstructure:"newRelic"`
	NewRelicLicenseKey string         `mapstructure:"newRelicLicenseKey"`
}

type ApplicationConfig struct {
}

type PostgresConfig struct {
	ConnMaxLifetime    int  `mapstructure:"connMaxLifetime"`
	MaxOpenConnections int  `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int  `mapstructure:"maxIdleConnections"`
	MaxIdleLifetime    int  `mapstructure:"maxIdleLifetime"`
	ConnectTimeout     int  `mapstructure:"connectTimeout"`
	Master             PSQL `mapstructure:"master"`
	Slave              PSQL `mapstructure:"slave"`
}

type PSQL struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
	DBName   string `mapstructure:"dbName"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
}

type Swagger struct {
	Title   string `mapstructure:"title"`
	Version string `mapstructure:"version"`
	Url     string `mapstructure:"url"`
	Schemes string `mapstructure:"schemes"`
}

type NewRelic struct {
	ApplicationName string `mapstructure:"applicationName"`
	IsActive        bool   `mapstructure:"isActive"`
}

func New() (conf Config) {
	var once sync.Once
	once.Do(func() {
		v := viper.New()
		retried := 0
		err := InitialiseRemote(v, retried)
		if err != nil {
			log.Printf("No remote server configured will load configuration from file and environment variables: %+v", err)
			if err := InitialiseFileAndEnv(v, "config.local"); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					configFileName := fmt.Sprintf("%s.yaml", "config.local")
					log.Printf("No '" + configFileName + "' file found on search paths. Will either use environment variables or defaults")
				} else {
					log.Fatalf("Error occured during loading config: %s", err.Error())
				}
			}
		}
		err = v.Unmarshal(&conf)
		if err != nil {
			log.Fatalf("%v", err)
		}
	})
	return conf
}

func InitialiseRemote(v *viper.Viper, retried int) error {
	if consulEnv := os.Getenv("CONSUL_URL"); consulEnv != "" {
		consulEndpoint = consulEnv
	}
	log.Printf("Initialising remote config, consul endpoint: %s, consul path: %s, retried: %d", consulEndpoint, consulPath, retried)
	v.AddRemoteProvider("consul", consulEndpoint, consulPath)
	v.SetConfigType("yaml")
	err := v.ReadRemoteConfig()
	if err != nil && retried < 1 {
		time.Sleep(500 * time.Millisecond)
		return InitialiseRemote(v, retried+1)
	}
	return err
}

func SetupMasterDB(conf Config, nrapp *newrelic.Application, unitTest bool) (db *sqlx.DB) {
	log.Printf("Connecting to postgresql (master) %s", fmt.Sprintf("host=%s user=%s dbname=%s search_path=%s sslmode=%s",
		conf.Postgres.Master.Host, conf.Postgres.Master.User, conf.Postgres.Master.DBName, conf.Postgres.Master.Schema, "disable"))

	var (
		masterDB *sqlx.DB
		err      error
	)

	if nrapp != nil {
		masterDB, err = sqlx.Open("nrpostgres",
			fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=%s connect_timeout=%d",
				conf.Postgres.Master.Host, conf.Postgres.Master.Port, conf.Postgres.Master.User, conf.Postgres.Master.Password, conf.Postgres.Master.DBName, conf.Postgres.Master.Schema, "disable", conf.Postgres.ConnectTimeout))
	} else {
		masterDB, err = sqlx.Open("postgres",
			fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=%s connect_timeout=%d",
				conf.Postgres.Master.Host, conf.Postgres.Master.Port, conf.Postgres.Master.User, conf.Postgres.Master.Password, conf.Postgres.Master.DBName, conf.Postgres.Master.Schema, "disable", conf.Postgres.ConnectTimeout))
	}
	if err != nil {
		log.Fatalf("open db connection failed (master) %v", err)
	}
	masterDB.SetMaxOpenConns(conf.Postgres.MaxOpenConnections)
	masterDB.SetMaxIdleConns(conf.Postgres.MaxIdleConnections)
	masterDB.SetConnMaxIdleTime(time.Duration(conf.Postgres.MaxIdleLifetime) * time.Millisecond)
	if err := masterDB.Ping(); err != nil && !unitTest {
		log.Fatalf("ping db connection failed (master) %v", err)
	}

	return masterDB
}

func SetupSlaveDB(conf Config, nrapp *newrelic.Application, unitTest bool) (db *sqlx.DB) {
	log.Printf("Connecting to postgresql (slave) %s", fmt.Sprintf("host=%s user=%s dbname=%s search_path=%s sslmode=%s",
		conf.Postgres.Slave.Host, conf.Postgres.Slave.User, conf.Postgres.Slave.DBName, conf.Postgres.Slave.Schema, "disable"))

	var (
		slaveDB *sqlx.DB
		err     error
	)

	if nrapp != nil {
		slaveDB, err = sqlx.Open("nrpostgres",
			fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=%s connect_timeout=%d",
				conf.Postgres.Slave.Host, conf.Postgres.Slave.Port, conf.Postgres.Slave.User, conf.Postgres.Slave.Password, conf.Postgres.Slave.DBName, conf.Postgres.Slave.Schema, "disable", conf.Postgres.ConnectTimeout))
	} else {
		slaveDB, err = sqlx.Open("postgres",
			fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=%s connect_timeout=%d",
				conf.Postgres.Slave.Host, conf.Postgres.Slave.Port, conf.Postgres.Slave.User, conf.Postgres.Slave.Password, conf.Postgres.Slave.DBName, conf.Postgres.Slave.Schema, "disable", conf.Postgres.ConnectTimeout))
	}
	if err != nil {
		log.Fatalf("open db connection failed (slave) %v", err)
	}
	slaveDB.SetMaxOpenConns(conf.Postgres.MaxOpenConnections)
	slaveDB.SetMaxIdleConns(conf.Postgres.MaxIdleConnections)
	slaveDB.SetConnMaxIdleTime(time.Duration(conf.Postgres.MaxIdleLifetime) * time.Millisecond)
	if err := slaveDB.Ping(); err != nil && !unitTest {
		log.Fatalf("ping db connection failed (slave) %v", err)
	}

	return slaveDB
}

func InitialiseFileAndEnv(v *viper.Viper, configName string) error {
	var searchPath = []string{
		"/etc/go-template",
		"$HOME/.go-template",
		".",
	}
	v.SetConfigName(configName)
	for _, path := range searchPath {
		v.AddConfigPath(path)
	}
	v.SetEnvPrefix("go-template")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	return v.ReadInConfig()
}

func SetupRedis(conf Config) *redisClient.Client {
	redis := redisClient.NewClient(redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Password,
		DB:       0,
	}))
	return redis
}

func SetupNewRelic(config Config) *newrelic.Application {
	if strings.ToLower(config.Env) == "local" || !config.NewRelic.IsActive {
		logger.Warn(context.Background(), "SetupNewRelic - NewRelicLicenseKey is optional on local, skipping...")
		return nil
	}

	//initiate new relic apps
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.NewRelic.ApplicationName),
		newrelic.ConfigLicense(config.NewRelicLicenseKey),
		func(cfg *newrelic.Config) {
			cfg.DistributedTracer.Enabled = true
			cfg.Logger = nrlogrus.Transform(logger.GetLogger())
		},
	)
	if err != nil {
		logger.Warn(context.Background(), "SetupNewRelic.NewApplication - %v", err)
	}

	// Wait for the application to connect.
	if err = app.WaitForConnection(5 * time.Second); nil != err {
		logger.Warn(context.Background(), "SetupNewRelic.WaitForConnection - %v", err)
	}

	//end initiation
	return app
}
