package strafiedSamplingTestTrainingSplit

import (
	"math"
	"math/rand"
)

type Datapoints2Dstring struct {
	X     []string
	Y     []string
	Class []string
}

func (data *Datapoints2Dstring) SetXYClass(X string, Y string, class string) {
	data.X = append(data.X, X)
	data.Y = append(data.Y, Y)
	data.Class = append(data.Class, class)
}

func FilterIndex(data [][]string, index int) []string {
	var filteredData []string
	for pos, _ := range data {
		if pos != 0 {
			filteredData = append(filteredData, data[pos][index])
		}
	}
	return filteredData
}
func StratifiedHoldout(data [][]string, ColIndex1 int, ColIndex2 int, split float64, indexClass int) (Datapoints2Dstring, Datapoints2Dstring) {
	Datapoints2DRaw := SortData(Datapoints2Dstring{X: FilterIndex(data, ColIndex1), Y: FilterIndex(data, ColIndex2), Class: FilterIndex(data, indexClass)})
	Datapoints2DRPositive := Datapoints2Dstring{X: []string{}, Y: []string{}, Class: []string{}}
	Datapoints2DRNegative := Datapoints2Dstring{X: []string{}, Y: []string{}, Class: []string{}}

	for pos, _ := range Datapoints2DRaw.X {
		x := Datapoints2DRaw.X[pos]
		y := Datapoints2DRaw.Y[pos]
		class := Datapoints2DRaw.Class[pos]
		if class == "0" {
			Datapoints2DRNegative.SetXYClass(x, y, class)
		} else {
			Datapoints2DRPositive.SetXYClass(x, y, class)
		}
	}
	Data2DTestPositive, Datapoints2DRTrainPositive := SplitData(Datapoints2DRPositive, split)
	Data2DTestNegative, Datapoints2DRTrainNegative := SplitData(Datapoints2DRNegative, split)
	Data2DTest := UnionData(Data2DTestPositive, Data2DTestNegative)
	Data2DTrain := UnionData(Datapoints2DRTrainPositive, Datapoints2DRTrainNegative)
	return Data2DTrain, Data2DTest
}

func UnionData(data2d1 Datapoints2Dstring, data2d2 Datapoints2Dstring) Datapoints2Dstring {
	data2d1.X = append(data2d1.X, data2d2.X...)
	data2d1.Y = append(data2d1.Y, data2d2.Y...)
	data2d1.Class = append(data2d1.Class, data2d2.Class...)
	return data2d1
}
func SplitData(data2d Datapoints2Dstring, split float64) (Datapoints2Dstring, Datapoints2Dstring) {
	Datapoints2DRTest := Datapoints2Dstring{}
	Datapoints2DRTrain := Datapoints2Dstring{}
	var lenData float64 = float64(len(data2d.X))
	splitOnList := int(math.Floor(lenData * (1 - split)))
	Datapoints2DRTrain.X = data2d.X[:splitOnList]
	Datapoints2DRTrain.Y = data2d.Y[:splitOnList]
	Datapoints2DRTrain.Class = data2d.Class[:splitOnList]
	Datapoints2DRTest.X = data2d.X[splitOnList:]
	Datapoints2DRTest.Y = data2d.Y[splitOnList:]
	Datapoints2DRTest.Class = data2d.Class[splitOnList:]
	return Datapoints2DRTest, Datapoints2DRTrain
}

func SortData(data2d Datapoints2Dstring) Datapoints2Dstring {
	rand.Seed(42)
	perm := rand.Perm(len(data2d.X))
	for changed, change := range perm {
		data2d.X[changed], data2d.X[change] = data2d.X[change], data2d.X[changed]
		data2d.Y[changed], data2d.Y[change] = data2d.Y[change], data2d.Y[changed]
		data2d.Class[changed], data2d.Class[change] = data2d.Class[change], data2d.Class[changed]
	}
	return data2d
}
