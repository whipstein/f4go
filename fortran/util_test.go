package fortran

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tcs := []struct {
		in                                                 string
		pos                                                int
		leftOther, leftVariable, rightVariable, rightOther string
	}{
		{
			in:            "t(2**3)",
			pos:           3,
			leftOther:     "t(",
			leftVariable:  "2",
			rightVariable: "3",
			rightOther:    ")",
		},
		{
			in:            "t(replace,2**3)",
			pos:           5,
			leftOther:     "t(replace,",
			leftVariable:  "2",
			rightVariable: "3",
			rightOther:    ")",
		},
		{
			in:            "t(2**3,replace)",
			pos:           3,
			leftOther:     "t(",
			leftVariable:  "2",
			rightVariable: "3",
			rightOther:    ",replace)",
		},
		{
			in:            "t(a(:1) // b(:1))",
			pos:           7,
			leftOther:     "t(",
			leftVariable:  "a(:1)",
			rightVariable: "b(:1)",
			rightOther:    ")",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.in, func(t *testing.T) {
			var p parser
			p.initVars = map[string]goType{}
			p.initVars["a"] = goType{baseType: "byte"}
			p.initVars["b"] = goType{baseType: "byte"}

			nodes := scan([]byte(tc.in))
			leftOther, leftVariable, rightVariable, rightOther :=
				p.split(&nodes, tc.pos)

			compare := []struct {
				name   string
				expect string
				actual string
			}{
				{
					name:   "leftOther",
					expect: strings.Replace(tc.leftOther, " ", "", -1),
					actual: strings.Replace(nodesToString(leftOther), " ", "", -1),
				},
				{
					name:   "leftVariable",
					expect: strings.Replace(tc.leftVariable, " ", "", -1),
					actual: strings.Replace(nodesToString(leftVariable), " ", "", -1),
				},
				{
					name:   "rightVariable",
					expect: strings.Replace(tc.rightVariable, " ", "", -1),
					actual: strings.Replace(nodesToString(rightVariable), " ", "", -1),
				},
				{
					name:   "rightOther",
					expect: strings.Replace(tc.rightOther, " ", "", -1),
					actual: strings.Replace(nodesToString(rightOther), " ", "", -1),
				},
			}

			for _, c := range compare {
				t.Run(c.name, func(t *testing.T) {
					if c.expect != c.actual {
						t.Fatalf("%s not same: `%s` != `%s`",
							c.name, c.expect, c.actual)
					}
				})
			}
		})
	}
}