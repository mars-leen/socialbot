package configService

var (
	host       = ""
	uploadHost = ""
	cors       = Cors{
		Enable:           true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           7200,
	}
)

type Cors struct {
	Enable           bool     `json:"Enable"`
	AllowOrigins     []string `json:"AllowOrigins"`
	AllowMethods     []string `json:"AllowMethods"`
	AllowHeaders     []string `json:"AllowHeaders"`
	AllowCredentials bool     `json:"AllowCredentials"`
	MaxAge           int      `json:"MaxAge"`
}

func GetCors() Cors {
	return cors
}


func GetHostName() string {
	return host
}

func SetUploadHost(host string) {
	uploadHost = host
}

func GetUploadHost() string {
	return uploadHost
}

func GetUploadFullUrl(base string) string {
	return uploadHost + base
}
