package configs

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DbInfoConfig struct {
	postgresUser     string
	postgresPassword string
	postgresHost     string
	postgresPort     string
	postgresDB       string
}

// 起動環境値
var profile string

func init() {
	profile = getEnv("APP_PROFILE", "dev")
}

/*
viperのload
*/
func LoadConfig() {
	viper.SetConfigType("yaml")                   // 設定ファイルの形式
	viper.SetConfigName("config-" + profile)      // 設定ファイル名を拡張子抜きで指定する
	viper.AddConfigPath("./configs/")             // 設定ファイルの探索パスを指定する
	viper.AddConfigPath(".")                      // 現在のワーキングディレクトリを探索することもできる
	readViperConfig()                             // 読み込み
	viper.WatchConfig()                           // 設定フィアルの変更を検知する
	viper.OnConfigChange(func(e fsnotify.Event) { // 検知時に実行する関数の設定
		fmt.Println("設定ファイルが変更されました:", e.Name)
		readViperConfig()
	})
	bindEnvs()    // 環境変数
	setDefaults() // デフォルト設定
}

/*
設定ファイルの読み込み
*/
func readViperConfig() {
	err := viper.ReadInConfig() // 設定ファイルを探索して読み取る
	if err != nil {             // 設定ファイルの読み取りエラー対応
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}
}

/*
viperに環境変数をバインド
*/
func bindEnvs() {
	viper.BindEnv("postgres.host", "POSTGRES_HOST")
	viper.BindEnv("postgres.port", "POSTGRES_PORT")
	viper.BindEnv("postgres.user", "POSTGRES_USER")
	viper.BindEnv("postgres.password", "POSTGRES_PASSWORD")
	viper.BindEnv("postgres.dbname", "POSTGRES_DB")
}

/*
ivperへのデフォルト設定（必要なやつ）
*/
func setDefaults() {
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.user", "wanrun")
	viper.SetDefault("postgres.password", "pass")
	viper.SetDefault("postgres.dbname", "dbname")
}

// 環境変数の取得
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

/*
DB情報のconfig構造体の取得
*/
func DbInfo() *DbInfoConfig {
	config := &DbInfoConfig{
		postgresUser:     viper.GetString("postgres.user"),
		postgresPassword: viper.GetString("postgres.password"),
		postgresHost:     viper.GetString("postgres.host"),
		postgresPort:     viper.GetString("postgres.port"),
		postgresDB:       viper.GetString("postgres.db"),
	}
	return config
}

/*
Getter
*/
func (c *DbInfoConfig) PostgresUser() string {
	return c.postgresUser
}

func (c *DbInfoConfig) PostgresPassword() string {
	return c.postgresPassword
}

func (c *DbInfoConfig) PostgresHost() string {
	return c.postgresHost
}

func (c *DbInfoConfig) PostgresPort() string {
	return c.postgresPort
}

func (c *DbInfoConfig) PostgresDB() string {
	return c.postgresDB
}
