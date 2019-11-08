package utils
import(
	"os"
	"encoding/json"
)

type configuration struct {
	Port      string
	Mongo_url string
	Redis_url string
	Smtp_username string
	Smtp_password string
	Smtp_hostname string
	Smtp_active_Url string
	Github_client_id string
    Github_client_secret string
    Github_AuthURL string 
    Github_UserURL string
    Github_TokenURL string
}
func LoadConf() (conf configuration) {
	// 打开文件
	file, _ := os.Open("conf.json")

	// 关闭文件
	defer file.Close()

	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)

	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	decoder.Decode(&conf)
	return conf
}