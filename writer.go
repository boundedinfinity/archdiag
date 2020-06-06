package archdiag

import (
	"io"
	"os"
)

func (t *ArchDiag) write(w io.Writer) error {
	if err := t.Generate(w); err != nil {
		return err
	}

	return nil
}

func (t *ArchDiag) WriteToStdout() error {
	return t.write(os.Stdout)
}

func (t *ArchDiag) WriteToStderr() error {
	return t.write(os.Stderr)
}

func (t *ArchDiag) WriteToFile(path string) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	if err := t.Generate(f); err != nil {
		return err
	}

	return nil
}
