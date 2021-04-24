package gox_database

import (
	"github.com/devlibx/gox-base"
	"github.com/pkg/errors"
)

// Function to setup a driver to handle a DB
type DatabasePlugin func(config *Config, function gox.CrossFunction) (Db, error)

// Storage to keep all database plugins
var databasePlugins = map[string]DatabasePlugin{}

// Storage to keep all database
var databases = map[string]Db{}

// Register database to plugin store
func RegisterDatabasePlugin(dbType string, plugin DatabasePlugin) {
	databasePlugins[dbType] = plugin
}

// Get the registered database
func GetOrCreate(databaseName string, config *Config, function gox.CrossFunction) (Db, error) {
	if db, ok := databases[databaseName]; !ok {
		if dbPlugin, ok := databasePlugins[config.Type]; !ok {
			return nil, errors.New("failed to create db using plugin - make sure you imported '_ github.com/devlibx/gox-mysql'")
		} else {
			if db, err := dbPlugin(config, function); err != nil {
				return nil, errors.Wrap(err, "failed to create db using plugin")
			} else {
				databases[databaseName] = db
				return db, nil
			}
		}
	} else {
		return db, nil
	}
}
