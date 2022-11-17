// Package xmlstruct generates Go structs from multiple XML documents.
package xmlstruct

import (
	"encoding/xml"
	"unicode"

	"golang.org/x/exp/slices"
)

const (
	DefaultFormatSource                 = true
	DefaultHeader                       = "// This file is automatically generated. DO NOT EDIT."
	DefaultIntType                      = "int"
	DefaultPackageName                  = "main"
	DefaultTimeLayout                   = "2006-01-02T15:04:05Z"
	DefaultUsePointersForOptionalFields = true
)

var (
	DefaultExportNameFunc = TitleFirstRune
	DefaultNameFunc       = IgnoreNamespaceNameFunc
)

// An ExportNameFunc returns the exported Go identifier for the given xml.Name.
type ExportNameFunc func(xml.Name) string

// A NameFunc modifies xml.Names observed in the XML documents.
type NameFunc func(xml.Name) xml.Name

// observeOptions contains options for observing XML documents.
type observeOptions struct {
	nameFunc   NameFunc
	timeLayout string
}

// generateOptions contains options for generating Go source.
type generateOptions struct {
	exportNameFunc               ExportNameFunc
	header                       string
	importPackageNames           map[string]struct{}
	intType                      string
	usePointersForOptionalFields bool
}

// TitleFirstRune returns name.Local with the initial rune capitalized.
func TitleFirstRune(name xml.Name) string {
	runes := []rune(name.Local)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// The IdentityNameFunc returns name unchanged. The same local name in different
// namespaces will be treated as distinct names.
func IdentityNameFunc(name xml.Name) xml.Name {
	return name
}

// IgnoreNamespaceNameFunc returns name with name.Space cleared. The same local
// name in different namespaces will be treated as identical names.
func IgnoreNamespaceNameFunc(name xml.Name) xml.Name {
	return xml.Name{
		Local: name.Local,
	}
}

// sortXMLNames sorts xmlNames by namespace and then local name and returns the
// sorted slice.
func sortXMLNames(xmlNames []xml.Name) []xml.Name {
	slices.SortFunc(xmlNames, func(a, b xml.Name) bool {
		switch {
		case a.Space < b.Space:
			return true
		case a.Space == b.Space:
			return a.Local < b.Local
		default:
			return false
		}
	})
	return xmlNames
}
