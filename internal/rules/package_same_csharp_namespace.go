package rules

import (
	"go.redsock.ru/protopack/internal/core"
)

var _ core.Rule = (*PackageSameCsharpNamespace)(nil)

// PackageSameCsharpNamespace checks that all files with a given package have the same value for the csharp_namespace option.
type PackageSameCsharpNamespace struct {
	// dir => package
	cache map[string]string
}

func (p *PackageSameCsharpNamespace) lazyInit() {
	if p.cache == nil {
		p.cache = make(map[string]string)
	}
}

// Message implements lint.Rule.
func (p *PackageSameCsharpNamespace) Message() string {
	return "different proto files in the same package should have the same csharp_namespace"
}

// Validate implements lint.Rule.
func (p *PackageSameCsharpNamespace) Validate(protoInfo core.ProtoInfo) ([]core.Issue, error) {
	p.lazyInit()

	var res []core.Issue

	if len(protoInfo.Info.ProtoBody.Packages) == 0 {
		return res, nil
	}

	packageName := protoInfo.Info.ProtoBody.Packages[0].Name
	for _, option := range protoInfo.Info.ProtoBody.Options {
		if option.OptionName == "csharp_namespace" {
			if p.cache[packageName] == "" {
				p.cache[packageName] = option.Constant
				continue
			}

			if p.cache[packageName] != option.Constant {
				res = core.AppendIssue(res, p, option.Meta.Pos, option.Constant, option.Comments)
			}
		}
	}

	return res, nil
}
