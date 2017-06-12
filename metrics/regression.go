package metrics

import (
	"log"
	"strconv"
)

type Residuals struct {
	SS_Res float64
	Sum    float64
	Count  int64
	ys     []float64
}

func (ss *Residuals) RSquared() float64 {
	average := ss.Sum / float64(ss.Count)
	var ssTot float64
	for _, y := range ss.ys {
		tot := y - average
		ssTot += tot * tot
	}
	return 1.0 - (ss.SS_Res / ssTot)
}

func (ss *Residuals) MSE() float64 {
	return ss.SS_Res / float64(ss.Count)
}

func (ss *Residuals) Update(yText, yHatText, yWeight string) {
	y, err := strconv.ParseFloat(yText, 64)
	if err != nil {
		log.Fatal("Error parsing float: ", err)
	}

	yHat, err := strconv.ParseFloat(yHatText, 64)
	if err != nil {
		log.Fatal("Error parsing float: ", err)
	}

	ss.Count++
	ss.Sum += y
	e := y - yHat
	ss.SS_Res += e * e
	ss.ys = append(ss.ys, y)
}
