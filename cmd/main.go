package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/olenindenis/vault-extractor/internal/commands"
	"github.com/olenindenis/vault-extractor/pkg/envs"
	"github.com/olenindenis/vault-extractor/pkg/vault"

	"github.com/urfave/cli/v2"
)

func main() {
	var (
		envName  string
		fileName string
	)

	app := &cli.App{
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "conf",
						Usage:       "env file name used to connect to vault, if empty then used os envs (-conf=.env)",
						Destination: &envName,
					},
					&cli.StringFlag{
						Name:        "file",
						Value:       ".env",
						Usage:       ".env custom file name",
						Destination: &fileName,
					},
				},
				Name:    "env",
				Aliases: []string{"e"},
				Usage:   "extract envs from vault end save as .env file (-file=.env.dev)",
				Action: func(cCtx *cli.Context) error {
					if err := run(commands.CmdModeEnv, envName, fileName); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "conf",
						Usage:       "env file name used to connect to vault, if empty then used os envs",
						Destination: &envName,
					},
					&cli.StringFlag{
						Name:        "file",
						Value:       "config.json",
						Usage:       "config.json custom file name",
						Destination: &fileName,
					},
				},
				Name:    "json",
				Aliases: []string{"j"},
				Usage:   "extract envs from vault end save as config.json file",
				Action: func(cCtx *cli.Context) error {
					if err := run(commands.CmdModeJson, envName, fileName); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(mod commands.CmdMod, envName, fileName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var loader envs.Loader
	if len(envName) != 0 {
		loader = envs.NewLoader(envs.WithFileName(envName))
	} else {
		loader = envs.NewLoader()
	}

	if err := loader.Load(); err != nil {
		return fmt.Errorf("run, loader Load, %w", err)
	}

	extractor, err := vault.NewClient(os.Getenv("VAULT_HOST"), os.Getenv("VAULT_TOKEN"))
	if err != nil {
		return fmt.Errorf("run, vault NewClient, %w", err)
	}

	cmd := commands.NewConfigFileMakerCommand(extractor)
	if err = cmd.MakeConfigFile(ctx, mod, envName, fileName); err != nil {
		return fmt.Errorf("run, cmd MakeConfigFile, %w", err)
	}

	return nil
}
