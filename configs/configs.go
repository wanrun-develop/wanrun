package configs

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DBInfoConfig struct {
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
func LoadConfig() error {
	viper.SetConfigType("yaml")                  // 設定ファイルの形式
	viper.SetConfigName("config-" + profile)     // 設定ファイル名を拡張子抜きで指定する
	viper.AddConfigPath("./configs/")            // 設定ファイルの探索パスを指定する
	viper.AddConfigPath(".")                     // 現在のワーキングディレクトリを探索することもできる
	if err := viper.ReadInConfig(); err != nil { // 設定ファイルを探索して読み取る
		_ = fmt.Errorf("設定ファイル読み込みエラー: %s \n", err) //ここでは停止させない
		return err
	}
	viper.WatchConfig()                           // 設定フィアルの変更を検知する
	viper.OnConfigChange(func(e fsnotify.Event) { // 検知時に実行する関数の設定
		fmt.Println("設定ファイルが変更されました:", e.Name)
		if err := viper.ReadInConfig(); err != nil { // 設定ファイルを探索して読み取る
			_ = fmt.Errorf("設定ファイル読み込みエラー: %s \n", err) //ここでは停止させない
		}
	})
	bindEnvs()    // 環境変数
	setDefaults() // デフォルト設定

	return nil
}

/*
viperに環境変数をバインド
*/
func bindEnvs() {
	_ = viper.BindEnv("postgres.host", "POSTGRES_HOST")
	_ = viper.BindEnv("postgres.port", "POSTGRES_PORT")
	_ = viper.BindEnv("postgres.user", "POSTGRES_USER")
	_ = viper.BindEnv("postgres.password", "POSTGRES_PASSWORD")
	_ = viper.BindEnv("postgres.dbname", "POSTGRES_DB")
}

/*
ivperへのデフォルト設定（必要なやつ）
*/
func setDefaults() {
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.user", "wanrun")
	viper.SetDefault("postgres.password", "__dummdy__")
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
func DbInfo() *DBInfoConfig {
	config := &DBInfoConfig{
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
func (c *DBInfoConfig) PostgresUser() string {
	return c.postgresUser
}

func (c *DBInfoConfig) PostgresPassword() string {
	return c.postgresPassword
}

func (c *DBInfoConfig) PostgresHost() string {
	return c.postgresHost
}

func (c *DBInfoConfig) PostgresPort() string {
	return c.postgresPort
}

func (c *DBInfoConfig) PostgresDB() string {
	return c.postgresDB
}

func FetchCondigStr(key string) string {
	return viper.GetString(key)
}
