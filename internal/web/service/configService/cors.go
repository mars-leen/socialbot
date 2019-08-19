package configService

var (
	cors       = &Cors{
		Enable:           true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		//AllowHeaders:     []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Authorization", "x-requested-with"},
		AllowHeaders:     []string{"*"},
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

func GetCors() *Cors {
	return cors
}
