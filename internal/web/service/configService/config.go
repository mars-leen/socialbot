package configService


const (
	HostKey = "host"
	ServerKey= "server"
	CorsKey = "cors"
	StorageKey = "storage"
)

var (
	host       = ""
)
func GetHostName() string {
	return host
}
