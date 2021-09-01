package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/terotoi/marque/core"
	"github.com/terotoi/marque/utils"
)

const version = "0.4"

func setup(cfg *core.Config) (*core.Site, *sqlx.DB) {
	rand.Seed(time.Now().Unix())

	db, err := core.SetupDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	var site core.Site
	site.Config = cfg
	site.JWTSecret = []byte(cfg.JWTSecret)
	if err != nil {
		log.Fatal(err)
	}

	return &site, db
}

func usage() {
	fh := flag.CommandLine.Output()
	fmt.Fprintf(fh, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(fh, "  %s [options...] <command>\n", os.Args[0])
	fmt.Fprintf(fh, "  version: %s\n", version)
	fmt.Fprintf(fh, "  options:\n")
	flag.PrintDefaults()
	fmt.Fprintln(fh)
	fmt.Fprintf(fh, "   commands: serve, import, export\n")
	fmt.Fprintln(fh)
	fmt.Fprintf(fh, "   serve  - Serve bookmarks.\n")
	fmt.Fprintf(fh, "   import - Import bookmarks from a custom JSON file.\n")
	fmt.Fprintf(fh, "            example: %s -file filename.json import\n", os.Args[0])
	fmt.Fprintf(fh, "   export - Export bookmarks to a custom JSON file.\n")
	fmt.Fprintf(fh, "            example: %s -file filename.json export\n", os.Args[0])
}

func startServe(cfgFile string) error {
	log.Printf("Loading configuration from %s", cfgFile)

	cfg, err := core.LoadConfig(cfgFile)
	if err != nil {
		return err
	}

	site, db := setup(cfg)
	defer db.Close()

	return serve(site, db)
}

func main() {
	cfgFile := flag.String("c", "$HOME/.config/marque/config.json", "configuration file to use")
	dataDir := flag.String("d", "$HOME/.config/marque", "data directory")
	listenAddress := flag.String("l", ":9999", "listen to this address")
	createInitialUser := flag.Bool("i", false, "create an initial admin user")

	flag.Usage = usage
	flag.Parse()

	*cfgFile = utils.ReplaceEnvs(*cfgFile)
	*dataDir = utils.ReplaceEnvs(*dataDir)

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}
	command := args[0]

	var err error
	switch command {
	case "createconfig":
		_, err = core.GenerateConfig(*cfgFile, *dataDir, *listenAddress, *createInitialUser)

	case "serve":
		err = startServe(*cfgFile)

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
