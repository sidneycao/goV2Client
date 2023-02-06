package conf

type VSub struct {
	Name string
	Url  string
}

var SubConfigNow map[string]VSub = map[string]VSub{}
