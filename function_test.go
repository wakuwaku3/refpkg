package refpkg

import (
	"reflect"
	"testing"
)

func TestGetMembers(t *testing.T) {
	actual, err := GetMembers("github.com/wakuwaku3/refpkg/sample")
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]MemberInfo{
		"PublicStruct": {Private: false, Name: "PublicStruct", TypeInfo: TypeInfo{
			Name:      "PublicStruct",
			BasedType: BasedTypeStruct,
		}},
		"PublicInterface": {Private: false, Name: "PublicInterface", TypeInfo: TypeInfo{
			Name:      "PublicInterface",
			BasedType: BasedTypeInterface,
		}},
		"PublicConst": {Private: false, Name: "PublicConst", TypeInfo: TypeInfo{
			Name:      "string",
			BasedType: BasedTypeConst,
		}},
		"PublicVar": {Private: false, Name: "PublicVar", TypeInfo: TypeInfo{
			Name:      "string",
			BasedType: BasedTypeVar,
		}},
		"PublicFunc": {Private: false, Name: "PublicFunc", TypeInfo: TypeInfo{
			Name:      "func()",
			BasedType: BasedTypeFunc,
		}},
		"privateStruct": {Private: true, Name: "privateStruct", TypeInfo: TypeInfo{
			Name:      "privateStruct",
			BasedType: BasedTypeStruct,
		}},
		"privateInterface": {Private: true, Name: "privateInterface", TypeInfo: TypeInfo{
			Name:      "privateInterface",
			BasedType: BasedTypeInterface,
		}},
		"privateConst": {Private: true, Name: "privateConst", TypeInfo: TypeInfo{
			Name:      "string",
			BasedType: BasedTypeConst,
		}},
		"privateVar": {Private: true, Name: "privateVar", TypeInfo: TypeInfo{
			Name:      "string",
			BasedType: BasedTypeVar,
		}},
		"privateFunc": {Private: true, Name: "privateFunc", TypeInfo: TypeInfo{
			Name:      "func()",
			BasedType: BasedTypeFunc,
		}},
	}
	if len(actual) != len(expected) {
		t.Fatal(actual)
	}
	for i, a := range actual {
		if e, ok := expected[a.Name]; !ok || !reflect.DeepEqual(a, e) {
			t.Fatal(i, a, e)
		}
	}
}
