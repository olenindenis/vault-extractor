package converters

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/wk8/go-ordered-map/v2"
)

func SaveAsEnvFile(_ context.Context, envName, fileName string, data map[string]interface{}) error {
	var buffer bytes.Buffer
	existsData := orderedmap.New[string, string]()

	if _, err := os.Stat(envName); !errors.Is(err, os.ErrNotExist) {
		file, err := os.Open(fileName)
		if err != nil {
			return fmt.Errorf("MakeEnvFile, os.Open, %w, %s", err, envName)
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		var (
			key string
			val string
		)
		for {
			line, _, err := reader.ReadLine()

			if err == io.EOF {
				break
			}

			splits := strings.Split(string(line), "=")
			key = splits[0]
			val = splits[1]

			existsData.Set(key, val)
		}

		for pair := existsData.Oldest(); pair != nil; pair = pair.Next() {
			buffer.WriteString(fmt.Sprintf("%s=%s\n", pair.Key, pair.Value))
		}
	}

	for k, v := range data {
		oldVal, ok := existsData.Get(k)
		if ok {
			continue
		}

		if oldVal == v.(string) {
			continue
		}

		buffer.WriteString(fmt.Sprintf("%s=%s\n", k, v.(string)))
	}

	if err := os.WriteFile(fileName, buffer.Bytes(), os.FileMode(0666)); err != nil {
		return fmt.Errorf("MakeEnvFile, os.WriteFile, %w, %s", err, fileName)
	}

	return nil
}
