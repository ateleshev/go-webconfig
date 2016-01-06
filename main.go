package webconfig

import (
	"fmt"
	"io/ioutil"
	"path"

	"encoding/json"    // JSON (http://www.json.org/)
	"gopkg.in/yaml.v2" // YAML (http://www.yaml.org/)
)

const (
	CONFIG_TYPE_JSON = "json"
	CONFIG_TYPE_YAML = "yaml"

	CONFIG_APP      = "app"
	CONFIG_MAIN     = "main"
	CONFIG_SERVER   = "server"
	CONFIG_DATABASE = "database"
)

var (
	configType = CONFIG_TYPE_JSON
)

func UseYaml() { // {{{
	configType = CONFIG_TYPE_YAML
} // }}}

func UseJson() { // {{{
	configType = CONFIG_TYPE_JSON
} // }}}

// ==[ Filepath ]==

func Filepath(dir string, config ConfigInterface) string { // {{{
	return path.Join(dir, fmt.Sprintf("%s.%s", config.Name(), configType))
} // }}}

// ==[ Load ]==

func Load(dir string, config ConfigInterface) error { // {{{
	err := load(Filepath(dir, config), config)
	config.setLoaded(err == nil)

	return err
} // }}}

func load(file string, config interface{}) error { // {{{
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	switch configType {
	case CONFIG_TYPE_YAML:
		return yaml.Unmarshal(data, config)
	case CONFIG_TYPE_JSON:
		return json.Unmarshal(data, config)
	default:
		return fmt.Errorf("Unknown parser: '%s'", configType)
	}
} // }}}
