package pmm

import (
	"reflect"
	"testing"
)

func TestPkgMissingFromA(t *testing.T) {

	testCases := []struct {
		inputA   []PkgInfo
		inputB   []PkgInfo
		expected []PkgInfo
		passing  bool
	}{
		{
			[]PkgInfo{
				PkgInfo{Name: "cu-ddns"},
				PkgInfo{Name: "sonar"},
			},
			[]PkgInfo{
				PkgInfo{Name: "cu-ddns"},
				PkgInfo{Name: "sonar"},
				PkgInfo{Name: "gotham"},
			},
			[]PkgInfo{
				PkgInfo{Name: "gotham"},
			},
			true,
		},
		{
			[]PkgInfo{
				PkgInfo{Name: "cu-ddns"},
				PkgInfo{Name: "sonar"},
			},
			[]PkgInfo{
				PkgInfo{Name: "cu-ddns"},
				PkgInfo{Name: "sonar"},
			},
			[]PkgInfo{},
			true,
		},
		{
			[]PkgInfo{
				PkgInfo{Name: "sonar"},
			},
			[]PkgInfo{
				PkgInfo{Name: "Sonar"},
			},
			[]PkgInfo{},
			false,
		},
	}

	for i, tc := range testCases {

		actual := PkgMissingFromA(tc.inputA, tc.inputB)

		if (!reflect.DeepEqual(actual, tc.expected)) && tc.passing {
			t.Errorf("PkgMissingFromA test[%d]: expected %v, actual %v", i, tc.expected, actual)
		}
	}
}
