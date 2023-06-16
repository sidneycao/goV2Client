# goV2Client  
使用golang编写的V2ray客户端，支持添加/删除/更新订阅。
<br/>  
使用前请根据你们的情况修改 conf/conf.go 中以下三个变量的值  
```
// 保存订阅信息
var subConfigFile = "/usr/local/etc/goV2Config/sub.json"
// 保存节点列表信息
var nodeConfigFile = "/usr/local/etc/goV2Config/node.json"
// v2ray进程的配置文件
var v2rayConfigFile = "/usr/local/etc/v2ray/default.json"
```
请注意，本工具使用systemd来重启v2ray，请在使用前将v2ray添加到systemd  
```
cmd := exec.Command("systemctl", "restart", "v2ray")
err := cmd.Run()
if err != nil {
	log.Panic(err)
}
log.Println("success to restart v2ray process")
```  
```
另外，本工具只会更新v2ray配置文件中的"outbounds"字段内容，其余字段内容需要保存于另外的文件中，并在启动时使用多配置文件模式启动。  
```
/usr/local/bin/v2ray run -config /usr/local/etc/v2ray/inbounds.json -c /usr/local/etc/v2ray/default.json
```
订阅管理：  
  --sub --add(-a) {name} {url}  
    添加一个订阅，订阅后节点添加到node list  
  --sub --update(-u) {name}  
    更新订阅  
  --sub --del(-d) {name}  
    删除订阅  
  --sub --list(-l)   
    查看所有订阅  
节点管理：  
  --node --list(-l)  
    查看所有节点  
  --node --set(-s) {node_id}  
    使用该节点  
其他:  
  -h, --help  
    显示此帮助信息  
```