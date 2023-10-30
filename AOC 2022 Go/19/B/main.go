package main

import (
	"bufio"
	"fmt"
	"os"
)

type Blueprint struct {
	ID int

	oreRobotCost  int
	clayRobotCost int

	obsidianRobotCostOre  int
	obsidianRobotCostClay int

	geodeRobotCostOre      int
	geodeRobotCostObsidian int
}

type Resources struct {
	oreCount      int
	oreRobotCount int

	clayCount      int
	clayRobotCount int

	obsidianCount      int
	obsidianRobotCount int

	geodeCount      int
	geodeRobotCount int
}

func (r *Resources) Collect() {
	r.oreCount += r.oreRobotCount
	r.clayCount += r.clayRobotCount
	r.obsidianCount += r.obsidianRobotCount
	r.geodeCount += r.geodeRobotCount
}

func (r Resources) MakeNoRobot() Resources {
	r.Collect()

	return r
}

func (r Resources) MakeOreRobot(b Blueprint) Resources {
	r.Collect()

	r.oreRobotCount++
	r.oreCount -= b.oreRobotCost

	return r
}

func (r Resources) MakeClayRobot(b Blueprint) Resources {
	r.Collect()

	r.clayRobotCount++
	r.oreCount -= b.clayRobotCost

	return r
}

func (r Resources) MakeObsidianRobot(b Blueprint) Resources {
	r.Collect()

	r.obsidianRobotCount++
	r.oreCount -= b.obsidianRobotCostOre
	r.clayCount -= b.obsidianRobotCostClay

	return r
}

func (r Resources) MakeGeodeRobot(b Blueprint) Resources {
	r.Collect()

	r.geodeRobotCount++
	r.obsidianCount -= b.geodeRobotCostObsidian
	r.oreCount -= b.geodeRobotCostOre

	return r
}

func (b Blueprint) mineGeode(minutes int, r Resources) int {
	if minutes == 0 {
		return r.geodeCount
	}

	max := 0

	// do nothing path
	result := b.mineGeode(minutes-1, r.MakeNoRobot())
	if result > max {
		max = result
	}

	// Make 1 Ore Robot (There is no reason to make ore robot if you could have done it in previous minute)
	if r.oreCount >= b.oreRobotCost && r.oreCount-r.oreRobotCount < b.oreRobotCost {
		result = b.mineGeode(minutes-1, r.MakeOreRobot(b))
		if result > max {
			max = result
		}
	}

	// Make 1 Clay Robot
	if r.oreCount >= b.clayRobotCost && r.oreCount-r.oreRobotCount < b.clayRobotCost {
		result = b.mineGeode(minutes-1, r.MakeClayRobot(b))
		if result > max {
			max = result
		}
	}

	// Make 1 Obsidian Robot
	if r.oreCount >= b.obsidianRobotCostOre && r.clayCount >= b.obsidianRobotCostClay {
		result = b.mineGeode(minutes-1, r.MakeObsidianRobot(b))
		if result > max {
			max = result
		}
	}

	// Make 1 Geode Robot
	if r.oreCount >= b.geodeRobotCostOre && r.obsidianCount >= b.geodeRobotCostObsidian {
		result = b.mineGeode(minutes-1, r.MakeGeodeRobot(b))
		if result > max {
			max = result
		}
	}

	return max
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var blueprint Blueprint
	answer := 1

	for i := 0; i < 3; i++ {
		if !sc.Scan() {
			break
		}
		fmt.Sscanf(
			sc.Text(),
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&blueprint.ID, &blueprint.oreRobotCost, &blueprint.clayRobotCost, &blueprint.obsidianRobotCostOre, &blueprint.obsidianRobotCostClay, &blueprint.geodeRobotCostOre, &blueprint.geodeRobotCostObsidian)

		// if blueprint.ID == 0 ||
		// 	blueprint.oreRobotCost == 0 || blueprint.clayRobotCost == 0 ||
		// 	blueprint.obsidianRobotCostOre == 0 || blueprint.obsidianRobotCostClay == 0 ||
		// 	blueprint.geodeRobotCostOre == 0 || blueprint.geodeRobotCostObsidian == 0 {
		// 	panic("Failed to read input")
		// }

		numberOfGeode := blueprint.mineGeode(32, Resources{0, 1, 0, 0, 0, 0, 0, 0})

		fmt.Println(blueprint.ID, ": ", numberOfGeode)

		answer *= numberOfGeode
	}

	fmt.Println(answer)
}
