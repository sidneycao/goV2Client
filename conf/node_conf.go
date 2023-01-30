package conf

type V2Node struct {
	SubName    string
	Vmess      VmessJson
	Source     string
	ConfigJson string
}

type VmessJson struct {
	Ps   string `json:"ps"`
	Add  string `json:"add"`
	Port string `json:"port"`
	Id   string `json:"id"`
	Aid  string `json:"aid"`
	Net  string `json:"net"`
	Type string `json:"type"`
	//TLS  string `json:"tls"`
}
