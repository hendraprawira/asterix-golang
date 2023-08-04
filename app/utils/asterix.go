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
		Lat: -6.949612491503703,
		Lon: 107.61957049369812,
	}

	geoCrdRefEndtAZ := models.OwnUnit{
		Lat: -6.949612491503703,
		Lon: 107.61957049369812,
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
	resolusi := GetRes(int(C240.I048.Res))
	substringStart := 0
	substringEnd := resolusi
	cell := 0
	geoJson := models.FeatureCollection{}
	geoJson.EndAz = C240.I041.EndAz
	geoJson.StartAz = C240.I041.StartAz

	for i := 0; i < (len(vidioBlockArr) / resolusi); i++ {
		opac, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(vidioBlockArr[substringStart:substringEnd], " "), " ", ""), 16, 64)
		opacs := (float64(opac) * 255) / 100 / 100

		if opacs > 0.8 {

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
				EndAz:  C240.I041.EndAz,
				Radius: 20000.0,
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
	geoJson.Type = "FeatureCollection"
	jsonData, _ := json.Marshal(geoJson)
	return jsonData
}

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
