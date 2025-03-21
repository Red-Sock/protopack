package rules

import (
	"path/filepath"

	"go.redsock.ru/protopack/internal/core"
)

var _ core.Rule = (*DirectorySamePackage)(nil)

// DirectorySamePackage this rule checks that all files in a given directory are in the same package.
type DirectorySamePackage struct {
	// dir => package
	cache map[string]string
}

func (d *DirectorySamePackage) lazyInit() {
	if d.cache == nil {
		d.cache = make(map[string]string)
	}
}

// Message implements lint.Rule.
func (d *DirectorySamePackage) Message() string {
	return "all files in the same directory must have the same package name"
}

// Validate implements lint.Rule.
func (d *DirectorySamePackage) Validate(protoInfo core.ProtoInfo) ([]core.Issue, error) {
	d.lazyInit()
	var res []core.Issue

	directory := filepath.Dir(protoInfo.Path)
	for _, pack := range protoInfo.Info.ProtoBody.Packages {
		if d.cache[directory] == "" {
			d.cache[directory] = pack.Name
			continue
		}

		if d.cache[directory] != pack.Name {
			res = core.AppendIssue(res, d, pack.Meta.Pos, pack.Name, pack.Comments)
		}
	}

	return res, nil
}
