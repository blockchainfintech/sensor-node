package count

import (
	"time"
	"math"
)

const maxNum = 5000

func GetCount(t time.Time) (int, error) {
	// TODO actually count the number of ack's in area
	//if t == nil { t := time.Now() }
	return int(normalDistrib(9.0, float64(t.Hour() + t.Minute() / 60) ) * maxNum), nil
}

func normalDistrib(mean float64, x float64) float64 {
	// Calculate the value between 0 and ~1 for std distribution with specified vals
	// mean and x are in hours between 0 and 24
	return 20*( 1/math.Sqrt( 2 * math.Pi * 64 ) )*math.Exp( -math.Pow((x-mean), 2)/(2 * 64) )
}