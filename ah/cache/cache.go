package cache

import (
	"os"
	"github.com/9seconds/ah/homestorage"
	"github.com/9seconds/ah/shell"
)


type Cache struct {
	homeStorage *homestorage.HomeStorage
	shell *shell.Shell
}


func (c *Cache) Init(homeStorage *homestorage.HomeStorage, shell *shell.Shell) {
	c.homeStorage = homeStorage
	c.shell = shell

	c.shell.Discover()
}


func (c *Cache) GetConfigKey(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	if homeValue, homeOk := c.homeStorage.GetConfigKey(key); homeOk {
		_ = os.Setenv(key, homeValue)
		return homeValue
	}

	result := c.shell.GetEnv(key)
	if result == "" {
		return result
	}

	c.SetConfigKey(key, result)

	return result

	return result
}

func (c *Cache) SetConfigKey(key string, value string) {
	c.homeStorage.SetConfigKey(key, value)
	_ = os.Setenv(key, value)
}
