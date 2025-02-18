package config

import "fmt"

type (
	DatabaseConfig struct {
		MyDbKub BaseDatabaseConfig `mapstructure:"my_db_kub"`
		EieiDb  BaseDatabaseConfig `mapstructure:"eiei_db"`
	}

	BaseDatabaseConfig struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	}
)

type (
	RedisConfig struct {
		MyRdb      BaseRedisConfig[MyRdbExpirationConfig] `mapstructure:"my_rdb"`
		MyOtherRdb BaseRedisConfig[MyRdbExpirationConfig] `mapstructure:"my_other_rdb"`
	}

	BaseRedisConfig[T any] struct {
		Host             string                `mapstructure:"host"`
		Password         string                `mapstructure:"password"`
		Port             string                `mapstructure:"port"`
		Db               int                   `mapstructure:"db"`
		Connection       RedisConnectionConfig `mapstructure:"connection"`
		ExpirationConfig T                     `mapstructure:"expiration_config"`
	}

	RedisConnectionConfig struct {
		Default            bool   `mapstructure:"default"`
		MaxRetries         int    `mapstructure:"max_retries"`
		MinRetryBackoff    string `mapstructure:"min_retry_backoff"`
		MaxRetryBackoff    string `mapstructure:"max_retry_backoff"`
		DialTimeout        string `mapstructure:"dial_timeout"`
		ReadTimeout        string `mapstructure:"read_timeout"`
		WriteTimeout       string `mapstructure:"write_timeout"`
		PoolSize           int    `mapstructure:"pool_size"`
		MinIdleConns       int    `mapstructure:"min_idle_conns"`
		MaxConnAge         string `mapstructure:"max_conn_age"`
		PoolTimeout        string `mapstructure:"pool_timeout"`
		IdleTimeout        string `mapstructure:"idle_timeout"`
		IdleCheckFrequency string `mapstructure:"idle_check_frequency"`
	}

	MyRdbExpirationConfig struct {
		InqUser int `mapstructure:"inq_user"`
	}
)

func (m BaseDatabaseConfig) NewMysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
	)
}
