package gox_database

import (
	"github.com/harishb2k/gox-base/serialization"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testConfig = `
databases:
  configs:
    master:
      type: mysql
      user: user_1
      password: password_1
      url: [localhost]
      port: 1234
      db: test_1
    scylla:
      type: scylla
      user: user_2
      password: password_2
      url: [localhost]
      port: 12345
      db: test_2
`

type testParsingConfig struct {
	Databases Configs `yaml:"databases"`
}

func TestParsingConfig(t *testing.T) {
	var err error
	configs := testParsingConfig{}
	err = serialization.ReadYamlFromString(testConfig, &configs)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(configs.Databases.Configs))

	assert.NotNil(t, configs.Databases.Configs["master"])
	assert.Equal(t, "mysql", configs.Databases.Configs["master"].Type)
	assert.Equal(t, "user_1", configs.Databases.Configs["master"].User)
	assert.Equal(t, "password_1", configs.Databases.Configs["master"].Password)
	assert.Equal(t, []string{"localhost"}, configs.Databases.Configs["master"].Url)
	assert.Equal(t, 1234, configs.Databases.Configs["master"].Port)
	assert.Equal(t, "test_1", configs.Databases.Configs["master"].Db)

	assert.NotNil(t, configs.Databases.Configs["scylla"])
	assert.Equal(t, "scylla", configs.Databases.Configs["scylla"].Type)
	assert.Equal(t, "user_2", configs.Databases.Configs["scylla"].User)
	assert.Equal(t, "password_2", configs.Databases.Configs["scylla"].Password)
	assert.Equal(t, []string{"localhost"}, configs.Databases.Configs["scylla"].Url)
	assert.Equal(t, 12345, configs.Databases.Configs["scylla"].Port)
	assert.Equal(t, "test_2", configs.Databases.Configs["scylla"].Db)

}
