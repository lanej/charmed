// Package lexical maps Chroma's lexical classifications to Charmed's theme.
package lexical

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/lanej/charmed/theme"
)

// Chroma groups these literals as KeywordConstant. Keep their private token
// types in the keyword category so ordinary keyword styling still inherits.
const (
	True chroma.TokenType = 1090 + iota
	False
	Null
)

// These refinements stay in their original Chroma categories for profiles that
// use the broader name and punctuation roles.
const (
	goBuiltinType chroma.TokenType = 2090 + iota
	goBuiltinFunction
	shellVariableReference
	goOperator chroma.TokenType = 5090
)

// Kind refines constant literals and Go operators, builtin calls and conversions.
func Kind(token chroma.Token, language string) chroma.TokenType {
	if token.Type == chroma.KeywordConstant {
		switch token.Value {
		case "true", "True":
			return True
		case "false", "False":
			return False
		case "null", "nil", "None":
			return Null
		}
	}

	if strings.EqualFold(language, "go") || strings.EqualFold(language, "golang") {
		switch {
		case token.Type == chroma.NameBuiltin && GoType(token.Value):
			return goBuiltinType
		case token.Type == chroma.NameBuiltin:
			return goBuiltinFunction
		case token.Type == chroma.Punctuation && len(token.Value) == 1 && strings.Contains("|^<>=!~", token.Value):
			return goOperator
		}
	}

	if token.Type == chroma.NameVariable {
		switch strings.ToLower(language) {
		case "shell", "sh", "bash", "zsh", "ksh", "fish":
			return shellVariableReference
		}
	}

	return token.Type
}

// GoType recognizes the predeclared types in source snippets and conversions.
func GoType(name string) bool {
	switch name {
	case "any", "comparable", "bool", "byte", "rune", "string", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "float32", "float64", "complex64", "complex128", "error":
		return true
	default:
		return false
	}
}

// Iterator preserves the original types for a caller-supplied Chroma style.
func Iterator(tokens chroma.Iterator, style *chroma.Style, language string) chroma.Iterator {
	return func() chroma.Token {
		token := tokens()
		if kind := Kind(token, language); kind != token.Type && style.Has(kind) {
			token.Type = kind
		}

		return token
	}
}

// Color returns a specific lexical role when one is configured. Otherwise
// callers retain their existing category defaults or Chroma inheritance.
func Color(kind chroma.TokenType, t theme.Theme) color.Color {
	switch kind { //nolint:exhaustive // Unlisted Chroma kinds inherit the existing lexical category.
	case chroma.NameFunction, chroma.NameFunctionMagic, goBuiltinFunction:
		return t.Lexical.Function
	case chroma.KeywordType, chroma.NameClass, goBuiltinType:
		return t.Lexical.Type
	case chroma.Operator, chroma.OperatorWord, chroma.NameOperator, goOperator:
		return t.Lexical.Operator
	case True:
		return t.Lexical.BooleanTrue
	case False:
		return t.Lexical.BooleanFalse
	case Null:
		return t.Lexical.Null
	case chroma.NameConstant, chroma.KeywordConstant:
		return t.Lexical.Constant
	case chroma.Name, chroma.NameOther, chroma.NameVariable:
		return t.Lexical.Variable
	case shellVariableReference:
		return t.Lexical.VariableReference
	case chroma.NameBuiltinPseudo, chroma.NameVariableMagic:
		return t.Lexical.Receiver
	case chroma.LiteralStringEscape, chroma.LiteralStringInterpol, chroma.NameDecorator:
		return t.Lexical.Special
	case chroma.GenericHeading, chroma.GenericSubheading:
		return t.Lexical.Section
	default:
		return nil
	}
}

// Entries adds only configured roles, preserving a partial theme's inherited styles.
func Entries(entries chroma.StyleEntries, t theme.Theme) {
	for _, kinds := range [][]chroma.TokenType{chroma.TokenTypeValues(), {True, False, Null, goBuiltinType, goBuiltinFunction, goOperator, shellVariableReference}} {
		for _, kind := range kinds {
			if c := Color(kind, t); c != nil {
				entries[kind] = hex(c)
			}
		}
	}
	// Name is also Chroma's generic category. Restore the more specific key
	// and member roles after assigning ordinary variables their own color.
	if t.Lexical.Variable != nil && t.Syntax.Name != nil {
		for _, kind := range []chroma.TokenType{chroma.NameTag, chroma.NameAttribute, chroma.NameProperty, chroma.NameVariableClass, chroma.NameVariableInstance} {
			entries[kind] = hex(t.Syntax.Name)
		}
	}
}

func hex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}
