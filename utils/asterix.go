package utils

import (
	"asterix-golang/models"
	"encoding/binary"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
	"time"
)

// func AsterixParse(data []string) models.Cat240 {
// 	lastIndex := len(data) - 1
// 	lastIndexV := lastIndex - 2
// 	a := data[0:1]
// 	value1, _ := strconv.ParseInt(strings.Join(a, " "), 16, 64)

// 	b := data[1:3]
// 	value2, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(b, " "), " ", ""), 16, 64)

// 	c := data[5:6]
// 	value3, _ := strconv.ParseInt(strings.Join(c, " "), 16, 64)

// 	d := data[6:7]
// 	value4, _ := strconv.ParseInt(strings.Join(d, " "), 16, 64)

// 	e := data[7:8]
// 	value5, _ := strconv.ParseInt(strings.Join(e, " "), 16, 64)

// 	f := data[8:12]
// 	value6, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(f, " "), " ", ""), 16, 64)

// 	g := data[12:14]
// 	value7, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(g, " "), " ", ""), 16, 64)
// 	value7a := (float64(value7) / math.Pow(2, 16)) * float64(360)

// 	h := data[14:16]
// 	value8, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(h, " "), " ", ""), 16, 64)
// 	value8a := (float64(value8) / math.Pow(2, 16)) * float64(360)

// 	i := data[16:20]
// 	value9, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(i, " "), " ", ""), 16, 64)

// 	j := data[20:24]
// 	value10, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(j, " "), " ", ""), 16, 64)

// 	k := data[25:26]
// 	value11, _ := strconv.ParseInt(strings.Join(k, " "), 16, 64)

// 	calc := value11 - 1
// 	calca := math.Pow(2, float64(calc))

// 	l := data[26:28]
// 	value12, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(l, " "), " ", ""), 16, 64)

// 	m := data[28:31]
// 	value13, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(m, " "), " ", ""), 16, 64)

// 	value13a := ((value13 * int64(calca)) / value11) / 2

// 	n := data[32 : value13a+32]
// 	value14 := strings.ToUpper((strings.ReplaceAll(strings.Join(n, " "), " ", "")))

// 	o := data[lastIndexV : lastIndex+1]
// 	value15, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(o, " "), " ", ""), 16, 64)
// 	tod := (float64(value15) / 128)

// 	dataStruct := models.Cat240{}

// 	if value2 <= 1020 {
// 		dataStruct = models.Cat240{
// 			Category: value1,
// 			Length:   value2,
// 			Time:     time.Now().Unix(),
// 			I010: models.I010{
// 				SAC: value3,
// 				SIC: value4,
// 			},
// 			I000: models.I000{
// 				TYP: value5,
// 			},
// 			I020: models.I020{
// 				MsgIndex: value6,
// 			},
// 			I041: models.I041{
// 				StartAz: float64(value7a),
// 				EndAz:   float64(value8a),
// 				StartRg: value9,
// 				CellDur: value10,
// 			},
// 			I048: models.I048{
// 				Res: value11,
// 			},
// 			I049: models.I049{
// 				Nbvb:    value12,
// 				NbCells: value13,
// 			},
// 			I050: &models.I050{
// 				Video: value14,
// 			},
// 			I040: models.I040{
// 				ToD: tod,
// 			},
// 		}
// 	} else if value2 > 1020 && value2 <= 16320 {
// 		dataStruct = models.Cat240{
// 			Category: value1,
// 			Length:   value2,
// 			Time:     time.Now().Unix(),
// 			I010: models.I010{
// 				SAC: value3,
// 				SIC: value4,
// 			},
// 			I000: models.I000{
// 				TYP: value5,
// 			},
// 			I020: models.I020{
// 				MsgIndex: value6,
// 			},
// 			I041: models.I041{
// 				StartAz: float64(value7a),
// 				EndAz:   float64(value8a),
// 				StartRg: value9,
// 				CellDur: value10,
// 			},
// 			I048: models.I048{
// 				Res: value11,
// 			},
// 			I049: models.I049{
// 				Nbvb:    value12,
// 				NbCells: value13,
// 			},
// 			I051: &models.I051{
// 				Video: value14,
// 			},
// 			I040: models.I040{
// 				ToD: tod,
// 			},
// 		}
// 	} else if value2 > 16320 && value2 <= 65024 {
// 		dataStruct = models.Cat240{
// 			Category: value1,
// 			Length:   value2,
// 			Time:     time.Now().Unix(),
// 			I010: models.I010{
// 				SAC: value3,
// 				SIC: value4,
// 			},
// 			I000: models.I000{
// 				TYP: value5,
// 			},
// 			I020: models.I020{
// 				MsgIndex: value6,
// 			},
// 			I041: models.I041{
// 				StartAz: float64(value7a),
// 				EndAz:   float64(value8a),
// 				StartRg: value9,
// 				CellDur: value10,
// 			},
// 			I048: models.I048{
// 				Res: value11,
// 			},
// 			I049: models.I049{
// 				Nbvb:    value12,
// 				NbCells: value13,
// 			},
// 			I052: &models.I052{
// 				Video: value14,
// 			},
// 			I040: models.I040{
// 				ToD: tod,
// 			},
// 		}
// 	}
// 	return dataStruct
// }

func AsterixParse(data []byte) models.Cat240 {
	lastIndex := len(data) - 1
	lastIndexV := lastIndex - 2
	value1 := data[0:1]
	value2 := binary.BigEndian.Uint16(data[1:3])
	value3 := data[5:6]
	value4 := data[6:7]
	value5 := data[7:8]
	value6 := binary.BigEndian.Uint32(data[8:12])
	value7 := binary.BigEndian.Uint16(data[12:14])
	value7a := (float64(value7) / math.Pow(2, 16)) * float64(360)

	value8 := binary.BigEndian.Uint16(data[14:16])
	value8a := (float64(value8) / math.Pow(2, 16)) * float64(360)

	value9 := binary.BigEndian.Uint32(data[16:20])
	value10 := binary.BigEndian.Uint32(data[20:24])
	value11 := data[25:26]
	calc := int(value11[0]) - 1
	calca := math.Pow(2, float64(calc))

	value12 := binary.BigEndian.Uint16(data[26:28])
	hexString1 := hex.EncodeToString(data[28:31])
	value13, _ := strconv.ParseInt(strings.ReplaceAll(hexString1, " ", ""), 16, 64)
	value13a := ((int64(value13) * int64(calca)) / int64(value11[0])) / 2

	hexString := hex.EncodeToString(data[32 : value13a+32])
	result := strings.Split(hexString, "")
	value14 := strings.ToUpper((strings.ReplaceAll(strings.Join(result, " "), " ", "")))

	o := binary.BigEndian.Uint16(data[lastIndexV : lastIndex+1])
	tod := (float64(o) / 128)

	dataStruct := models.Cat240{}

	if value2 <= 1020 {
		dataStruct = models.Cat240{
			Category: int64(value1[0]),
			Length:   value2,
			Time:     time.Now().Unix(),
			I010: models.I010{
				SAC: uint32(value3[0]),
				SIC: uint32(value4[0]),
			},
			I000: models.I000{
				TYP: uint32(value5[0]),
			},
			I020: models.I020{
				MsgIndex: value6,
			},
			I041: models.I041{
				StartAz: float64(value7a),
				EndAz:   float64(value8a),
				StartRg: value9,
				CellDur: value10,
			},
			I048: models.I048{
				Res: uint32(value11[0]),
			},
			I049: models.I049{
				Nbvb:    value12,
				NbCells: value13,
			},
			I050: &models.I050{
				Video: value14,
			},
			I040: models.I040{
				ToD: tod,
			},
		}
	} else if value2 > 1020 && value2 <= 16320 {
		dataStruct = models.Cat240{
			Category: int64(value1[0]),
			Length:   value2,
			Time:     time.Now().Unix(),
			I010: models.I010{
				SAC: uint32(value3[0]),
				SIC: uint32(value4[0]),
			},
			I000: models.I000{
				TYP: uint32(value5[0]),
			},
			I020: models.I020{
				MsgIndex: value6,
			},
			I041: models.I041{
				StartAz: float64(value7a),
				EndAz:   float64(value8a),
				StartRg: value9,
				CellDur: value10,
			},
			I048: models.I048{
				Res: uint32(value11[0]),
			},
			I049: models.I049{
				Nbvb:    value12,
				NbCells: value13,
			},
			I051: &models.I051{
				Video: value14,
			},
			I040: models.I040{
				ToD: tod,
			},
		}
	} else if value2 > 16320 && value2 <= 65024 {
		dataStruct = models.Cat240{
			Category: int64(value1[0]),
			Length:   value2,
			Time:     time.Now().Unix(),
			I010: models.I010{
				SAC: uint32(value3[0]),
				SIC: uint32(value4[0]),
			},
			I000: models.I000{
				TYP: uint32(value5[0]),
			},
			I020: models.I020{
				MsgIndex: value6,
			},
			I041: models.I041{
				StartAz: float64(value7a),
				EndAz:   float64(value8a),
				StartRg: value9,
				CellDur: value10,
			},
			I048: models.I048{
				Res: uint32(value11[0]),
			},
			I049: models.I049{
				Nbvb:    value12,
				NbCells: value13,
			},
			I052: &models.I052{
				Video: value14,
			},
			I040: models.I040{
				ToD: tod,
			},
		}
	}
	return dataStruct
}
