package core

import (
	"encoding/base64"
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
}

func setDefaults(cfg *Config) {
	cfg.ListenAddress = "127.0.0.1:9999"
	cfg.DatabaseType = "sqlite3"

	homeDir := os.Getenv("HOME")
	dbLoc := fmt.Sprintf("%s/.config/marque/marque.db", homeDir)
	cfg.DatabaseString = fmt.Sprintf("file:%s?cache=shared&mode=rwc", dbLoc)
}

// LoadConfig loads the configuration file.
func LoadConfig() (*Config, error) {
	homeDir := os.Getenv("HOME")

	cfgDir := fmt.Sprintf("%s/.config/marque", homeDir)

	var cfg Config

	if _, err := os.Stat(cfgDir); os.IsNotExist(err) {
		if err = os.Mkdir(cfgDir, 0700); err != nil {
			log.Printf("Failed to create directory %s", cfgDir)
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	cfgFile := fmt.Sprintf("%s/.config/marque/config.json", homeDir)
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		setDefaults(&cfg)
	} else if err != nil {
		return nil, err
	} else {
		contents, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(contents, &cfg); err != nil {
			return nil, err
		}
	}

	var rewrite bool
	if cfg.PasswordSecret == "" {
		secret, err := utils.GenerateSecret(512)
		if err != nil {
			return nil, err
		}

		cfg.PasswordSecret = base64.RawStdEncoding.EncodeToString(secret)
		rewrite = true
	}

	if cfg.JWTSecret == "" {
		secret, err := utils.GenerateSecret(512)
		if err != nil {
			return nil, err
		}
		cfg.JWTSecret = base64.RawStdEncoding.EncodeToString(secret)
		rewrite = true
	}

	if rewrite {
		data, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(cfgFile, data, 0640); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
