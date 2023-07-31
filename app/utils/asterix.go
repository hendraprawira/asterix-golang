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

func AsterixGeoJSONParse(data []byte) {

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

	var cellStartAZ float64 = (float64(cellDur) * math.Pow(10, -15)) * float64(startRG+1-1) * (299792458 / 2)  //  range awal cell dari own unit terhadap start azimuth
	var cellEndAZ float64 = (float64(cellDur) * math.Pow(10, -15)) * float64(startRG+1-1) * (299792458 / 2)    //  range awal cell dari own unit terhadap end azimuth
	var ranges float64 = (float64(cellDur)*math.Pow(10, -15))*float64(startRG+2-1)*(299792458/2) - cellStartAZ // 9.765624948527275

	latAwalAZ, lonAwalAZ := geo1.At(geoCrdRefStartAZ.Lat, geoCrdRefStartAZ.Lon, cellStartAZ, startAz)
	latAkhirAZ, lonAkhirAZ := geo1.At(geoCrdRefEndtAZ.Lat, geoCrdRefEndtAZ.Lon, cellEndAZ, endAz)

	len240 := binary.BigEndian.Uint16(data[1:3])
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

	dataStruct := models.Cat240s{
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

	if len240 <= 1020 {
		dataStruct.I050.Video = value14
	} else if len240 > 1020 && len240 <= 16320 {
		dataStruct.I051.Video = value14
	} else if len240 > 16320 && len240 <= 65024 {
		dataStruct.I052.Video = value14
	}

	genGeoJson := models.GenGeoJson{
		C240:       dataStruct,
		LatAwalAz:  latAwalAZ,
		LonAwalAz:  lonAwalAZ,
		LatAkhirAz: latAkhirAZ,
		LonAkhirAz: lonAkhirAZ,
		Ranges:     ranges,
	}
	generateGeoJSON(genGeoJson)
}

func generateGeoJSON(genGeoJson models.GenGeoJson) {
	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
	latAwalAZ1 := genGeoJson.LatAwalAz   // -6.9496124915037
	lonAwalAZ1 := genGeoJson.LonAwalAz   // 107.61957049369812
	latAkhirAZ1 := genGeoJson.LatAkhirAz // -6.9496124915037
	lonAkhirAZ1 := genGeoJson.LonAkhirAz // 107.61957049369812
	substringStart := 0
	substringEnd := 0
	resolusi := GetRes(int(genGeoJson.C240.I048.Res))
	vidioBlock := GetVideoBlock(genGeoJson.C240) // 2
	vidioBlockArr := strings.Split(vidioBlock, "")
	geoJson := models.FeatureCollection{}

	for i := 0; i < int(genGeoJson.C240.I049.NbCells); i++ {
		geoCoordinateStart := models.OwnUnit{
			Lat: latAwalAZ1,
			Lon: lonAwalAZ1,
		}
		geoCoordinateEnd := models.OwnUnit{
			Lat: latAkhirAZ1,
			Lon: lonAkhirAZ1,
		}
		latStart, lonStart := geo1.At(geoCoordinateStart.Lat, geoCoordinateStart.Lon, genGeoJson.Ranges, genGeoJson.C240.I041.StartAz)
		latEndAz, lonEndAz := geo1.At(geoCoordinateEnd.Lat, geoCoordinateEnd.Lon, genGeoJson.Ranges, genGeoJson.C240.I041.EndAz)

		cellPoint1 := []float64{lonAwalAZ1, latAwalAZ1}
		cellPoint2 := []float64{lonStart, latStart}
		cellPoint3 := []float64{lonEndAz, latEndAz}
		cellPoint4 := []float64{lonAkhirAZ1, latAkhirAZ1}

		polygonCell := [][][]float64{{cellPoint1, cellPoint2, cellPoint3, cellPoint4}}

		lonAwalAZ1 = lonStart
		latAwalAZ1 = latStart
		lonAkhirAZ1 = lonEndAz
		latAkhirAZ1 = latEndAz

		geoJsonGeometry := models.Geometry{
			Coordinates: polygonCell,
			Type:        "Polygon",
		}

		substringEnd = substringEnd + resolusi
		opac, _ := strconv.ParseInt(strings.ReplaceAll(strings.Join(vidioBlockArr[substringStart:substringEnd], " "), " ", ""), 16, 64)
		opacs := (float64(opac) * 255) / 10000

		geoJsonProperties := models.Properties{Opacity: opacs}
		geoJsonFeature := models.Feature{
			Type:       "Feature",
			Geometry:   geoJsonGeometry,
			Properties: geoJsonProperties,
		}
		substringStart = substringEnd
		geoJson.Features = append(geoJson.Features, &geoJsonFeature)

	}
	geoJson.Type = "FeatureCollection"
	jsonData, _ := json.Marshal(geoJson)
	SendWebSocketMessage(jsonData)
}

// Get Vidio Block
func GetVideoBlock(c240 models.Cat240s) string {
	var videoBlock string
	if c240.I050.Video != "" {
		videoBlock = c240.I050.Video
	} else if c240.I051.Video != "" {
		videoBlock = c240.I051.Video
	} else if c240.I052.Video != "" {
		videoBlock = c240.I052.Video
	}
	return videoBlock
}

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
