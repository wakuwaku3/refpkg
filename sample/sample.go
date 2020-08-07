package sample

type (
	// PublicStruct です
	PublicStruct struct{}
	// PublicInterface です
	PublicInterface  interface{}
	privateStruct    struct{}
	privateInterface interface{}
)

const (
	// PublicConst です
	PublicConst  = ""
	privateConst = ""
)

var (
	// PublicVar です
	PublicVar  = ""
	privateVar = ""
)

// PublicFunc です
func PublicFunc()  {}
func privateFunc() {}
