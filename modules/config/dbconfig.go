package config

type DatabaseConfig struct {
	Host       string
	Port       string
	User       string
	Pwd        string
	Name       string
	MaxIdleCon int
	MaxOpenCon int
	Dsn        string
	Params     map[string]string
}

type DatabaseList map[string]DatabaseConfig

func (dbc DatabaseConfig) GetDSN() string {
	if dbc.Dsn != "" {
		return dbc.Dsn
	}
	return dbc.User + ":" + dbc.Pwd + "@tcp(" + dbc.Host + ":" + dbc.Port + ")/" +
		dbc.Name + dbc.ParamStr()
}

func (dbc DatabaseConfig) ParamStr() string {
	p := ""
	if dbc.Params == nil {
		dbc.Params = make(map[string]string)
	}
	if _, ok := dbc.Params["charset"]; !ok {
		dbc.Params["charset"] = "utf8mb4"
	}
	if len(dbc.Params) > 0 {
		p = "?"
		for k, v := range dbc.Params {
			p += k + "=" + v + "&"
		}
		p = p[:len(p)-1]
	}
	return p
}
