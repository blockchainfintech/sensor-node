package count

import (
	"time"
	"math"
)

func GetCount(t time.Time, maxNum float64, meanTime float64) (int, error) {
	// TODO actually count the number of ack's in area
	//if t == nil { t := time.Now() }
	//x := float64(t.Hour() + t.Minute() / 60 + t.Second() / 3600)
	x := float64(t.Second()) / 6.0 // big boy fast for demo

	return int(normalDistrib(meanTime, 8.0, x) * maxNum), nil
}

func normalDistrib(mean float64, stddev float64, x float64) float64 {
	// Calculate the value between 0 and ~1 for std distribution with specified vals
	// mean and x are in hours between 0 and 24
	return 20*( 1/math.Sqrt( 2 * math.Pi * math.Pow(stddev, 2) ) )*math.Exp( -math.Pow((x-mean), 2)/(2 * math.Pow(stddev, 2)) )
}