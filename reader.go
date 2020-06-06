package archdiag

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func (t *ArchDiag) read(reader io.Reader, ext Format) error {
	bs, err := ioutil.ReadAll(reader)

	if err != nil {
		return err
	}

	switch Format(ext) {
	case FormatJson:
		if err := json.Unmarshal(bs, &t.G); err != nil {
			return err
		}
	case FormatYaml:
		if err := yaml.Unmarshal(bs, &t.G); err != nil {
			return err
		}
	}

	return nil
}

func (t *ArchDiag) ReadFromFile(path string) error {
	ext := filepath.Ext(path)
	ext = strings.Replace(ext, ".", "", -1)

	var format Format

	switch ext {
	case "json":
		format = FormatJson
	case "yml":
		format = FormatYaml
	case "yaml":
		format = FormatYaml
	default:
		return fmt.Errorf("unsupposed file extention '%v'", ext)
	}

	f, err := os.Open(path)

	if err != nil {
		return err
	}

	defer f.Close()

	if err := t.read(f, format); err != nil {
		return fmt.Errorf("parse error %v: %w", path, err)
	}

	return nil
}
