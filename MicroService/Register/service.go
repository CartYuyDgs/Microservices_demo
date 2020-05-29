package Register

type Service struct {
	Name  string  `json:"name"`
	Nodes []*Node `json:"nodes"`
}

type Node struct {
	Id   string `json:"id"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}
