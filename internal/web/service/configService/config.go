package configService


const (
	HostKey = "host"
	ServerKey= "server"
	CorsKey = "cors"
	StorageKey = "storage"
)

var (
	host   = "http://localhost:8080"
	hostName       = "century"
)
func GetHost() string {
	return host
}
func GetHostName() string {
	return hostName
}
