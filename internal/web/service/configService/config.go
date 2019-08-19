package configService


const (
	HostKey = "host"
	ServerKey= "server"
	CorsKey = "cors"
	StorageKey = "storage"
)

var (
	//host   = "http://127.0.0.1:8081"
	host   = "http://192.168.31.228:8081"
	hostName       = "century"
)
func GetHost() string {
	return host
}
func GetHostName() string {
	return hostName
}
