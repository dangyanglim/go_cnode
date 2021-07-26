module go_cnode

go 1.12

replace golang.org/x/net => github.com/golang/net v0.0.0-20190424024845-afe8014c977f

replace golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190422165155-953cdadca894

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190422183909-d864b10871cd

replace golang.org/x/text => github.com/golang/text v0.3.0

replace google.golang.org/appengine => github.com/golang/appengine v1.5.0

require (
	github.com/garyburd/redigo v1.6.0
	github.com/gin-contrib/cors v0.0.0-20190424000812-bd1331c62cae
	github.com/gin-gonic/gin v1.7.0
	github.com/gorilla/sessions v1.1.3 // indirect
	github.com/russross/blackfriday v2.0.0+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/tommy351/gin-sessions v0.0.0-20150617141853-353060947eb6
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)
