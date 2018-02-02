package filter
import (
	"github.com/astaxie/beego/session"
)
type Sesson struct {

}
var GlobalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	GlobalSessions, _ = session.NewManager("memory",sessionConfig)
	go GlobalSessions.GC()
}