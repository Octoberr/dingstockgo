package redis

type RdbConf struct {
	Address string `json:"address"`
	Port    string `json:"port"`
	DB      int    `json:"db"`
	PassWD  string `json:"passwd"`
}
