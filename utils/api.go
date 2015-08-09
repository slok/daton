package utils

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/slok/daton/configuration"
)

func ApiUrl(relativeUrl string) string {
	host := viper.GetString("AppHost")
	prefix := fmt.Sprintf("api/v%d", configuration.ApiVersion)

	return fmt.Sprintf("%s/%s/%s", host, prefix, relativeUrl)
}
