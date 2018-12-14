package chronal_calibration

type ChronalCalibration struct {
	Frequencies []int
}

func NewChronalCalibration(frequencies []int) *ChronalCalibration {
	return &ChronalCalibration{Frequencies: frequencies}
}

// Part One
func (cc *ChronalCalibration) Result() (currFrequency int) {
	currFrequency = 0
	for _, f := range cc.Frequencies {
		currFrequency += f
	}
	return
}

// Part Two
func (cc *ChronalCalibration) Repeated() int {
	currFrequency := 0
	visited := map[int]bool{0: true}

	for {
		for _, f := range cc.Frequencies {
			currFrequency += f
			if visited[currFrequency] {
				return currFrequency
			}
			visited[currFrequency] = true
		}
	}
}
