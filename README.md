###Example 
##### Config file with all database properties
A config file which as all DB defined. This config has 2 database in it; first is mysql
, and second is scylla
```yaml
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
```

##### Code 
```go
type mainConfig struct {
	Databases Configs `yaml:"databases"`
}

var err error
configs := testParsingConfig{}
err = gox.ReadYaml("./config.yaml", &configs)
assert.NoError(t, err)
assert.Equal(t, 2, len(configs.Databases.Configs))

```