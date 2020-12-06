package pmm

import (
	"reflect"
	"testing"
)

func TestPkgMissingFromA(t *testing.T) {

	testCases := []struct {
		inputA   []pkgInfo
		inputB   []pkgInfo
		expected []pkgInfo
		passing  bool
	}{
		{
			[]pkgInfo{
				pkgInfo{Name: "cu-ddns"},
				pkgInfo{Name: "sonar"},
			},
			[]pkgInfo{
				pkgInfo{Name: "cu-ddns"},
				pkgInfo{Name: "sonar"},
				pkgInfo{Name: "gotham"},
			},
			[]pkgInfo{
				pkgInfo{Name: "gotham"},
			},
			true,
		},
		{
			[]pkgInfo{
				pkgInfo{Name: "cu-ddns"},
				pkgInfo{Name: "sonar"},
			},
			[]pkgInfo{
				pkgInfo{Name: "cu-ddns"},
				pkgInfo{Name: "sonar"},
			},
			[]pkgInfo{},
			true,
		},
		{
			[]pkgInfo{
				pkgInfo{Name: "sonar"},
			},
			[]pkgInfo{
				pkgInfo{Name: "Sonar"},
			},
			[]pkgInfo{},
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
