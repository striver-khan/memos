package parser

import (
	"github.com/usememos/memos/plugin/gomark/ast"
	"github.com/usememos/memos/plugin/gomark/parser/tokenizer"
)

type TagParser struct{}

func NewTagParser() *TagParser {
	return &TagParser{}
}

func (*TagParser) Match(tokens []*tokenizer.Token) (ast.Node, int) {
	matchedTokens := tokenizer.GetFirstLine(tokens)
	if len(matchedTokens) < 2 {
		return nil, 0
	}
	if matchedTokens[0].Type != tokenizer.PoundSign {
		return nil, 0
	}

	contentTokens := []*tokenizer.Token{}
	for _, token := range matchedTokens[1:] {
		if token.Type == tokenizer.Space || token.Type == tokenizer.PoundSign {
			break
		}
		contentTokens = append(contentTokens, token)
	}
	if len(contentTokens) == 0 {
		return nil, 0
	}

	return &ast.Tag{
		Content: tokenizer.Stringify(contentTokens),
	}, len(contentTokens) + 1
}
