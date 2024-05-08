package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
  data := make(map[string][]float64, 10000)
  stations := make([]string, 0, 10000)

  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    parts := strings.Split(scanner.Text(), ";")
    station := parts[0]
    measurement, _ := strconv.ParseFloat(parts[1], 64)

    measurements, ok := data[station]
    if ok {
      min, max, _, _ := split(measurements)
      measurements[0] = minimum(min, measurement)
      measurements[1] = maximum(max, measurement)
      measurements[2] += measurement
      measurements[3] += 1
    } else {
      data[station] = []float64 { measurement, measurement, measurement, 1 }
      stations = append(stations, station)
    }
  }

  if err := scanner.Err(); err != nil {
    log.Println(err)
  }

  sortAndPrint(stations, data)
}

func split(data []float64) (float64, float64, float64, float64) {
  return data[0], data[1], data[2], data[3]
}

func minimum(a, b float64) float64 {
  if a <= b { return a }
  return b
}

func maximum(a, b float64) float64 {
  if a >= b { return a }
  return b
}

func sortAndPrint(stations []string, data map[string][]float64) {
  sort.Strings(stations)

  var builder strings.Builder
  for _, station := range(stations) {
    min, max, sum, count := split(data[station])
    mean := math.Round(sum / count * 10) / 10
    output := fmt.Sprintf("%v=%v/%v/%v", station, min, mean, max)
    if builder.Len() != 0 {
      builder.WriteString(", ")
    }
    builder.WriteString(output)
  }

  fmt.Print(builder.String())
}
