package file

import (
	"github.com/magiconair/properties"
	"log"
	"os"
	"strings"
)

type Config struct {
	DatasourceHost     string `properties:"datasource.host,default=localhost"`
	DatasourcePort     string `properties:"datasource.port,default=5432"`
	DatasourceUser     string `properties:"datasource.user,default=postgres"`
	DatasourcePassword string `properties:"datasource.password,default=pgadmin"`
	DatasourceTable    string `properties:"datasource.table,default=postgres"`
	TokenExpired       int64  `properties:"token.expired,default=1800"`
	TokenRefreshCron   string `properties:"token.refresh.cron,default=*/30 * * * * ?"`
	TokenSecret        string `properties:"token.secret,default=dahuawudimeimeida,jishufeichanghao,xinxiangshicheng,wanshiruyi"`
}

var cfg Config

func init() {
	filePath := os.Getenv("config.file.path")
	if strings.Compare(filePath, "") == 0 {
		filePath = "config.properties"
	} else {
		if !strings.HasSuffix(filePath, "properties") {
			suffix := "config.properties"
			if !strings.HasSuffix(filePath, "\\/") {
				suffix = "\\/" + suffix
			}
			filePath += suffix
		}
	}
	config := properties.MustLoadFile(filePath, properties.UTF8)
	if err := config.Decode(&cfg); err != nil {
		log.Fatal(err)
	}
}

func Read(fileName string) string {
	bytes, _ := os.ReadFile(fileName)
	return string(bytes)
}

func Write(fileName string, fileData string) {
	data := []byte(fileData)
	os.WriteFile(fileName, data, 0664)
}

func GetEnvParam() Config {
	return cfg
}
