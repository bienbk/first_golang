package model

type AddResult struct {
	Calc int32 `json:"total"`
}

type SubResult struct {
	Calc int32 `json:"sub"`
}

type PostDivParam struct {
	Param  float64 `json:"a"`
	Param2 float64 `json:"b"`
}
type DivResult struct {
	Calc float64 `json:"div"`
}

type PostMultParam struct {
	Param  float32 `json:"c"`
	Param2 float32 `json:"d"`
}
type MultResult struct {
	Calc float32 `json:"Multiplication"`
}
