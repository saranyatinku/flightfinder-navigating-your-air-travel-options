package geo

import "fmt"

// ConversionError indicates problem with coordinates conversion
type ConversionError struct {
	Msg string
}

func (e ConversionError) Error() string {
	return e.Msg
}

// ConvertDegMinSecHemToLongitude turns 4 component longitude into 1 component longitude
func ConvertDegMinSecHemToLongitude(deg, min, sec int, hem string) (Longitude, error) {
	d := float32(deg) + float32(min)/60.0 + float32(sec)/3600.0
	switch hem {
	case "E":
		return Longitude(d), nil
	case "W":
		return Longitude(-d), nil
	default:
		return Longitude(0), ConversionError{fmt.Sprintf("Invalid longitude hemisphere: %q", hem)}
	}
}

// ConvertDegMinSecHemToLatitude turns 4 component latitude into 1 component latitude
func ConvertDegMinSecHemToLatitude(deg, min, sec int, hem string) (Latitude, error) {
	d := float32(deg) + float32(min)/60.0 + float32(sec)/3600.0
	switch hem {
	case "N":
		return Latitude(d), nil
	case "S":
		return Latitude(-d), nil
	default:
		return Latitude(0), ConversionError{fmt.Sprintf("Invalid latitude hemisphere: %q", hem)}
	}
}
