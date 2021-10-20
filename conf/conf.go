package conf

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/hitzhangjie/babynames/assets"

	"github.com/hitzhangjie/codeblocks/tar"
)

var CfgPath = filepath.Join(os.TempDir(), "babynames")

func init() {
	dat := assets.DictsGo

	os.RemoveAll(CfgPath)

	err := tar.Untar(CfgPath, bytes.NewBuffer(dat))
	if err != nil {
		panic(err)
	}
}
