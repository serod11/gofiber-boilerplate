/*
 * @author serod
 */

package config

type Pgsql struct {
	Host        string
	ReplicaHost string
	Port        string
	Config      string
	Dbname      string
	Username    string
	Password    string
	LogMode     string
	LogZap      bool
}

func (p *Pgsql) Dsn() string {
	return "host=" + p.Host + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " TimeZone=Asia/Ulaanbaatar" + p.Config
}
