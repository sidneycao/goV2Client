package conf

type V2Sub struct {
	Name string
	Url  string
}

var SubConfig map[string]V2Sub = map[string]V2Sub{}
