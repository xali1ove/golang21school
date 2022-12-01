package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"github.com/spf13/pflag"
	"io"
	"os"
)

func ParseValue() ([]float64, bool) {
	ok := false
	var num float64
	mass := make([]float64, 0)

	for _, err := fmt.Scanln(&num); err != io.EOF; _, err = fmt.Scanln(&num) {
		if num > 100000 || num < -100000 {
			errRange := fmt.Errorf("%s", "num is outy of range!!")
			fmt.Println(errRange)
			ok = true
			break
		}
		if err != nil {
			errType := fmt.Errorf("%s", "error in type of num")
			fmt.Println(errType)
			ok = true
			break
		}
		mass = append(mass, num)
	}
	return mass, ok
}

func GetResult(tMean, tMedian, tMode, tSD *bool, mass []float64) {
	if *tMean {
		resMean, _ := stats.Mean(mass)
		fmt.Println("Mean: ", resMean)
	}
	if *tMedian {
		resMedian, _ := stats.Median(mass)
		fmt.Println("Median: ", resMedian)
	}
	if *tMode {
		tmpArr, _ := stats.Mode(mass)
		if len(tmpArr) == 0 {
			min, _ := stats.Min(mass)
			fmt.Println("Mode: ", min)
		} else {
			min, _ := stats.Min(tmpArr)
			fmt.Println("Mode: ", min)
		}
	}
	if *tSD {
		tmpRes := stats.NormFit(mass)
		fmt.Printf("SD: %.2f\n", tmpRes[1])
	}
}

func main() {
	tMean := pflag.Bool("Mean", false, "The arithmetic mean in the array")
	tMedian := pflag.Bool("Median", false, "Median is a middle number of a sorted sequence")
	tMode := pflag.Bool("Mode", false, "Mode is a number which is occurring most frequently")
	tSD := pflag.Bool("SD", false, "Mean square deviation")
	help := pflag.Bool("help", false, "Navigation")
	pflag.Parse()
	if len(os.Args) < 2 {
		*tMean = true
		*tMedian = true
		*tMode = true
		*tSD = true
	}
	if *help {
		pflag.Usage()
		fmt.Println("Please select any flag")
		return
	}
	mass, ok := ParseValue()
	if !ok {
		GetResult(tMean, tMedian, tMode, tSD, mass)
	}
}
