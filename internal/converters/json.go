package converters

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

func SaveAsJsonFile(_ context.Context, fileName string, data map[string]interface{}) error {
	tmpStruct := make(map[string]string, len(data))
	for k, v := range data {
		tmpStruct[k] = v.(string)
	}

	jsonString, err := json.Marshal(tmpStruct)
	if err != nil {
		return fmt.Errorf("MakeEnvFile, json.Marshal, %w", err)
	}

	if err = os.WriteFile(fileName, jsonString, os.FileMode(0666)); err != nil {
		return fmt.Errorf("MakeEnvFile, os.WriteFile, %w, %s", err, fileName)
	}

	return nil
}
