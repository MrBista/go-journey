package viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	// berfungsi saat untuk build misal kita ga ingin build lagi dan ingin merubah config aja jadi ga perlu build lagi hanya perlu mengubah confignya saja
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestViperJson(t *testing.T) {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	assert.Nil(t, err)

	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestViperYaml(t *testing.T) {
	config := viper.New()

	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	assert.Nil(t, err)

	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestVipereNV(t *testing.T) {
	config := viper.New()

	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	assert.Nil(t, err)

	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
}
