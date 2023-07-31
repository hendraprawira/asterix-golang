package models

type Cat240 struct {
	Category int64  `json:"category"`
	Length   uint16 `json:"len"`
	Time     int64  `json:"ts"`
	I010     I010   `json:"i010"`
	I000     I000   `json:"i000"`
	I020     I020   `json:"i020"`
	I041     I041   `json:"i041"`
	I048     I048   `json:"i048"`
	I049     I049   `json:"i049"`
	I050     *I050  `json:"i050,omitempty"`
	I051     *I051  `json:"i051,omitempty"`
	I052     *I052  `json:"i052,omitempty"`
	I040     I040   `json:"i040"`
}

type Cat240s struct {
	I041 I041 `json:"i041"`
	I048 I048 `json:"i048"`
	I049 I049 `json:"i049"`
	I050 I050 `json:"i050,omitempty"`
	I051 I051 `json:"i051,omitempty"`
	I052 I052 `json:"i052,omitempty"`
}

type I010 struct {
	SAC uint32 `json:"SAC"`
	SIC uint32 `json:"SIC"`
}

type I000 struct {
	TYP uint32 `json:"TYP"`
}

type I020 struct {
	MsgIndex uint32 `json:"MSG_INDEX"`
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

type I040 struct {
	ToD float64 `json:"ToD"`
}
