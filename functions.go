package refpkg

import (
	"go/importer"
	"go/token"
	"go/types"
	"strings"
)

type (
	// MemberInfo は package Member の情報です
	MemberInfo struct {
		Private  bool
		Name     string
		TypeInfo TypeInfo
	}
	// TypeInfo は 型情報です
	TypeInfo struct {
		Name      string
		BasedType BasedType
	}
	// BasedType は 基底の型です
	BasedType string
)

const (
	// BasedTypeStruct は struct です
	BasedTypeStruct BasedType = "struct"
	// BasedTypeInterface は interafce です
	BasedTypeInterface BasedType = "interafce"
	// BasedTypeFunc は func です
	BasedTypeFunc BasedType = "func"
	// BasedTypeConst は const です
	BasedTypeConst BasedType = "const"
	// BasedTypeVar は var です
	BasedTypeVar BasedType = "var"
)

var imp = importer.ForCompiler(token.NewFileSet(), "source", nil)

// GetMembers は package を指定することで package に定義された情報を取得します。
// pkgName は "github.com/wakuwaku3/refpkg/sample" の様に指定します。
func GetMembers(pkgName string) ([]MemberInfo, error) {
	pkg, err := imp.Import(pkgName)
	if err != nil {
		return nil, err
	}

	scope := pkg.Scope()
	names := scope.Names()
	members := make([]MemberInfo, len(names))
	for i, name := range names {
		obj := scope.Lookup(name)
		members[i] = *cnvObjToMemberInfo(pkgName, obj)
	}
	return members, nil
}

func cnvObjToMemberInfo(pkgName string, obj types.Object) *MemberInfo {
	return &MemberInfo{
		Private:  !obj.Exported(),
		Name:     obj.Name(),
		TypeInfo: *cnvType(pkgName, obj.Type()),
	}
}

func cnvType(pkgName string, t types.Type) *TypeInfo {
	return &TypeInfo{
		Name:      strings.ReplaceAll(strings.ReplaceAll(t.String(), pkgName+".", ""), "untyped ", ""),
		BasedType: getBasedType(t, ""),
	}
}

func getBasedType(t types.Type, prevName string) BasedType {
	name := t.String()
	if name == prevName {
		if strings.HasPrefix(name, "untyped") {
			return BasedTypeConst
		}
		return BasedTypeVar
	} else if strings.HasPrefix(name, "interface{") {
		return BasedTypeInterface
	} else if strings.HasPrefix(name, "struct{") {
		return BasedTypeStruct
	} else if strings.HasPrefix(name, "func(") {
		return BasedTypeFunc
	}
	return getBasedType(t.Underlying(), name)
}
