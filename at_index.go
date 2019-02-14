package ansigo

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

var codeRE = regexp.MustCompile("\x1b\\[(\\d+)(;\\d+)*m")

func ANSIAtIndex(s string, index int) []string {
	if index < 5 || index >= len(s) {
		return []string{}
	}
	var activeForeground, activeBackground string
	activeAttributes := map[int]bool{
		1:  false,
		2:  false,
		3:  false,
		4:  false,
		5:  false,
		6:  false,
		7:  false,
		8:  false,
		9:  false,
		11: false,
		12: false,
		13: false,
		14: false,
		15: false,
		16: false,
		17: false,
		18: false,
		19: false,
		20: false,
		21: false,
		51: false,
		52: false,
		53: false,
		60: false,
		61: false,
		62: false,
		63: false,
		64: false,
	}
	endsToStarts := map[int][]int{
		22: []int{1, 2},
		23: []int{3, 20},
		24: []int{4, 21},
		25: []int{5, 6},
		27: []int{7},
		28: []int{8},
		29: []int{9},
		10: []int{11, 12, 13, 14, 15, 16, 17, 18, 19},
		54: []int{51, 52},
		55: []int{53},
		65: []int{60, 61, 62, 63, 64},
	}
	for _, match := range codeRE.FindAllStringSubmatch(s[:index], -1) {
		core, _ := strconv.Atoi(match[1])
		if core == 39 {
			activeForeground = ""
		} else if core == 49 {
			activeBackground = ""
		} else if _, ok := activeAttributes[core]; ok {
			activeAttributes[core] = true
		} else if starts, ok := endsToStarts[core]; ok {
			for _, a := range starts {
				activeAttributes[a] = false
			}
		} else if (core >= 30 && core <= 38) || (core >= 90 && core <= 97) {
			activeForeground = match[0]
		} else if (core >= 40 && core <= 48) || (core >= 100 && core <= 107) {
			activeBackground = match[0]
		}
	}
	activeCodes := []string{}
	if activeForeground != "" {
		activeCodes = append(activeCodes, activeForeground)
	}
	if activeBackground != "" {
		activeCodes = append(activeCodes, activeBackground)
	}
	for code, status := range activeAttributes {
		if status {
			activeCodes = append(activeCodes, fmt.Sprintf("\x1b[%dm", code))
		}
	}
	sort.Strings(activeCodes)
	return activeCodes
}
