package api

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/easyp-tech/easyp/internal/config"
	"github.com/easyp-tech/easyp/internal/flags"
	wfs "github.com/easyp-tech/easyp/internal/fs"
)

var _ Handler = (*Init)(nil)

// Init is a handler for initialization EasyP configuration.
type Init struct{}

var (
	flagInitDirectoryPath = &cli.StringFlag{
		Name:       "dir",
		Usage:      "directory path to initialize",
		Required:   true,
		HasBeenSet: true,
		Value:      ".",
		Aliases:    []string{"d"},
		EnvVars:    []string{"EASYP_INIT_DIR"},
	}
)

// Command implements Handler.
func (i Init) Command() *cli.Command {
	return &cli.Command{
		Name:        "init",
		Aliases:     []string{"i"},
		Usage:       "initialize configuration",
		UsageText:   "initialize configuration",
		Description: "initialize configuration",
		Action:      i.Action,
		Flags: []cli.Flag{
			flagInitDirectoryPath,
		},
	}
}

// Action implements Handler.
func (i Init) Action(ctx *cli.Context) error {
	rootPath := ctx.String(flagInitDirectoryPath.Name)
	dirFS := wfs.Disk(rootPath)

	cfg, err := config.New(ctx.Context, ctx.String(flags.Config.Name))
	if err != nil {
		return fmt.Errorf("config.New: %w", err)
	}

	app, err := buildCore(ctx.Context, *cfg)
	if err != nil {
		return fmt.Errorf("buildCore: %w", err)
	}

	err = app.Initialize(ctx.Context, dirFS, []string{"DEFAULT"})
	if err != nil {
		return fmt.Errorf("initer.Initialize: %w", err)
	}

	return nil
}
