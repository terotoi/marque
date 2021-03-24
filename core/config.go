package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/terotoi/marque/utils"
)

// Config contains configuration options of the application.
type Config struct {
	ListenAddress  string `json:"listen_address"`
	DatabaseType   string `json:"db_type"`
	DatabaseString string `json:"db_string"`
	PasswordSecret string `json:"password_secret"`
	JWTSecret      string `json:"jwt_secret"`
	DataDir        string `json:"data_dir"`
}

// SecretsConfig is an optional configuration for the secrets.
// It is automatically generated if any of the secrets is not
// given in the main config.
type SecretsConfig struct {
	PasswordSecret string `json:"password_secret"`
	JWTSecret      string `json:"jwt_secret"`
}

// LoadConfig loads the configuration file.
func LoadConfig(cfgFile string) (*Config, error) {
	var cfg Config

	d, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(d, &cfg); err != nil {
		return nil, err
	}

	secretsFile := fmt.Sprintf("%s/secrets.json", cfg.DataDir)

	var secrets SecretsConfig
	if cfg.DataDir != "" {
		d, err := ioutil.ReadFile(secretsFile)

		if err != nil && !os.IsNotExist(err) {
			return nil, err
		}

		if err == nil {
			if err = json.Unmarshal(d, &secrets); err != nil {
				return nil, err
			}
		}
	}

	var writeSecrets bool
	if cfg.PasswordSecret == "" {
		if secrets.PasswordSecret != "" {
			cfg.PasswordSecret = secrets.PasswordSecret
		} else {
			log.Println("Generating new password secret.")
			cfg.PasswordSecret, err = utils.GenerateSecret(32)
			if err != nil {
				return nil, err
			}
			secrets.PasswordSecret = cfg.PasswordSecret
			writeSecrets = true
		}
	}

	if cfg.JWTSecret == "" {
		if secrets.JWTSecret != "" {
			cfg.JWTSecret = secrets.JWTSecret
		} else {
			log.Println("Generating new JWT secret.")
			cfg.JWTSecret, err = utils.GenerateSecret(32)
			if err != nil {
				return nil, err
			}
			secrets.JWTSecret = cfg.JWTSecret
			writeSecrets = true
		}
	}

	if writeSecrets {
		if cfg.DataDir == "" {
			return nil, fmt.Errorf("secrets not given and data_dir is empty")
		}

		md, err := json.Marshal(secrets)
		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(secretsFile, md, 0600); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
