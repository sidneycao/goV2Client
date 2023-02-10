# goV2Client  
使用golang编写的V2ray客户端，支持添加/删除/更新订阅。  
  
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