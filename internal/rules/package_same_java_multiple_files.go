package rules

import (
	"go.redsock.ru/protopack/internal/core"
)

var _ core.Rule = (*PackageSameJavaMultipleFiles)(nil)

// PackageSameJavaMultipleFiles checks that all files with a given package have the same value for the java_multiple_files option.
type PackageSameJavaMultipleFiles struct {
	// dir => package
	cache map[string]string
}

func (p *PackageSameJavaMultipleFiles) lazyInit() {
	if p.cache == nil {
		p.cache = make(map[string]string)
	}
}

// Message implements lint.Rule.
func (p *PackageSameJavaMultipleFiles) Message() string {
	return "all files in the same package must have the same java_multiple_files option"
}

// Validate implements lint.Rule.
func (p *PackageSameJavaMultipleFiles) Validate(protoInfo core.ProtoInfo) ([]core.Issue, error) {
	p.lazyInit()

	var res []core.Issue

	if len(protoInfo.Info.ProtoBody.Packages) == 0 {
		return res, nil
	}

	packageName := protoInfo.Info.ProtoBody.Packages[0].Name
	for _, option := range protoInfo.Info.ProtoBody.Options {
		if option.OptionName == "java_multiple_files" {
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
