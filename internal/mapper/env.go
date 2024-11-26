package mapper

import (
	"FaisalBudiono/coolify-env-fetcher/internal/coolify"
	"fmt"
	"io"
)

type dotENV struct{}

func NewDotENV() *dotENV {
	return &dotENV{}
}

func (d *dotENV) WriteFile(w io.Writer, es []coolify.EnvObject) error {
	for _, e := range es {
		if !e.IsBuildTime {
			continue
		}

		con := fmt.Sprintf("%s=%s\n", e.Key, e.Value)
		_, err := w.Write([]byte(con))
		if err != nil {
			return err
		}
	}

	return nil
}
