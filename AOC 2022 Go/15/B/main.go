package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"golang.org/x/exp/slices"
)

type Sensor struct {
	x, y, dis int
}

type Beacon struct {
	x, y int
}

type Range struct {
	min, max int
}

var (
	sensors        []Sensor
	beacons        []Beacon
	noBeaconRanges []Range
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var sensor Sensor
	var beacon Beacon
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		if !slices.Contains(beacons, beacon) {
			beacons = append(beacons, beacon)
		}

		dx := sensor.x - beacon.x
		if dx < 0 {
			dx *= -1
		}
		dy := sensor.y - beacon.y
		if dy < 0 {
			dy *= -1
		}

		dis := dx + dy
		sensor.dis = dis
		sensors = append(sensors, sensor)
	}

	noBeaconRanges = make([]Range, 0)
	for y := 0; y <= 4000000; y++ {
		noBeaconRanges = noBeaconRanges[:0]
		for i := range sensors {
			dy := sensors[i].y - y
			if dy < 0 {
				dy *= -1
			}

			dx := sensors[i].dis - dy
			if dx < 0 {
				continue
			}

			noBeaconRanges = append(noBeaconRanges, Range{sensors[i].x - dx, sensors[i].x + dx})
		}

		// sort ranges to make combining process easier
		sort.Slice(noBeaconRanges, func(i, j int) bool {
			return noBeaconRanges[i].min < noBeaconRanges[j].min
		})

		// combines ranges
		for i := 0; i < len(noBeaconRanges)-1; i++ {
			for j := i + 1; j < len(noBeaconRanges); {
				// if overlaps
				if noBeaconRanges[i].max >= noBeaconRanges[j].min {
					if noBeaconRanges[i].max < noBeaconRanges[j].max {
						noBeaconRanges[i].max = noBeaconRanges[j].max
					}
					noBeaconRanges = append(noBeaconRanges[:j], noBeaconRanges[j+1:]...)
				} else {
					j++
				}
			}
		}

		// special cases to test if they contain empty spot
		if len(noBeaconRanges) >= 2 {
			fmt.Println(y, noBeaconRanges)
			for i := 0; i < len(noBeaconRanges)-1; i++ {
				minX := noBeaconRanges[i].max + 1
				maxX := noBeaconRanges[i+1].min - 1
				for x := minX; x <= maxX; x++ {
					if !slices.Contains(beacons, Beacon{x, y}) {
						fmt.Println(x*4000000 + y)
					}
				}
			}
		}
		if noBeaconRanges[0].min > 0 {
			fmt.Println(y, noBeaconRanges)
			maxX := noBeaconRanges[0].min - 1
			for x := 0; x <= maxX; x++ {
				if !slices.Contains(beacons, Beacon{x, y}) {
					fmt.Println(x*4000000 + y)
				}
			}
		}
		if noBeaconRanges[0].max < 4000000 {
			fmt.Println(y, noBeaconRanges)
			minX := noBeaconRanges[len(noBeaconRanges)-1].max + 1
			for x := minX; x <= 4000000; x++ {
				if !slices.Contains(beacons, Beacon{x, y}) {
					fmt.Println(x*4000000 + y)
				}
			}
		}
	}
}
