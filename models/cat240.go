// package models

// type Cat240 struct {
// 	Category int64 `json:"Category"`
// 	Length   int64 `json:"Length"`
// 	Time     int64 `json:"Time"`
// 	I010     I010  `json:"I010"`
// 	I000     I000  `json:"I000"`
// 	I020     I020  `json:"I020"`
// 	I041     I041  `json:"I041"`
// 	I048     I048  `json:"I048"`
// 	I049     I049  `json:"I049"`
// 	I050     *I050 `json:"I050,omitempty"`
// 	I051     *I051 `json:"I051,omitempty"`
// 	I052     *I052 `json:"I052,omitempty"`
// 	I040     I040  `json:"I040"`
// }

// type I010 struct {
// 	SAC int64 `json:"SAC"`
// 	SIC int64 `json:"SIC"`
// }

// type I000 struct {
// 	TYP int64 `json:"TYP"`
// }

// type I020 struct {
// 	MsgIndex int64 `json:"MSG_INDEX"`
// }

// type I041 struct {
// 	StartAz float64 `json:"START_AZ"`
// 	EndAz   float64 `json:"END_AZ"`
// 	StartRg int64   `json:"START_RG"`
// 	CellDur int64   `json:"CELL_DUR"`
// }

// type I048 struct {
// 	Res int64 `json:"RES"`
// }

// type I049 struct {
// 	Nbvb    int64 `json:"NB_VB"`
// 	NbCells int64 `json:"NB_CELLS"`
// }

// type I050 struct {
// 	Video string `json:"VIDEO,omitempty"`
// }

// type I051 struct {
// 	Video string `json:"VIDEO,omitempty"`
// }

// type I052 struct {
// 	Video string `json:"VIDEO,omitempty"`
// }

// type I040 struct {
// 	ToD float64 `json:"ToD"`
// }

package models

type Cat240 struct {
	Category int64  `json:"Category"`
	Length   uint16 `json:"Length"`
	Time     int64  `json:"Time"`
	I010     I010   `json:"I010"`
	I000     I000   `json:"I000"`
	I020     I020   `json:"I020"`
	I041     I041   `json:"I041"`
	I048     I048   `json:"I048"`
	I049     I049   `json:"I049"`
	I050     *I050  `json:"I050,omitempty"`
	I051     *I051  `json:"I051,omitempty"`
	I052     *I052  `json:"I052,omitempty"`
	I040     I040   `json:"I040"`
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
