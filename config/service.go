package config

const (
	//account rpc服务监听的端口
	AccountServicePort = ":29000"
	//account服务客户端连接rpc地址
	AccountServiceIngress = "127.0.0.1:29000"

	//Download rpc服务监听的端口
	DownloadServicePort = ":29001"
	//Download服务客户端连接rpc地址
	DownloadServiceIngress = "127.0.0.1:29001"

	//upload rpc服务监听的端口
	UploadServicePort = ":29002"
	//account服务客户端连接rpc地址
	UploadServiceIngress = "127.0.0.1:29002"

	//dbproxy rpc服务监听的端口
	DbproxyServicePort = ":29003"
	//dbproxy服务客户端连接rpc地址
	DbproxyServiceIngress = "127.0.0.1:29003"
)
