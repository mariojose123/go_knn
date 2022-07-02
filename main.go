package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/mariojose123/knngo/knn"
	"github.com/mariojose123/knngo/strafiedSamplingTestTrainingSplit"
	"io/ioutil"
	"os"
	"time"
)

type result struct {
	X          int           `json:"colx"`
	Y          int           `json:"coly"`
	K          int           `json:"k"`
	Totalright int           `json:"totalright"`
	Perfomance time.Duration `json: "Perfomance"`
	UseMinMax  bool          `json: "UseMinMax"`
	Distance   string        `json: "distance"`
}

func readCsvFile(file string) (records [][]string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	return data, err
}

func main() {
	var csvFileLocation string
	fmt.Print("Please tell the location of  Breast Cancer Dataset")
	fmt.Scanln(&csvFileLocation)
	data, _ := readCsvFile(csvFileLocation)
	fmt.Print("Okay lets Data Science")
	results := []result{}
	ks := []int{1, 3, 5, 7}
	minmaxbools := []bool{true, false}
	distnacetypes := []string{"euclidian", "Man"}
	for _, minmaxbool := range minmaxbools {
		for _, distancetype := range distnacetypes {
			for _, k := range ks {
				for i := 1; i < 32; i++ {
					for j := 1; j < 32-i+1; j++ {
						if i != j {
							train, test := strafiedSamplingTestTrainingSplit.StratifiedHoldout(data, i, j, 0.2, 31)
							knnPreTrained := knn.NewKnn(train.X, train.Y, train.Class, minmaxbool)
							t := time.Now()
							knnPosTest, _ := knnPreTrained.AddKnnTest(test.X, test.Y, k, distancetype, "1", minmaxbool)
							Perfomance := time.Since(t)
							var result result
							result.X = i
							result.Y = j
							result.K = k
							result.Perfomance = Perfomance
							result.UseMinMax = minmaxbool
							result.Distance = distancetype
							result.Totalright = knnPosTest.CompareClasses(test.Class)
							results = append(results, result)
						}
					}
				}
			}
		}
	}
	jsonString, _ := json.Marshal(results)
	ioutil.WriteFile("test2.json", jsonString, os.ModePerm)
}
