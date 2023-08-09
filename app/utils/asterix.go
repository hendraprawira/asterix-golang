package utils

import (
	"asterix-golang/app/models"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
)

func AsterixGeoJSONParse(data []byte) (datas []byte) {
	opacity, _ := strconv.ParseFloat(os.Getenv("OPACITY"), 32)
	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)

	geoCrdRefStartAZ := models.OwnUnit{
		Lat: 47.39885782790412,
		Lon: -122.45102114364046,
	}

	geoCrdRefEndtAZ := models.OwnUnit{
		Lat: 47.39885782790412,
		Lon: -122.45102114364046,
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
	resBit := int(res[0]) - 1
	resBits := math.Pow(2, float64(resBit))

	nbvb := binary.BigEndian.Uint16(data[26:28])
	hexStringVideo := hex.EncodeToString(data[28:31])
	videoBlock, _ := strconv.ParseInt(strings.ReplaceAll(hexStringVideo, " ", ""), 16, 64)
	videoBlockLen := ((int64(videoBlock) * int64(resBits)) / int64(res[0])) / 2

	hexString := hex.EncodeToString(data[32 : videoBlockLen+32])
	result := strings.Split(hexString, "")
	value14 := strings.ToUpper((strings.ReplaceAll(strings.Join(result, " "), " ", "")))
	videoBlockArr := strings.Split(value14, "")

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
			Nbvb:    nbvb,
			NbCells: videoBlock,
		},
	}
	resolusi := GetRes(int(C240.I048.Res))
	substringStart := 0
	substringEnd := resolusi
	cell := 0
	geoJson := models.FeatureCollection{}
	geoJson.EndAz = C240.I041.EndAz
	geoJson.StartAz = C240.I041.StartAz

	// var ranges1 float64 = (float64(C240.I041.CellDur)) * (math.Pow(10, -15)) * float64(0+2-1) * (299792458 / 2) //    Distance Meter from ownUnit
	// radius2 := distanceCellStart + (ranges1 * float64(C240.I049.NbCells))

	// var ranges1 float64 = (float64(C240.I041.CellDur)) * (math.Pow(10, -15)) * float64(0+2-1) * (299792458 / 2) //    Distance Meter from ownUnit
	radius2 := 20000.0

	for i := 0; i < (len(videoBlockArr) / resolusi); i++ {
		opac, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(videoBlockArr[substringStart:substringEnd], " "), " ", ""), 16, 64)
		opacs := (float64(opac) / math.Pow(2, float64(resolusi)))

		if opacs > opacity {

			geoCrdRefStartAZ := models.OwnUnit{
				Lat: getLatCrdRefStartAZ,
				Lon: getLonCrdRefStartAZ,
			}

			geoCrdRefEndtAZ := models.OwnUnit{
				Lat: getLatCrdRefEndAZ,
				Lon: getLonCrdRefEndAZ,
			}

			var startCell float64 = (float64(C240.I041.CellDur)) * (math.Pow(10, -15)) * float64(i+1-1) * (299792458 / 2) //    Distance Meter from ownUnit
			latPoint1, lonPoint1 := geo1.At(geoCrdRefStartAZ.Lat, geoCrdRefStartAZ.Lon, startCell, startAz)
			latPoint4, lonPoint4 := geo1.At(geoCrdRefEndtAZ.Lat, geoCrdRefEndtAZ.Lon, startCell, endAz)

			nextPointStartAZ := models.OwnUnit{
				Lat: latPoint1,
				Lon: lonPoint1,
			}

			nextPointEndAZ := models.OwnUnit{
				Lat: latPoint4,
				Lon: lonPoint4,
			}
			var ranges float64 = (float64(C240.I041.CellDur)) * (math.Pow(10, -15)) * float64(0+2-1) * (299792458 / 2) //    Distance Meter from ownUnit
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
	geoJson.Radius = radius2
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
