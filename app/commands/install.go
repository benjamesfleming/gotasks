package commands

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gookit/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// Random Hex
// generate n hex bytes
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Extract File
// extract any file from a given box to a given location, and templated with given data
func extractFile(box *rice.Box, fn, dest string, data interface{}) {
	log.Info().Msgf("Parsing Template [%s]", fn)
	t, _ := template.New("file").Parse(
		box.MustString(fn),
	)

	log.Info().Msgf("Executing Template [%s]", fn)
	tpl := new(bytes.Buffer)
	if err := t.Execute(tpl, data); err != nil {
		log.Error().Msgf("Failed To Execute Template [%s]", fn)
		log.Fatal().Msg("Aborting")
		panic(err)
	}

	log.Info().Msgf("Writing Template [%s] File To Disk", fn)
	if err := ioutil.WriteFile(dest+"/"+fn, tpl.Bytes(), 0644); err != nil {
		log.Error().Msgf("Failed To Save Template [%s] To Disk", fn)
		log.Fatal().Msg("Aborting")
		panic(err)
	}

	log.Info().Msgf("  -> %s", dest+"/"+fn)
}

// InstallCommand ...
var InstallCommand = &cli.Command{
	HelpName: "gotasks install",
	Name:     "install",
	Usage:    "add supporting files",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "with-sqlite3",
			Usage: "Generate a blank sqlite3 database",
		},
		&cli.StringFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Usage:    "Save configuration template to `DIR`",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

		fmt.Println("----")

		// Checking Given Arguments
		// check that the users arguments are valid
		dest := ctx.String("output")
		if stat, err := os.Stat(dest); err != nil || !stat.IsDir() {
			if os.IsNotExist(err) || !stat.IsDir() {
				log.Error().Msgf("Failed To Access Output [%s], Ensure Directory Exists", dest)
			} else {
				log.Error().Err(err).Msg("Failed To Access Output Directory")
			}
			log.Fatal().Msg("Aborting")
			panic(err)
		}

		// Generating Unique Key
		// this gerneates the unique key to load into the template
		log.Info().Msg("Installing Supporting Files")
		log.Info().Msg("Generating Unique Key")
		key, err := randomHex(32)
		if err != nil {
			log.Error().Msg("Failed To Genereate A Unique Key")
			log.Fatal().Msg("Aborting")
			panic(err)
		}
		log.Info().Msgf("  -> %s", key)

		// Load Templates Box
		// this finds the config template on either the file system or the embedded resources
		log.Info().Msg("Loading Templates Box")
		box, err := rice.FindBox("../../templates")
		if err != nil {
			log.Error().Msg("Failed To Load The Templates Box")
			log.Fatal().Msg("Aborting")
			panic(err)
		}

		// Extract Files
		// extract all the files to their correct install locations
		extractFile(box, "gotasks.sample.toml", dest, key)

		// Create Database File
		// create a blank sqlite3 database
		if ctx.Bool("with-sqlite3") {
			log.Info().Msg("Creating a Blank Database File")
			if _, err := os.Create(dest + "/database.db"); err != nil {
				log.Error().Msg("Failed To Create The Database")
				log.Fatal().Msg("Aborting")
				panic(err)
			}
			log.Info().Msgf("  -> %s", dest+"/database.db")
		} else {
			log.Warn().Msg("Skipping Database Creation, Use --with-sqlite3 To Generate Database")
		}

		log.Info().Msg("Done")
		fmt.Println("----")
		fmt.Printf(
			"NEXT STEPS:\n   %v\n   %v\n   %v\n   %v\n",
			color.Sprintf("1. <info>cd</> <warn>%v</>", dest),
			color.Sprintf("2. <info>cp</> <warn>%v %v</>", "gotasks.sample.toml", "gotasks.toml"),
			color.Sprintf("3. follow setup steps in <warn>%v</>", "gotasks.toml"),
			color.Sprintf("4. <info>gotasks start --help</>"),
		)
		return nil
	},
}
