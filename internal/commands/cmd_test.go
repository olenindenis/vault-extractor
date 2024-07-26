package commands

import (
	"context"
	"os"
	"testing"

	"github.com/olenindenis/vault-extractor/pkg/vault"
)

func TestMakeConfigFileForEnvType(t *testing.T) {
	testFileName := ".env.test"
	ctx := context.TODO()

	extractor := vault.NewNilClient()
	cmd := NewConfigFileMakerCommand(extractor)
	err := cmd.MakeConfigFile(ctx, CmdModeEnv, testFileName, testFileName)
	if err != nil {
		t.Fatalf(`MakeConfigFile("%s", "%s") = %v, want match for nil`, CmdModeEnv, testFileName, err)
	}

	if err = os.Remove(testFileName); err != nil {
		t.Fatalf(`Remove("%s") = %v, want match for nil`, testFileName, err)
	}
}

func TestMakeConfigFileForJsonType(t *testing.T) {
	testFileName := "tmp.json"
	ctx := context.TODO()

	extractor := vault.NewNilClient()
	cmd := NewConfigFileMakerCommand(extractor)
	err := cmd.MakeConfigFile(ctx, CmdModeJson, testFileName, testFileName)
	if err != nil {
		t.Fatalf(`MakeConfigFile("%s", "%s") = %v, want match for nil`, CmdModeJson, testFileName, err)
	}

	if err = os.Remove(testFileName); err != nil {
		t.Fatalf(`Remove("%s") = %v, want match for nil`, testFileName, err)
	}
}
