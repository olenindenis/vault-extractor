package commands

import (
	"context"
	"fmt"
	"os"

	"vault-extractor/internal/converters"
)

type CmdMod string

const (
	CmdModeEnv  CmdMod = "env"
	CmdModeJson CmdMod = "json"
)

type (
	Extractor interface {
		Extract(ctx context.Context, path, mountPath string) (map[string]interface{}, error)
	}
	ConfigFileMakerCommand struct {
		extractor Extractor
	}
)

func NewConfigFileMakerCommand(extractor Extractor) *ConfigFileMakerCommand {
	return &ConfigFileMakerCommand{
		extractor: extractor,
	}
}

func (m *ConfigFileMakerCommand) MakeConfigFile(ctx context.Context, mod CmdMod, envName, fileName string) error {
	data, err := m.extractor.Extract(ctx, os.Getenv("VAULT_PATH"), os.Getenv("VAULT_MOUNT_PATH"))
	if err != nil {
		return fmt.Errorf("MakeConfigFile, vault extract, %w", err)
	}

	if len(data) == 0 {
		return fmt.Errorf("MakeConfigFile, empty data, %w", err)
	}

	if mod == CmdModeEnv {
		// extractor add envs used for vault client to you new env fileName
		if err = converters.SaveAsEnvFile(ctx, envName, fileName, data); err != nil {
			return fmt.Errorf("MakeConfigFile, converters SaveAsEnvFile, %w", err)
		}
	}

	if mod == CmdModeJson {
		if err = converters.SaveAsJsonFile(ctx, fileName, data); err != nil {
			return fmt.Errorf("MakeConfigFile, converters SaveAsJsonFile, %w", err)
		}
	}

	return nil
}
