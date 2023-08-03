package utils

// import (
// 	"fmt"
// 	"image"
// 	"image/color"
// 	"image/draw"
// 	"image/png"
// 	"os"

// 	"github.com/paulmach/orb"
// 	"github.com/paulmach/orb/geojson"
// )

// func Raster() {
// 	// Replace "path/to/input.geojson" with the path to your GeoJSON file.
// 	geoJSONFilePath := "path/to/input.geojson"

// 	// Parse the GeoJSON file.
// 	f, err := os.Open(geoJSONFilePath)
// 	if err != nil {
// 		fmt.Println("Error opening GeoJSON file:", err)
// 		return
// 	}
// 	defer f.Close()

// 	// Decode the GeoJSON data into an orb.Geometry.
// 	geometry, err := geojson.NewBBox()
// 	if err != nil {
// 		fmt.Println("Error decoding GeoJSON:", err)
// 		return
// 	}

// 	// Set up the bounding box and image size.
// 	bounds := geometry.Bound()
// 	width := 800
// 	height := 600

// 	// Create an image representation.
// 	img := image.NewRGBA(image.Rect(0, 0, width, height))
// 	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

// 	// Convert coordinates to image coordinates.
// 	// This function maps geo coordinates to the image coordinates.
// 	// You may need to adjust this based on your specific use case.
// 	convertToImageCoordinates := func(p orb.Point) image.Point {
// 		x := int((p[0] - bounds.Min[0]) / (bounds.Max[0] - bounds.Min[0]) * float64(width))
// 		y := int((p[1] - bounds.Min[1]) / (bounds.Max[1] - bounds.Min[1]) * float64(height))
// 		return image.Point{x, y}
// 	}

// 	// Draw the GeoJSON geometry on the image.
// 	switch geom := geometry.(type) {
// 	case orb.Point:
// 		img.Set(convertToImageCoordinates(geom), color.Black)
// 	case orb.MultiPoint:
// 		for _, p := range geom {
// 			img.Set(convertToImageCoordinates(p), color.Black)
// 		}
// 	case orb.LineString:
// 		for i := 0; i < len(geom)-1; i++ {
// 			start := convertToImageCoordinates(geom[i])
// 			end := convertToImageCoordinates(geom[i+1])
// 			drawLine(img, start, end, color.Black)
// 		}
// 	case orb.MultiLineString:
// 		for _, line := range geom {
// 			for i := 0; i < len(line)-1; i++ {
// 				start := convertToImageCoordinates(line[i])
// 				end := convertToImageCoordinates(line[i+1])
// 				drawLine(img, start, end, color.Black)
// 			}
// 		}
// 	case orb.Polygon:
// 		for _, ring := range geom {
// 			for i := 0; i < len(ring)-1; i++ {
// 				start := convertToImageCoordinates(ring[i])
// 				end := convertToImageCoordinates(ring[i+1])
// 				drawLine(img, start, end, color.Black)
// 			}
// 		}
// 	case orb.MultiPolygon:
// 		for _, polygon := range geom {
// 			for _, ring := range polygon {
// 				for i := 0; i < len(ring)-1; i++ {
// 					start := convertToImageCoordinates(ring[i])
// 					end := convertToImageCoordinates(ring[i+1])
// 					drawLine(img, start, end, color.Black)
// 				}
// 			}
// 		}
// 	}

// 	// Save the image to a PNG file.
// 	outputPNGFilePath := "output.png"
// 	outputFile, err := os.Create(outputPNGFilePath)
// 	if err != nil {
// 		fmt.Println("Error creating PNG file:", err)
// 		return
// 	}
// 	defer outputFile.Close()

// 	err = png.Encode(outputFile, img)
// 	if err != nil {
// 		fmt.Println("Error encoding PNG:", err)
// 		return
// 	}

// 	fmt.Println("Image saved to:", outputPNGFilePath)
// }

// // drawLine draws a line between two points on the image.
// func drawLine(img *image.RGBA, start, end image.Point, col color.Color) {
// 	dx := end.X - start.X
// 	dy := end.Y - start.Y
// 	steps := 0

// 	if abs(dx) > abs(dy) {
// 		steps = abs(dx)
// 	} else {
// 		steps = abs(dy)
// 	}

// 	xIncrement := float64(dx) / float64(steps)
// 	yIncrement := float64(dy) / float64(steps)

// 	x := float64(start.X)
// 	y := float64(start.Y)

// 	for i := 0; i < steps; i++ {
// 		x += xIncrement
// 		y += yIncrement
// 		img.Set(int(x), int(y), col)
// 	}
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }
