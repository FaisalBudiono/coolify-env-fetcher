package mapper

import (
	"fmt"
	"io"

	"FaisalBudiono/coolify-env-fetcher/internal/coolify"
)

type dotENV struct{}

func NewDotENV() *dotENV {
	return &dotENV{}
}

func (d *dotENV) WriteFile(
	w io.Writer, es []coolify.EnvObject,
	isPreview bool,
) error {
	for _, e := range es {
		skipped := func() bool {
			if !e.IsBuildENV() {
				return true
			}

			if isPreview {
				return !e.IsPreview
			}

			return e.IsPreview
		}()

		if skipped {
			continue
		}

		con := fmt.Sprintf("%s=%s\n", e.Key, e.RealValue)
		_, err := w.Write([]byte(con))
		if err != nil {
			return err
		}
	}

	return nil
}
