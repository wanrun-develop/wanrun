package common

import (
	"github.com/wanrun-develop/wanrun/configs"
)

const (
	ENV           = "env"
	STAGE         = "stage"
	ENV_DEV       = "dev"
	ENV_STAGING   = "staging"
	ENV_PROD      = "prod"
	STAGING_LOCAL = "local"
	STAGING_CLOUD = "cloud"
)

// IsProduction: 本番環境かどうかを確認する関数
// secureな状態が求められる場合等に使用
// args:
//   - None
//
// return:
//   - bool: 本番環境かどうか
func IsProduction() bool {
	return configs.FetchConfigStr(ENV) == ENV_PROD
}

// IsStaging: ステージング環境かどうかを確認する関数
// secureな状態が求められない場合等に使用
// args:
//   - None
//
// return:
//   - bool: ステージング環境かどうか
func IsStaging() bool {
	return configs.FetchConfigStr(ENV) == ENV_STAGING
}

// IsDevelopment: 開発環境かどうかを確認する関数
// secureな状態が求められない場合等に使用
// args:
//   - None
//
// return:
//   - bool: 開発環境かどうか
func IsDevelopment() bool {
	return configs.FetchConfigStr(ENV) == ENV_DEV
}

// IsLocal: ローカル環境かどうかを確認する関数
//
// args:
//   - None
//
// return:
//   - bool: ローカル環境かどうか
func IsLocal() bool {
	return configs.FetchConfigStr(STAGE) == STAGING_LOCAL
}

// IsCloud: クラウド環境かどうかを確認する関数
//
// args:
//   - None
//
// return:
//   - bool: クラウド環境かどうか
func IsCloud() bool {
	return configs.FetchConfigStr(STAGE) == STAGING_CLOUD
}
