package main

const AppId string = "ucasnauth"

const HostUrl string = "http://210.77.16.21/"

const IndexPageBasename string = "index.jsp"
const SuccessPageBasename string = "success.jsp"
const InterFaceName string = "InterFace.do"

const PostContentType string = "application/x-www-form-urlencoded;charset=UTF-8"

const AuthBodyFormat string = "userId=%s&password=%s&service=&queryString=%s" +
	"&operatorPwd=&operatorUserId=&validcode=&passwordEncrypt=true"

const RsaPubKeyModHex string = "0x9C2899B8CEDDF9BEAFAD2DB8E431884A" +
	"79FD9B9C881E459C0E1963984779D661" +
	"2222CEE814593CC458845BBBA42B2D34" +
	"74C10B9D31ED84F256C6E3A1C795E68E" +
	"18585B84650076F122E763289A4BCB0D" +
	"E08762C3CEB591EC44D764A69817318F" +
	"BCE09D6ECB0364111F6F38E90DC44CA8" +
	"9745395A17483A778F1CC8DC990D87C3"
const RsaPubKeyExp int = 0x10001

const DataFilename string = "data.dat"
const NonceFilename string = "nonce.dat"
const SaltFilename string = "ucasnauthsalt.dat"

const UsageHint string = `Usage:
  UCASNAUTH [login/logout] [username password]
  If the command "login" or "logout" is not given, login will be executed.
  username and password are only valid for login.
  If username and password are not given, the value in last successfully
authentication will be used.
  You must specify username and password when you use it for the first time,
or when the saved data is corrupt or expired.`
