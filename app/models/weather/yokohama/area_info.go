package yokohama

type YokohamaAreaInfo struct {
	Area struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"area"`
	Pops *[]string `json:"pops"`
}

func (yai *YokohamaAreaInfo) GetCode() string {
	return yai.Area.Code
}

func (yai *YokohamaAreaInfo) GetPops() *[]string {
	return yai.Pops
}
