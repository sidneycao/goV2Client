package conf

type V2Node struct {
	SubName    string
	Vmess      VmessJson //解析后的vmess json结构
	Source     string    //原始vmess链接  vmess://
	ConfigJson string    //通过解析后的vmess json 转换得到的 v2ray config
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

/**
func ParseVmess2Conf() {

}
**/
