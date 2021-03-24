package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/terotoi/marque/core"
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

func main() {
	filename := flag.String("file", "", "file to to use [import]")
	cfgFile := flag.String("c", "$HOME/.config/marque/config.json", "configuration file to use")
	flag.Usage = usage
	flag.Parse()

	homeDir := os.Getenv("HOME")
	*cfgFile = strings.Replace(*cfgFile, "$HOME", homeDir, 1)
	*cfgFile = strings.Replace(*cfgFile, "~/", homeDir+"/", 1)
	log.Printf("Loading configuration from %s\n", *cfgFile)

	cfg, err := core.LoadConfig(*cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}
	command := args[0]

	site, db := setup(cfg)
	defer db.Close()

	switch command {
	case "serve":
		err = serve(site, db)

	case "import":
		if *filename == "" {
			flag.Usage()
		} else {
			err = ImportJSON(*filename, db)
		}

	case "export":
		if *filename == "" {
			flag.Usage()
		} else {
			err = ExportJSON(*filename, db)
		}

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
