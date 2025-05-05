package configs

import (
	"log"

	"github.com/spf13/viper"
)

type DBInfoConfig struct {
	postgresUser     string
	postgresPassword string
	postgresHost     string
	postgresPort     string
	postgresDB       string
}

func init() {
	if err := LoadConfig(); err != nil {
		log.Fatalf("設定の初期化に失敗しました: %s \n", err)
	}
}

var v *viper.Viper

/*
viperのload (環境変数のみ)
*/
func LoadConfig() error {
	v = viper.New()

	// 環境変数のバインドとデフォルト値の設定のみを行う
	bindEnvs()    // 環境変数
	setDefaults() // デフォルト設定

	return nil
}

/*
viperに環境変数をバインド
*/
func bindEnvs() {
	_ = v.BindEnv("postgres.host", "POSTGRES_HOST")
	_ = v.BindEnv("postgres.port", "POSTGRES_PORT")
	_ = v.BindEnv("postgres.user", "POSTGRES_USER")
	_ = v.BindEnv("postgres.password", "POSTGRES_PASSWORD")
	_ = v.BindEnv("postgres.dbname", "POSTGRES_DB")
	_ = v.BindEnv("stage", "STAGE") // アプリケーションの起動環境
	_ = v.BindEnv("env", "ENV")     // 環境情報（dev, staging, prod）
	_ = v.BindEnv("google.place.api.key", "GOOGLE_PLACE_API_KEY")
	_ = v.BindEnv("jwt.os.secret.key", "SECRET_KEY")                // jwt生成用の秘密鍵
	_ = v.BindEnv("jwt.exp.time", "JWT_EXP_TIME")                   // jwt生成用の有効期限（時間）
	_ = v.BindEnv("refresh.jwt.exp.time", "REFRESH_JWT_EXP_TIME")   // jwt生成用の有効期限（時間）
	_ = v.BindEnv("gcp.client.id", "GCP_CLIENT_ID")                 // oauthの際のgcp credentials
	_ = v.BindEnv("gcp.client.secret", "GCP_CLIENT_SECRET")         // oauthの際のgcp credentials
	_ = v.BindEnv("gcp.redirect.uri", "GCP_REDIRECT_URI")           // oauthの際のgcp credentials
	_ = v.BindEnv("aws.access.key", "AWS_ACCESS_KEY")               // awsのアクセスキー
	_ = v.BindEnv("aws.secret.access.key", "AWS_SECRET_ACCESS_KEY") // awsのシークレットアクセスキー
	_ = v.BindEnv("aws.s3.bucket.name", "AWS_S3_BUCKET_NAME")       // awsのbucket名
	_ = v.BindEnv("log.level", "LOG_LEVEL")                         // ログレベル
}

/*
viperへのデフォルト設定（必要なやつ）
*/
func setDefaults() {
	v.SetDefault("postgres.host", "localhost")
	v.SetDefault("postgres.port", "5432")
	v.SetDefault("postgres.user", "wanrun")
	v.SetDefault("postgres.password", "__dummdy__")
	v.SetDefault("postgres.dbname", "dbname")
	v.SetDefault("log.level", "info") // ログレベルのデフォルト値
	v.SetDefault("env", "dev")        // 環境情報のデフォルト値
	v.SetDefault("stage", "local")    // アプリケーションの起動環境のデフォルト値
}

/*
DB情報のconfig構造体の取得
*/
func DbInfo() *DBInfoConfig {
	return &DBInfoConfig{
		postgresUser:     v.GetString("postgres.user"),
		postgresPassword: v.GetString("postgres.password"),
		postgresHost:     v.GetString("postgres.host"),
		postgresPort:     v.GetString("postgres.port"),
		postgresDB:       v.GetString("postgres.dbname"),
	}
}
func (c DBInfoConfig) PostgresUser() string {
	return c.postgresUser
}
func (c DBInfoConfig) PostgresPassword() string {
	return c.postgresPassword
}
func (c DBInfoConfig) PostgresHost() string {
	return c.postgresHost
}
func (c DBInfoConfig) PostgresPort() string {
	return c.postgresPort
}
func (c DBInfoConfig) PostgresDB() string {
	return c.postgresDB
}

// db情報 end

/*
loadしたviperからkeyで値を取得
*/
func FetchConfigStr(key string) string {
	return v.GetString(key)
}

/*
loadしたviperからkeyで値を取得(int)
*/
func FetchConfigInt(key string) int {
	return v.GetInt(key)
}

/*
loadしたviperからkeyで値を取得(bool)
*/
func FetchConfigBool(key string) bool {
	return v.GetBool(key)
}
