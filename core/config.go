package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/terotoi/marque/utils"
)

// Config contains configuration options of the application.
type Config struct {
	ListenAddress   string `json:"listen_address"`
	DatabaseType    string `json:"db_type"`
	DatabaseConfig  string `json:"db_string"`
	PasswordSecret  string `json:"password_secret"`
	JWTSecret       string `json:"jwt_secret"`
	InitialUser     string `json:"initial_user"`
	InitialPassword string `json:"initial_password"`
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

	if cfg.PasswordSecret == "" {
		log.Println("Generating new password secret.")
		cfg.PasswordSecret, err = utils.GenerateSecret(32)
		if err != nil {
			return nil, err
		}
	}

	if cfg.JWTSecret == "" {
		log.Println("Generating new JWT secret.")
		cfg.JWTSecret, err = utils.GenerateSecret(32)
		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

func GenerateConfig(cfgFile, dataDir, listenAddress string, createInitialUser bool) (*Config, error) {
	if _, err := os.Stat(cfgFile); err == nil {
		return nil, fmt.Errorf("configuration file %s already exists", cfgFile)
	}

	if err := os.MkdirAll(path.Dir(cfgFile), 0700); err != nil && !os.IsExist(err) {
		return nil, err
	}

	pwSecret, err := utils.GenerateSecret(32)
	if err != nil {
		return nil, err
	}

	jwtSecret, err := utils.GenerateSecret(32)
	if err != nil {
		return nil, err
	}

	cfg := Config{
		ListenAddress:  listenAddress,
		DatabaseType:   "sqlite3",
		DatabaseConfig: fmt.Sprintf("file:%s/marque.db?cache=shared\u0026mode=rwc", dataDir),
		PasswordSecret: pwSecret,
		JWTSecret:      jwtSecret,
	}

	if createInitialUser {
		cfg.InitialUser = "admin"
		cfg.InitialPassword = "admin"
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return nil, err
	}

	log.Printf("Writing a configuration file to %s", cfgFile)

	if err := ioutil.WriteFile(cfgFile, data, 0600); err != nil {
		return nil, err
	}

	return &cfg, nil
}
