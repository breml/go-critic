package lint

import (
	"go/ast"
	"regexp"
)

type commentsChecker struct {
	ctx          *context
	badCommentRE *regexp.Regexp
}

func commentsCheck(ctx *context) func(*ast.File) {
	re := `//\s?\w+[^a-zA-Z]+$`
	c := commentsChecker{ctx: ctx, badCommentRE: regexp.MustCompile(re)}
	return c.CheckFile
}

func (c *commentsChecker) CheckFile(f *ast.File) {
	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.FuncDecl); ok {
			if decl.Doc != nil && c.badCommentRE.MatchString(decl.Doc.List[0].Text) {
				c.warn(decl)
			}
		}
	}
}

func (c *commentsChecker) warn(decl *ast.FuncDecl) {
	c.ctx.Warn(decl, "silencing go lint doc-comment warnings is unadvised")
}