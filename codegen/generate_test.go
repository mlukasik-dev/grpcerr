package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Run("generate not empty file without errors", func(t *testing.T) {
		t.Parallel()

		outdir, err := os.MkdirTemp("", "")
		defer os.RemoveAll(outdir)
		if err != nil {
			t.Fatalf("failed to setup test: %v", err)
		}
		if err := Generate("../testdata/grpc-errors.yaml", "./gen"); err != nil {
			t.Fatal(err)
		}

		gen, err := ioutil.ReadFile(filepath.Join(outdir, "gen.go"))
		if err != nil {
			t.Error(err)
		}
		if len(gen) == 0 {
			t.Errorf("generated empty file")
		}
	})
}
