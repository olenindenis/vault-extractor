package envs

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const defaultEnvFileName = ".env"

type Loader interface {
	Load() error
}

var _ Loader = (*DefaultLoader)(nil)

type DefaultLoader struct {
	fileName string
}

func NewLoader(options ...func(*DefaultLoader)) *DefaultLoader {
	loader := &DefaultLoader{
		fileName: defaultEnvFileName,
	}
	for _, o := range options {
		o(loader)
	}
	return loader
}

func WithFileName(fileName string) func(loader *DefaultLoader) {
	return func(l *DefaultLoader) {
		l.fileName = fileName
	}
}

func (l *DefaultLoader) Load() error {
	if _, err := os.Stat(l.fileName); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	var fileEnv map[string]string
	fileEnv, err := godotenv.Read()
	if err != nil {
		return fmt.Errorf("godotenv read: %w", err)
	}

	for key, val := range fileEnv {
		if len(os.Getenv(key)) == 0 {
			err = os.Setenv(key, val)
			if err != nil {
				return fmt.Errorf("os setenv: %w", err)
			}
		}
	}

	return nil
}
