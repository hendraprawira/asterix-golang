package utils

import (
	"asterix-golang/app/models"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"math"
	"strconv"
	"strings"

	"github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
)

func AsterixGeoJSONParse(data []byte) (datas []byte) {

	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)

	//dummy lat lon from ownunit/ship
	geoCrdRefStartAZ := models.OwnUnit{
		Lat: -9.2893,
		Lon: 106.4583,
	}

	geoCrdRefEndtAZ := models.OwnUnit{
		Lat: -9.2893,
		Lon: 106.4583,
	}
	cellDur := binary.BigEndian.Uint32(data[20:24])
	startRG := binary.BigEndian.Uint32(data[16:20])

	value7 := binary.BigEndian.Uint16(data[12:14])
	startAz := (float64(value7) / math.Pow(2, 16)) * float64(360)

	value8 := binary.BigEndian.Uint16(data[14:16])
	endAz := (float64(value8) / math.Pow(2, 16)) * float64(360)

	var distanceCellStart float64 = ((float64(cellDur) * math.Pow(10, -15)) * float64(startRG+1-1) * (299792458 / 2)) // 9.765624948527275
	getLatCrdRefStartAZ, getLonCrdRefStartAZ := geo1.At(geoCrdRefStartAZ.Lat, geoCrdRefStartAZ.Lon, distanceCellStart, startAz)
	getLatCrdRefEndAZ, getLonCrdRefEndAZ := geo1.At(geoCrdRefEndtAZ.Lat, geoCrdRefEndtAZ.Lon, distanceCellStart, endAz)

	// latAwalAZ, lonAwalAZ := CalculateNewCoordinates(geoCrdRefStartAZ.Lat, geoCrdRefStartAZ.Lon, startAz, cellStartAZ)
	// latAkhirAZ, lonAkhirAZ := CalculateNewCoordinates(geoCrdRefEndtAZ.Lat, geoCrdRefEndtAZ.Lon, endAz, cellEndAZ)

	res := data[25:26]
	calc := int(res[0]) - 1
	calca := math.Pow(2, float64(calc))

	value12 := binary.BigEndian.Uint16(data[26:28])
	hexString1 := hex.EncodeToString(data[28:31])
	value13, _ := strconv.ParseInt(strings.ReplaceAll(hexString1, " ", ""), 16, 64)
	value13a := ((int64(value13) * int64(calca)) / int64(res[0])) / 2

	hexString := hex.EncodeToString(data[32 : value13a+32])
	result := strings.Split(hexString, "")
	value14 := strings.ToUpper((strings.ReplaceAll(strings.Join(result, " "), " ", "")))
	vidioBlockArr := strings.Split(value14, "")

	C240 := models.Cat240s{
		I041: models.I041{
			StartAz: float64(startAz),
			EndAz:   float64(endAz),
			StartRg: startRG,
			CellDur: cellDur,
		},
		I048: models.I048{
			Res: uint32(res[0]),
		},
		I049: models.I049{
			Nbvb:    value12,
			NbCells: value13,
		},
	}
	// value6 := binary.BigEndian.Uint32(data[8:12])
	resolusi := GetRes(int(C240.I048.Res))
	substringStart := 0
	substringEnd := resolusi
	cell := 0
	geoJson := models.FeatureCollection{}

	for i := 0; i < (len(vidioBlockArr) / resolusi); i++ {
		opac, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(vidioBlockArr[substringStart:substringEnd], " "), " ", ""), 16, 64)
		opacs := (float64(opac) * 255) / 100 / 100

		if opacs > 0.9 {

			geoCrdRefStartAZ := models.OwnUnit{
				Lat: getLatCrdRefStartAZ,
				Lon: getLonCrdRefStartAZ,
			}

			geoCrdRefEndtAZ := models.OwnUnit{
				Lat: getLatCrdRefEndAZ,
				Lon: getLonCrdRefEndAZ,
			}

			var StartCell float64 = (float64(C240.I041.CellDur)) * (math.Pow(10, -15)) * float64(i+1-1) * (299792458 / 2) //    Distance Meter from ownUnit
			latPoint1, lonPoint1 := geo1.At(geoCrdRefStartAZ.Lat, geoCrdRefStartAZ.Lon, StartCell, startAz)
			latPoint4, lonPoint4 := geo1.At(geoCrdRefEndtAZ.Lat, geoCrdRefEndtAZ.Lon, StartCell, endAz)

			nextPointStartAZ := models.OwnUnit{
				Lat: latPoint1,
				Lon: lonPoint1,
			}

			nextPointEndAZ := models.OwnUnit{
				Lat: latPoint4,
				Lon: lonPoint4,
			}
			var ranges float64 = (float64(C240.I041.CellDur)) * (math.Pow(10, -15)) * float64(0+1-1) * (299792458 / 2) //    Distance Meter from ownUnit
			latPoint2, lonPoint2 := geo1.At(nextPointStartAZ.Lat, nextPointStartAZ.Lon, ranges, startAz)
			latPoint3, lonPoint3 := geo1.At(nextPointEndAZ.Lat, nextPointEndAZ.Lon, ranges, endAz)

			cellPoint1 := []float64{lonPoint1, latPoint1}
			cellPoint2 := []float64{lonPoint2, latPoint2}
			cellPoint3 := []float64{lonPoint3, latPoint3}
			cellPoint4 := []float64{lonPoint4, latPoint4}

			polygonCell := [][][]float64{{cellPoint1, cellPoint2, cellPoint3, cellPoint4}}

			geoJsonGeometry := models.Geometry{
				Coordinates: polygonCell,
				Type:        "Polygon",
			}

			geoJsonProperties := models.Properties{Opacity: opacs,
				EndAz: C240.I041.EndAz,
			}
			geoJsonFeature := models.Feature{
				Type:       "Feature",
				Geometry:   geoJsonGeometry,
				Properties: geoJsonProperties,
			}

			geoJson.Features = append(geoJson.Features, &geoJsonFeature)

		}
		substringStart = substringEnd
		substringEnd = substringEnd + resolusi
		cell = cell + 1

	}
	// if len(geoJson.Features) > 0 {
	geoJson.Type = "FeatureCollection"
	jsonData, _ := json.Marshal(geoJson)
	return jsonData
	// 	// SendWebSocketMessage(jsonData)
	// }
	// return nil
	// 	// geoJson.Type = "FeatureCollection"
	// 	// jsonData, _ := json.Marshal(geoJson)
	// 	// SendWebSocketMessage(jsonData)

}

// func AsterixGeoJSONParse(data []byte) (datas []byte) {

// 	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)

// 	//dummy lat lon from ownunit/ship
// 	geoCrdRefStartAZ := models.OwnUnit{
// 		Lat: -9.2893,
// 		Lon: 106.4583,
// 	}

// 	geoCrdRefEndtAZ := models.OwnUnit{
// 		Lat: -9.2893,
// 		Lon: 106.4583,
// 	}
// 	cellDur := binary.BigEndian.Uint32(data[20:24])
// 	startRG := binary.BigEndian.Uint32(data[16:20])

// 	value7 := binary.BigEndian.Uint16(data[12:14])
// 	startAz := (float64(value7) / math.Pow(2, 16)) * float64(360)

// 	value8 := binary.BigEndian.Uint16(data[14:16])
// 	endAz := (float64(value8) / math.Pow(2, 16)) * float64(360)

// 	var cellStartAZ float64 = (float64(cellDur) * math.Pow(10, -15)) * float64(startRG+1-1) * (299792458 / 2)           //  range awal cell dari own unit terhadap start azimuth
// 	var cellEndAZ float64 = (float64(cellDur) * math.Pow(10, -15)) * float64(startRG+1-1) * (299792458 / 2)             //  range awal cell dari own unit terhadap end azimuth
// 	var ranges float64 = ((float64(cellDur)*math.Pow(10, -15))*float64(startRG+2-1)*(299792458/2) - cellStartAZ) / 1000 // 9.765624948527275

// 	latAwalAZ, lonAwalAZ := geo1.At(geoCrdRefStartAZ.Lat, geoCrdRefStartAZ.Lon, cellStartAZ, startAz)
// 	latAkhirAZ, lonAkhirAZ := geo1.At(geoCrdRefEndtAZ.Lat, geoCrdRefEndtAZ.Lon, cellEndAZ, endAz)

// 	// latAwalAZ, lonAwalAZ := CalculateNewCoordinates(geoCrdRefStartAZ.Lat, geoCrdRefStartAZ.Lon, startAz, cellStartAZ)
// 	// latAkhirAZ, lonAkhirAZ := CalculateNewCoordinates(geoCrdRefEndtAZ.Lat, geoCrdRefEndtAZ.Lon, endAz, cellEndAZ)

// 	res := data[25:26]
// 	calc := int(res[0]) - 1
// 	calca := math.Pow(2, float64(calc))

// 	value12 := binary.BigEndian.Uint16(data[26:28])
// 	hexString1 := hex.EncodeToString(data[28:31])
// 	value13, _ := strconv.ParseInt(strings.ReplaceAll(hexString1, " ", ""), 16, 64)
// 	value13a := ((int64(value13) * int64(calca)) / int64(res[0])) / 2

// 	hexString := hex.EncodeToString(data[32 : value13a+32])
// 	result := strings.Split(hexString, "")
// 	value14 := strings.ToUpper((strings.ReplaceAll(strings.Join(result, " "), " ", "")))

// 	dataStruct := models.Cat240s{
// 		I041: models.I041{
// 			StartAz: float64(startAz),
// 			EndAz:   float64(endAz),
// 			StartRg: startRG,
// 			CellDur: cellDur,
// 		},
// 		I048: models.I048{
// 			Res: uint32(res[0]),
// 		},
// 		I049: models.I049{
// 			Nbvb:    value12,
// 			NbCells: value13,
// 		},
// 	}
// 	value6 := binary.BigEndian.Uint32(data[8:12])
// 	// fmt.Println(value6)

// 	genGeoJson := models.GenGeoJson{
// 		C240:       dataStruct,
// 		LatAwalAz:  latAwalAZ,
// 		LonAwalAz:  lonAwalAZ,
// 		LatAkhirAz: latAkhirAZ,
// 		LonAkhirAz: lonAkhirAZ,
// 		Ranges:     ranges,
// 		VideoBlock: value14,
// 		MsgIndex:   int64(value6),
// 	}
// 	datas = GenerateGeoJSON(genGeoJson)
// 	return datas

// }

// func GenerateGeoJSON(genGeoJson models.GenGeoJson) (data []byte) {
// 	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
// 	latAwalAZ1 := genGeoJson.LatAwalAz   // -6.9496124915037
// 	lonAwalAZ1 := genGeoJson.LonAwalAz   // 107.61957049369812
// 	latAkhirAZ1 := genGeoJson.LatAkhirAz // -6.9496124915037
// 	lonAkhirAZ1 := genGeoJson.LonAkhirAz // 107.61957049369812
// 	substringStart := 0
// 	substringEnd := 0
// 	resolusi := GetRes(int(genGeoJson.C240.I048.Res))
// 	vidioBlock := genGeoJson.VideoBlock // 2
// 	vidioBlockArr := strings.Split(vidioBlock, "")
// 	geoJson := models.FeatureCollection{}

// 	for i := 0; i < int(genGeoJson.C240.I049.NbCells); i++ {
// 		substringEnd = substringEnd + resolusi
// 		opac, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(vidioBlockArr[substringStart:substringEnd], " "), " ", ""), 16, 64)
// 		opacs := (float64(opac) * 255) / 10000
// 		substringStart = substringEnd

// 		geoCoordinateStart := models.OwnUnit{
// 			Lat: latAwalAZ1,
// 			Lon: lonAwalAZ1,
// 		}
// 		geoCoordinateEnd := models.OwnUnit{
// 			Lat: latAkhirAZ1,
// 			Lon: lonAkhirAZ1,
// 		}

// 		latStart, lonStart := geo1.At(geoCoordinateStart.Lat, geoCoordinateStart.Lon, genGeoJson.Ranges, genGeoJson.C240.I041.StartAz)
// 		latEndAz, lonEndAz := geo1.At(geoCoordinateEnd.Lat, geoCoordinateEnd.Lon, genGeoJson.Ranges, genGeoJson.C240.I041.EndAz)

// 		// latStart, lonStart := CalculateNewCoordinates(geoCoordinateStart.Lat, geoCoordinateStart.Lon, genGeoJson.C240.I041.StartAz, genGeoJson.Ranges)
// 		// latEndAz, lonEndAz := CalculateNewCoordinates(geoCoordinateEnd.Lat, geoCoordinateEnd.Lon, genGeoJson.C240.I041.EndAz, genGeoJson.Ranges)

// 		log.Print(latStart, lonStart)

// 		if opacs >= 0.8 {
// 			cellPoint1 := []float64{lonAwalAZ1, latAwalAZ1}
// 			cellPoint2 := []float64{lonStart, latStart}
// 			cellPoint3 := []float64{lonEndAz, latEndAz}
// 			cellPoint4 := []float64{lonAkhirAZ1, latAkhirAZ1}

// 			polygonCell := [][][]float64{{cellPoint1, cellPoint2, cellPoint3, cellPoint4}}
// 			geoJsonGeometry := models.Geometry{
// 				Coordinates: polygonCell,
// 				Type:        "Polygon",
// 			}

// 			geoJsonProperties := models.Properties{Opacity: opacs,
// 				EndAz: genGeoJson.C240.I041.EndAz,
// 			}
// 			geoJsonFeature := models.Feature{
// 				Type:       "Feature",
// 				Geometry:   geoJsonGeometry,
// 				Properties: geoJsonProperties,
// 			}

// 			geoJson.Features = append(geoJson.Features, &geoJsonFeature)
// 		}
// 		lonAwalAZ1 = lonStart
// 		latAwalAZ1 = latStart
// 		lonAkhirAZ1 = lonEndAz
// 		latAkhirAZ1 = latEndAz

// 	}
// 	if len(geoJson.Features) > 0 {
// 		geoJson.Type = "FeatureCollection"
// 		geoJson.MsgIndex = genGeoJson.MsgIndex
// 		jsonData, _ := json.Marshal(geoJson)
// 		return jsonData
// 		// SendWebSocketMessage(jsonData)
// 	}
// 	return nil
// 	// geoJson.Type = "FeatureCollection"
// 	// jsonData, _ := json.Marshal(geoJson)
// 	// SendWebSocketMessage(jsonData)
// }

// Get Vidio Resolution
func GetRes(res int) int {
	resolusi := 0
	if res == 3 {
		resolusi = 1
	} else if res == 4 {
		resolusi = 2
	} else if res == 5 {
		resolusi = 4
	} else if res == 6 {
		resolusi = 8
	}
	return resolusi
}
