package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Columns struct {
	Seeds                 []int64
	SeedToSoil            []Mapping
	SoilToFertilizer      []Mapping
	FertilizerToWater     []Mapping
	WaterToLight          []Mapping
	LightToTemperature    []Mapping
	TemperatureToHumidity []Mapping
	HumidityToLocation    []Mapping
}

type Mapping struct {
	RangeLen              int64
	SourceRangeStart      int64
	DestinationRangeStart int64
}

func main() {
	file, err := os.Open("../../inputs/day05.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var columns Columns

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	line = line[7:] // "seeds: "

	for _, seed := range strings.Split(line, " ") {
		var value int64
		fmt.Sscanf(seed, "%d", &value)
		columns.Seeds = append(columns.Seeds, value)
	}

	scanner.Scan()
	scanner.Scan() // "seed-to-soil map:"
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var length, source, destination int64
		fmt.Sscanf(line, "%d %d %d", &destination, &source, &length)
		columns.SeedToSoil = append(columns.SeedToSoil, Mapping{length, source, destination})
	}

	scanner.Scan() // "soil-to-fertilizer: "
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var length, source, destination int64
		fmt.Sscanf(line, "%d %d %d", &destination, &source, &length)
		columns.SoilToFertilizer = append(columns.SoilToFertilizer, Mapping{length, source, destination})
	}

	scanner.Scan() // "fertilizer-to-water: "
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var length, source, destination int64
		fmt.Sscanf(line, "%d %d %d", &destination, &source, &length)
		columns.FertilizerToWater = append(columns.FertilizerToWater, Mapping{length, source, destination})
	}

	scanner.Scan() // "water-to-light: "
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var length, source, destination int64
		fmt.Sscanf(line, "%d %d %d", &destination, &source, &length)
		columns.WaterToLight = append(columns.WaterToLight, Mapping{length, source, destination})
	}

	scanner.Scan() // "light-to-temperature: "
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var length, source, destination int64
		fmt.Sscanf(line, "%d %d %d", &destination, &source, &length)
		columns.LightToTemperature = append(columns.LightToTemperature, Mapping{length, source, destination})
	}

	scanner.Scan() // "temperature-to-humidity:"
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var length, source, destination int64
		fmt.Sscanf(line, "%d %d %d", &destination, &source, &length)
		columns.TemperatureToHumidity = append(columns.TemperatureToHumidity, Mapping{length, source, destination})
	}

	scanner.Scan() // "humidity-to-location:"
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var length, source, destination int64
		fmt.Sscanf(line, "%d %d %d", &destination, &source, &length)
		columns.HumidityToLocation = append(columns.HumidityToLocation, Mapping{length, source, destination})
	}

	result := solvePuzzle(columns)
	fmt.Println("Solution5a:", result)

	result2 := solvePuzzle2(columns)
	fmt.Println("Solution5b:", result2)
}

func solvePuzzle(data Columns) int64 {
	var res int64 = math.MaxInt64

	for _, seed := range data.Seeds {
		soil := adjustValue(seed, data.SeedToSoil)
		fertilizer := adjustValue(soil, data.SoilToFertilizer)
		water := adjustValue(fertilizer, data.FertilizerToWater)
		light := adjustValue(water, data.WaterToLight)
		temperature := adjustValue(light, data.LightToTemperature)
		humidity := adjustValue(temperature, data.TemperatureToHumidity)
		location := adjustValue(humidity, data.HumidityToLocation)

		if location < res {
			res = location
		}
	}

	return res
}

func solvePuzzle2(data Columns) int64 {
	var res int64 = 0

exit:
	for {
		humidity := restoreValue(res, data.HumidityToLocation)
		temperature := restoreValue(humidity, data.TemperatureToHumidity)
		light := restoreValue(temperature, data.LightToTemperature)
		water := restoreValue(light, data.WaterToLight)
		fertilizer := restoreValue(water, data.FertilizerToWater)
		soil := restoreValue(fertilizer, data.SoilToFertilizer)
		seed := restoreValue(soil, data.SeedToSoil)

		for i := 0; i < len(data.Seeds); i += 2 {
			base := data.Seeds[i]
			length := data.Seeds[i+1]

			if (base <= seed) && (seed < (base + length)) {
				break exit
			}
		}

		res += 1
	}
	return res
}

func adjustValue(n int64, mappings []Mapping) int64 {
	for _, mapping := range mappings {
		if (mapping.SourceRangeStart <= n) && (n <= (mapping.SourceRangeStart + mapping.RangeLen)) {
			return mapping.DestinationRangeStart + (n - mapping.SourceRangeStart)
		}
	}
	return n
}

func restoreValue(n int64, mappings []Mapping) int64 {
	for _, mapping := range mappings {
		if (mapping.DestinationRangeStart <= n) && (n <= (mapping.DestinationRangeStart + mapping.RangeLen)) {
			return mapping.SourceRangeStart + (n - mapping.DestinationRangeStart)
		}
	}
	return n
}
