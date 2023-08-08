package models

type Cat240s struct {
	I041 I041 `json:"i041"`
	I048 I048 `json:"i048"`
	I049 I049 `json:"i049"`
	I050 I050 `json:"i050,omitempty"`
	I051 I051 `json:"i051,omitempty"`
	I052 I052 `json:"i052,omitempty"`
}
type I041 struct {
	StartAz float64 `json:"START_AZ"`
	EndAz   float64 `json:"END_AZ"`
	StartRg uint32  `json:"START_RG"`
	CellDur uint32  `json:"CELL_DUR"`
}

type I048 struct {
	Res uint32 `json:"RES"`
}

type I049 struct {
	Nbvb    uint16 `json:"NB_VB"`
	NbCells int64  `json:"NB_CELLS"`
}

type I050 struct {
	Video string `json:"VIDEO,omitempty"`
}

type I051 struct {
	Video string `json:"VIDEO,omitempty"`
}

type I052 struct {
	Video string `json:"VIDEO,omitempty"`
}
