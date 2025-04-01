package config

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type LogLevel string

const (
	Info  LogLevel = "info"
	Debug LogLevel = "debug"
	Trace LogLevel = "trace"
	None  LogLevel = "none"
)

type Password string

func (p Password) MarshalText() ([]byte, error) {
	return []byte("*************"), nil
}

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	HTTP     HTTPConfig     `mapstructure:"http"`
	Postgres DatabaseConfig `mapstructure:"postgres"`
	Logging  LogConfig      `mapstructure:"logging"`
}

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Environment string `mapstructure:"environment"`
	Version     string `mapstructure:"version"`
}

type HTTPConfig struct {
	Host           string   `mapstructure:"host"`
	Port           int      `mapstructure:"port"`
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	Timeout        int      `mapstructure:"timeout"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"db"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"ssl_mode"`
	PoolMax  int    `mapstructure:"pool_max"`
}

type LogConfig struct {
	Level  LogLevel `mapstructure:"level"`
	Format string   `mapstructure:"format"`
}

func LoadConfig[E any]() (*E, error) {
	var err error
	var conf *E

	conf, err = readConfig[E](".env")
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func readConfig[E any](configFilePath string) (*E, error) {
	vp := viper.New()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vp.AutomaticEnv()
	vp.SetConfigFile(configFilePath)

	var config E
	err := Unmarshal(vp, &config, ".")
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %w", err)
	}

	return &config, nil
}

func Unmarshal(v *viper.Viper, rawVal any, keyDelim string, opts ...viper.DecoderConfigOption) error {
	keys := v.AllKeys()

	structKeys, err := decodeStructKeys(rawVal, keyDelim, opts...)
	if err != nil {
		return err
	}

	keys = append(keys, structKeys...)

	return decode(getSettings(v, keys, keyDelim), defaultDecoderConfig(rawVal, opts...))
}

func defaultDecoderConfig(output any, opts ...viper.DecoderConfigOption) *mapstructure.DecoderConfig {
	c := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func decode(input any, config *mapstructure.DecoderConfig) error {
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

func getSettings(v *viper.Viper, keys []string, keyDelim string) map[string]any {
	m := map[string]any{}
	for _, k := range keys {
		value := v.Get(k)
		if value == nil {
			continue
		}
		path := strings.Split(k, keyDelim)
		lastKey := strings.ToLower(path[len(path)-1])
		deepestMap := deepSearch(m, path[0:len(path)-1])
		deepestMap[lastKey] = value
	}
	return m
}

func deepSearch(m map[string]any, path []string) map[string]any {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
			m3 := make(map[string]any)
			m[k] = m3
			m = m3
			continue
		}
		m3, ok := m2.(map[string]any)
		if !ok {
			m3 = make(map[string]any)
			m[k] = m3
		}
		m = m3
	}
	return m
}

func decodeStructKeys(input any, keyDelim string, opts ...viper.DecoderConfigOption) ([]string, error) {
	var structKeyMap map[string]any

	err := decode(input, defaultDecoderConfig(&structKeyMap, opts...))
	if err != nil {
		return nil, err
	}

	flattenedStructKeyMap := flattenAndMergeMap(map[string]bool{}, structKeyMap, "", keyDelim)

	r := make([]string, 0, len(flattenedStructKeyMap))
	for v := range flattenedStructKeyMap {
		r = append(r, v)
	}

	return r, nil
}

func flattenAndMergeMap(shadow map[string]bool, m map[string]any, prefix string, keyDelim string) map[string]bool {
	if shadow != nil && prefix != "" && shadow[prefix] {
		return shadow
	}
	if shadow == nil {
		shadow = make(map[string]bool)
	}

	var m2 map[string]any
	if prefix != "" {
		prefix += keyDelim
	}
	for k, val := range m {
		fullKey := prefix + k
		switch val := val.(type) {
		case map[string]any:
			m2 = val
		case map[any]any:
			m2 = cast.ToStringMap(val)
		default:
			shadow[strings.ToLower(fullKey)] = true
			continue
		}
		shadow = flattenAndMergeMap(shadow, m2, fullKey, keyDelim)
	}
	return shadow
}
