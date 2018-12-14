package chronal_calibration

import (
	"fmt"
	"testing"
)

type chronalCalibrationResultTest struct {
	frequencies []int
	expected    int
}

type chronalCalibrationRepeatedTest struct {
	frequencies []int
	expected    int
}

var resultTests = []chronalCalibrationResultTest{
	{
		[]int{+7, -5, -9, -1, -3, +7, -8, -17, -4, -18, -15, +18, -11, -3, -15, +6, +17, -3, +19, +6, +4, -15, +18, -14, +8, -11, -7, -15, +19, -1, +10, +18, -1, +5, -12, +11, -5, +17, +9, -12, +7, +19, -9, +17, +12, -15, -12, +11, +14, -16, -13, +7, +5, -6, +20, -1, +8, -13, -6, -9, +4, +23, -5, -5, -19, +20, -6, +21, -4, +12, +10, -1, +16, +15, -12, -17, -3, +9, -19, -3, +7, -5, -7, -8, +2, +7, +16, -2, +13, +13, +20, -15, -10, -5, -1, +7, -15, +18, +3, -1, +16, -13, -1, -10, +15, -17, -18, -11, +1, +2, +7, +12, -14, +22, +14, +5, -13, -16, +25, +20, +2, +10, +16, -19, -1, -7, +3, +3, +10, +4, +16, +12, +8, +7, -12, -1, -7, +19, -4, +11, +18, -15, +12, +11, -10, -7, +12, -16, +2, -5, -5, +14, +9, +14, +19, +11, +6, +13, +14, +14, -6, -9, -9, -11, -5, -5, +15, -7, +1, -17, +7, -15, -7, +12, -6, -13, -19, +10, +5, -1, +6, +20, -12, -7, +18, -17, -15, -11, +16, +12, +13, -21, +9, +9, +11, -2, -13, -1, -6, +4, -16, +2, -5, +23, -13, -12, -19, +2, +2, -12, -6, +1, -21, -16, -18, +5, +18, +4, -11, +17, +3, +6, -1, -4, -18, +2, -4, -3, -15, -6, -1, -13, +10, +15, +6, +20, +18, -9, +11, -3, -18, -16, -1, +20, -7, +19, -24, +20, +12, +22, -15, +24, -21, +3, +13, -15, +23, +37, +9, +18, +3, -11, +20, +4, +15, +14, -5, +7, -6, +9, -2, -9, +12, +11, -18, +8, -15, +16, -7, +11, +11, +1, -8, +6, -8, +18, -21, +9, -7, +12, +4, +13, +15, +5, -10, -4, +10, -3, +24, -14, +10, +7, -11, -20, +19, +6, +7, +12, +1, +1, -7, +1, +20, +25, -14, -20, -15, -19, -23, -14, +9, -30, -13, -13, -20, +11, +15, -3, -9, +13, -12, -10, -9, +2, +19, -8, +3, -6, +18, -5, +15, -16, -10, -15, -6, -5, -10, +18, -16, +25, -15, -18, +12, +19, -12, +4, +6, +9, -27, -60, -48, -4, +2, -17, -12, +11, -18, -9, +5, -6, -28, +3, -13, +35, +5, +25, -22, -35, -15, -27, -10, -9, -64, -9, +10, -3, -25, +9, -14, -83, +11, +61, -8, -8, +12, +50, -20, -165, +7, -65, -3, -5, +281, +56, +121, -55464, +8, -9, -15, -3, -16, +10, -11, +9, +5, +4, +7, -20, -15, -1, -10, -11, +12, -11, +16, +21, -3, +4, +1, -9, -2, -6, -7, +9, -11, -1, -4, +15, -4, -13, -1, -1, +10, -17, -15, -17, +11, +9, +3, -9, +13, +2, -3, -17, +19, -16, -4, +8, -3, +5, -12, +16, +8, -15, -19, -8, -3, -13, -5, -16, +3, +2, +10, -5, +2, +14, +10, +8, -11, +2, -6, +9, +8, +7, +6, -3, -6, -11, +12, -23, -11, +13, +22, +18, +21, +18, +19, +6, -3, +5, +17, +7, -4, +13, -5, +13, +18, +10, -7, +12, +11, -10, -7, +14, +6, -16, +20, +16, -1, +6, -14, +15, -2, -25, +13, +39, -11, -9, +34, -60, -14, -22, -27, -12, -10, +3, +20, +3, +8, +7, -10, -14, +12, +9, -6, +8, -19, -10, -16, -14, +16, -4, -20, +10, -17, -9, -15, +16, +11, -20, -11, -14, -28, -16, -2, +21, -18, +17, +5, -6, -18, -15, -14, +15, +11, -16, -5, +19, -12, -15, -1, +11, -16, +11, +15, -19, +3, -9, -2, -9, +1, -5, +7, +11, +24, +14, +4, +17, -46, -5, -17, -8, -10, -12, -5, +10, +3, +9, +14, +7, -6, -16, +17, -18, +14, -10, +18, -17, -11, -17, -12, -15, -2, -4, -9, -13, +10, +4, -3, -9, -5, +7, +11, -17, -21, +9, +17, +17, +13, +12, -3, -3, +18, -5, +12, +10, +3, +10, +24, +21, +13, +22, +19, +72, +7, +11, -15, -5, +6, +83, +8, -153, +24, +8, +39, +43, -265, -17, +18, -4, -12, +7, -4, -2, +18, +12, +11, -7, -9, -4, +10, +15, +25, -20, -16, -22, -21, -11, +2, +13, +20, +10, +12, -6, -23, +3, -17, -7, +15, -19, -10, -8, -2, -8, -2, -6, +5, +7, -5, -3, -3, -15, +4, +5, +15, -12, -16, -13, +15, -4, -9, +12, -8, +2, +19, -4, +1, -23, -9, -10, -2, +5, -10, +2, -3, -10, -15, +19, -7, -9, -17, +10, -18, +16, -14, +7, -10, +12, +3, -7, +12, +21, -5, -11, -7, +3, +11, -16, +15, +14, +17, +8, -2, -1, +19, +10, +11, +22, -1, +18, +12, -19, -5, +17, -4, -4, +19, -46, -22, -11, +19, -12, -19, -11, -17, -15, -42, -11, -4, +20, +6, -38, +8, -58, +5, -11, -60, -7, -61, -5, +36, -82, -3, -35, -11, +33, -51, -17, +9, -277, +251, -55001, +3, -11, +1, +13, +3, -15, +1, +2, +7, -16, +13, +9, +18, -1, -10, +5, +17, -9, +2, +10, +18, +19, +3, -9, -18, -4, +6, -13, +6, +19, +15, +14, -13, +17, +14, -3, -7, -16, -9, +16, -5, -12, +19, +16, -15, +12, -5, +7, -20, -6, -1, +6, -7, +17, +22, +7, -13, +8, +16, +14, -18, +12, -4, -1, +10, -15, +14, +11, +9, -14, -3, +10, -8, -8, -16, +18, -4, +17, -3, -22, -18, -20, +19, -6, +14, -1, +4, +15, -4, +24, -31, -10, +7, -24, -1, +8, +4, -21, +8, -4, -16, +2, +4, -18, +6, -5, -19, -6, -5, -13, +17, +15, -5, +4, +2, -19, +12, -16, -8, -11, +3, +14, -5, +8, +17, -5, -11, +23, +2, -18, -30, -11, +5, +15, -14, -17, +13, +19, +2, -15, -7, -18, -14, -3, -2, -6, -12, +1, +20, +5, +2, -3, -20, -3, +9, -15, -4, -6, +2, +3, +19, +17, -3, -18, -17, -20, +14, +2, -18, +13, +13, -21, -13, -4, -18, +11, +1, +3, +10, -15, -4, -7, -2, -16, -11, -10, -5, +6, +1, -12, -8, -5, -10, +1, +5, +15, -14, -5, -3, -14, +19, +111653},
		406,
	},
}

func TestChronalCalibrationResult(t *testing.T) {
	for _, tc := range resultTests {
		t.Run(fmt.Sprintf("chronal calibration result"), func(t *testing.T) {
			calibration := NewChronalCalibration(tc.frequencies)
			result := calibration.Result()
			if tc.expected != result {
				t.Errorf("incorrect chronal calibration result: %v, expected: %v", result, tc.expected)
			}
		})
	}
}

var repeatedTests = []chronalCalibrationRepeatedTest{
	{
		[]int{+7, -5, -9, -1, -3, +7, -8, -17, -4, -18, -15, +18, -11, -3, -15, +6, +17, -3, +19, +6, +4, -15, +18, -14, +8, -11, -7, -15, +19, -1, +10, +18, -1, +5, -12, +11, -5, +17, +9, -12, +7, +19, -9, +17, +12, -15, -12, +11, +14, -16, -13, +7, +5, -6, +20, -1, +8, -13, -6, -9, +4, +23, -5, -5, -19, +20, -6, +21, -4, +12, +10, -1, +16, +15, -12, -17, -3, +9, -19, -3, +7, -5, -7, -8, +2, +7, +16, -2, +13, +13, +20, -15, -10, -5, -1, +7, -15, +18, +3, -1, +16, -13, -1, -10, +15, -17, -18, -11, +1, +2, +7, +12, -14, +22, +14, +5, -13, -16, +25, +20, +2, +10, +16, -19, -1, -7, +3, +3, +10, +4, +16, +12, +8, +7, -12, -1, -7, +19, -4, +11, +18, -15, +12, +11, -10, -7, +12, -16, +2, -5, -5, +14, +9, +14, +19, +11, +6, +13, +14, +14, -6, -9, -9, -11, -5, -5, +15, -7, +1, -17, +7, -15, -7, +12, -6, -13, -19, +10, +5, -1, +6, +20, -12, -7, +18, -17, -15, -11, +16, +12, +13, -21, +9, +9, +11, -2, -13, -1, -6, +4, -16, +2, -5, +23, -13, -12, -19, +2, +2, -12, -6, +1, -21, -16, -18, +5, +18, +4, -11, +17, +3, +6, -1, -4, -18, +2, -4, -3, -15, -6, -1, -13, +10, +15, +6, +20, +18, -9, +11, -3, -18, -16, -1, +20, -7, +19, -24, +20, +12, +22, -15, +24, -21, +3, +13, -15, +23, +37, +9, +18, +3, -11, +20, +4, +15, +14, -5, +7, -6, +9, -2, -9, +12, +11, -18, +8, -15, +16, -7, +11, +11, +1, -8, +6, -8, +18, -21, +9, -7, +12, +4, +13, +15, +5, -10, -4, +10, -3, +24, -14, +10, +7, -11, -20, +19, +6, +7, +12, +1, +1, -7, +1, +20, +25, -14, -20, -15, -19, -23, -14, +9, -30, -13, -13, -20, +11, +15, -3, -9, +13, -12, -10, -9, +2, +19, -8, +3, -6, +18, -5, +15, -16, -10, -15, -6, -5, -10, +18, -16, +25, -15, -18, +12, +19, -12, +4, +6, +9, -27, -60, -48, -4, +2, -17, -12, +11, -18, -9, +5, -6, -28, +3, -13, +35, +5, +25, -22, -35, -15, -27, -10, -9, -64, -9, +10, -3, -25, +9, -14, -83, +11, +61, -8, -8, +12, +50, -20, -165, +7, -65, -3, -5, +281, +56, +121, -55464, +8, -9, -15, -3, -16, +10, -11, +9, +5, +4, +7, -20, -15, -1, -10, -11, +12, -11, +16, +21, -3, +4, +1, -9, -2, -6, -7, +9, -11, -1, -4, +15, -4, -13, -1, -1, +10, -17, -15, -17, +11, +9, +3, -9, +13, +2, -3, -17, +19, -16, -4, +8, -3, +5, -12, +16, +8, -15, -19, -8, -3, -13, -5, -16, +3, +2, +10, -5, +2, +14, +10, +8, -11, +2, -6, +9, +8, +7, +6, -3, -6, -11, +12, -23, -11, +13, +22, +18, +21, +18, +19, +6, -3, +5, +17, +7, -4, +13, -5, +13, +18, +10, -7, +12, +11, -10, -7, +14, +6, -16, +20, +16, -1, +6, -14, +15, -2, -25, +13, +39, -11, -9, +34, -60, -14, -22, -27, -12, -10, +3, +20, +3, +8, +7, -10, -14, +12, +9, -6, +8, -19, -10, -16, -14, +16, -4, -20, +10, -17, -9, -15, +16, +11, -20, -11, -14, -28, -16, -2, +21, -18, +17, +5, -6, -18, -15, -14, +15, +11, -16, -5, +19, -12, -15, -1, +11, -16, +11, +15, -19, +3, -9, -2, -9, +1, -5, +7, +11, +24, +14, +4, +17, -46, -5, -17, -8, -10, -12, -5, +10, +3, +9, +14, +7, -6, -16, +17, -18, +14, -10, +18, -17, -11, -17, -12, -15, -2, -4, -9, -13, +10, +4, -3, -9, -5, +7, +11, -17, -21, +9, +17, +17, +13, +12, -3, -3, +18, -5, +12, +10, +3, +10, +24, +21, +13, +22, +19, +72, +7, +11, -15, -5, +6, +83, +8, -153, +24, +8, +39, +43, -265, -17, +18, -4, -12, +7, -4, -2, +18, +12, +11, -7, -9, -4, +10, +15, +25, -20, -16, -22, -21, -11, +2, +13, +20, +10, +12, -6, -23, +3, -17, -7, +15, -19, -10, -8, -2, -8, -2, -6, +5, +7, -5, -3, -3, -15, +4, +5, +15, -12, -16, -13, +15, -4, -9, +12, -8, +2, +19, -4, +1, -23, -9, -10, -2, +5, -10, +2, -3, -10, -15, +19, -7, -9, -17, +10, -18, +16, -14, +7, -10, +12, +3, -7, +12, +21, -5, -11, -7, +3, +11, -16, +15, +14, +17, +8, -2, -1, +19, +10, +11, +22, -1, +18, +12, -19, -5, +17, -4, -4, +19, -46, -22, -11, +19, -12, -19, -11, -17, -15, -42, -11, -4, +20, +6, -38, +8, -58, +5, -11, -60, -7, -61, -5, +36, -82, -3, -35, -11, +33, -51, -17, +9, -277, +251, -55001, +3, -11, +1, +13, +3, -15, +1, +2, +7, -16, +13, +9, +18, -1, -10, +5, +17, -9, +2, +10, +18, +19, +3, -9, -18, -4, +6, -13, +6, +19, +15, +14, -13, +17, +14, -3, -7, -16, -9, +16, -5, -12, +19, +16, -15, +12, -5, +7, -20, -6, -1, +6, -7, +17, +22, +7, -13, +8, +16, +14, -18, +12, -4, -1, +10, -15, +14, +11, +9, -14, -3, +10, -8, -8, -16, +18, -4, +17, -3, -22, -18, -20, +19, -6, +14, -1, +4, +15, -4, +24, -31, -10, +7, -24, -1, +8, +4, -21, +8, -4, -16, +2, +4, -18, +6, -5, -19, -6, -5, -13, +17, +15, -5, +4, +2, -19, +12, -16, -8, -11, +3, +14, -5, +8, +17, -5, -11, +23, +2, -18, -30, -11, +5, +15, -14, -17, +13, +19, +2, -15, -7, -18, -14, -3, -2, -6, -12, +1, +20, +5, +2, -3, -20, -3, +9, -15, -4, -6, +2, +3, +19, +17, -3, -18, -17, -20, +14, +2, -18, +13, +13, -21, -13, -4, -18, +11, +1, +3, +10, -15, -4, -7, -2, -16, -11, -10, -5, +6, +1, -12, -8, -5, -10, +1, +5, +15, -14, -5, -3, -14, +19, +111653},
		312,
	},
}

func TestChronalCalibrationRepeated(t *testing.T) {
	for _, tc := range repeatedTests {
		t.Run(fmt.Sprintf("chronal calibration repeated"), func(t *testing.T) {
			calibration := NewChronalCalibration(tc.frequencies)
			repeated := calibration.Repeated()
			if tc.expected != repeated {
				t.Errorf("incorrect chronal calibration repeated: %v, expected: %v", repeated, tc.expected)
			}
		})
	}
}