package knn

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/mariojose123/knngo/minmax"
)

type KnnPoints struct {
	X         float64       `json:"x"`
	Y         float64       `json:"y"`
	Class     int           `json:"class"`
	Train     bool          `json:"train"`
	Distancek float64       `json:"distancek"`
	Permance  time.Duration `json:"Permomance"`
}
type Knn struct {
	Ṕoints []KnnPoints `json:"knnPoints"`
}

type KnnResult struct {
	Points []KnnPoints `json:"knnPoints"`
}

func (knn KnnResult) CompareClasses(classes []string) int {
	correctClasses := 0
	for pos, elem := range knn.Points {
		result, _ := strconv.Atoi(classes[pos])
		if (elem.Class - result) == 0 {
			correctClasses += 1
		}
	}
	return correctClasses
}
func (knn Knn) AddKnnTest(KNNdatax []string, KNNdatay []string, N int, distanceMethod string, PositiveName string, Minmax bool) (KnnResult, Knn) {
	knnpoints1, _ := convertListToFloat(KNNdatax)
	knnpoints2, _ := convertListToFloat(KNNdatay)
	if Minmax {
		knnpoints1 = minmax.MinMaxData(knnpoints1)
		knnpoints2 = minmax.MinMaxData(knnpoints2)
	}

	knnTest := CreateKnnPoints(knnpoints1, knnpoints2, nil)
	return KnnResult{Points: knn.forloopKnnTest(knnTest.Ṕoints, N, distanceMethod)}, knn
}

func NewKnn(KNNdatax []string, KNNdatay []string, KNNdataClasses []string, Minmax bool) Knn {
	knnpoints1, _ := convertListToFloat(KNNdatax)
	knnpoints2, _ := convertListToFloat(KNNdatay)
	if Minmax {
		knnpoints1 = minmax.MinMaxData(knnpoints1)
		knnpoints2 = minmax.MinMaxData(knnpoints2)
	}
	knnclasses, _ := convertListToFloat(KNNdataClasses)
	knn := CreateKnnPoints(knnpoints1, knnpoints2, knnclasses)
	return knn
}

func CreateKnnPoints(knnpoints1 []float64, knnpoints2 []float64, knnclass []float64) Knn {
	knn := Knn{[]KnnPoints{}}
	var class int
	for pos, _ := range knnpoints1 {
		if knnclass != nil {
			if knnclass[pos] == 0 {
				class = 1
			} else {
				class = 0
			}
		}
		knn.Ṕoints = append(knn.Ṕoints, KnnPoints{X: knnpoints1[pos], Y: knnpoints2[pos], Class: class})
	}
	return knn
}

func (knn Knn) forloopKnnTest(newknnPoints []KnnPoints, N int, distanceMethod string) []KnnPoints {
	var class int
	fmt.Print(len(newknnPoints))
	for pos, pointT := range newknnPoints {
		timer := time.Now()
		class = knn.KnnAlgorithm(N, pointT, distanceMethod)
		newknnPoints[pos].Permance = time.Since(timer)
		newknnPoints[pos].Class = class
	}
	return newknnPoints
}

func (knn Knn) KnnAlgorithm(N int, pointT KnnPoints, distanceMethod string) int {
	var distance float64
	for pos, pointk := range knn.Ṕoints {
		if distanceMethod == "euclidian" {
			distance = Euclidian(pointk, pointT)
		} else {
			distance = Mahhatan(pointk, pointT)
		}
		knn.Ṕoints[pos].Distancek = distance
	}
	sort.Slice(knn.Ṕoints, func(i, j int) bool {
		return knn.Ṕoints[i].Distancek < knn.Ṕoints[j].Distancek
	})
	var classesCount [2]int
	for pos := 0; pos < N; pos++ {
		classesCount[knn.Ṕoints[pos].Class] += 1
	}
	if classesCount[0] > classesCount[1] {
		return 0
	} else {
		return 1
	}
}

func Euclidian(pointk KnnPoints, pointT KnnPoints) float64 {
	radicalx := math.Pow(pointk.X-pointT.X, 2)
	radicaly := math.Pow(pointk.Y-pointT.Y, 2)
	radical := radicalx + radicaly
	euclidianDistance := math.Sqrt(radical)
	return euclidianDistance
}

func Mahhatan(pointk KnnPoints, pointT KnnPoints) float64 {
	equationx := math.Abs(pointk.X - pointT.X)
	equationy := math.Abs(pointk.Y - pointT.Y)
	equation := equationx + equationy
	return equation
}

func convertListToFloat(knnpointsInt []string) ([]float64, error) {
	var knnpointsFloat []float64
	for _, stringElem := range knnpointsInt {
		floatElem, err := strconv.ParseFloat(stringElem, 64)
		if err != nil {
			return nil, err
		}
		knnpointsFloat = append(knnpointsFloat, floatElem)
	}
	return knnpointsFloat, nil
}
