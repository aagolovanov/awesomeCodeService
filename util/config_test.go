package util

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type test struct {
		Name     string
		Env      map[string]string
		Expected Config
	}

	tests := []test{
		{
			"New env config",
			map[string]string{
				"PORT":    "1234",
				"DB_ADDR": "hello",
				"DB_PORT": "123",
				"DB_PASS": "password",
				"TTL":     "228",
			},
			Config{
				Port:   1234,
				DBAddr: "hello:123",
				DBPass: "password",
				TTL:    228,
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {
			for k, v := range tt.Env {
				t.Setenv(k, v)
			}

			config := LoadConfig()

			if !reflect.DeepEqual(config, tt.Expected) {
				t.Errorf("Expected: %v, got: %v\n", tt.Expected, config)
			}
		})

	}
}
