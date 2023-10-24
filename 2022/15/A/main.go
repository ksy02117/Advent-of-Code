package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"golang.org/x/exp/slices"
)

type Coord struct {
	x, y int
}

type Range struct {
	min, max int
}

var (
	beacons        []Coord
	noBeaconRanges []Range
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var sensor, beacon Coord
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
		dy = sensor.y - 2000000
		if dy < 0 {
			dy *= -1
		}

		dx = dis - dy
		if dx < 0 {
			continue
		}

		noBeaconRanges = append(noBeaconRanges, Range{sensor.x - dx, sensor.x + dx})
	}

	sort.Slice(noBeaconRanges, func(i, j int) bool {
		return noBeaconRanges[i].min < noBeaconRanges[j].min
	})

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

	sum := 0
	for i := range noBeaconRanges {
		sum += noBeaconRanges[i].max - noBeaconRanges[i].min + 1
	}

	cnt := 0
	for i := range beacons {
		if beacons[i].y == 2000000 {
			cnt++
		}
	}

	fmt.Println(noBeaconRanges)
	fmt.Println(sum - cnt)
}
