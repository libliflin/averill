package averill

import (
	"bytes"
	"errors"
	"fmt"
	"math"
)

type Compound struct {
	Name         string
	DensityAt25C float64
	Description  string
}

var AllCompounds []Compound = []Compound{
	{"Water", 0.998, ""},
	{"Ethanol", 0.789, "\"alcohol\" in beverages"},
	{"Methanol", 0.792, "wood alcohol"},
	{"Ethylene glycol", 1.113, "used in antifreeze"},
	{"Diethyl ether", 0.708, "\"ether\": once widely used as an anesthetic"},
	{"Isopropanol", 0.785, "rubbing alcohol"},
	{"Benzene", 0.8787, ""},
	{"Toluene", 0.8669, ""},
	{"m-Xylene", 0.8684, ""},
	{"Isooctane", 0.6979, ""},
	{"Methyl t-butyl ether", 0.7405, ""},
	{"t-Butyl alcohol", 0.7856, ""}}

// if only one compound is known within epsilon = .005, return that
// compound, else return error with the name of the compounds
func Compound_from_mL_and_g(mL, g float64) (*Compound, error) {
	eps := .005
	d := g / mL
	var possible []*Compound
	for i := 0; i < len(AllCompounds); i++ {
		pd := AllCompounds[i].DensityAt25C
		if math.Abs(pd-d) < eps {
			append(possible, AllCompounds[i])
		}
	}
	if len(possible) == 1 {
		return possible[0], nil
	} else if len(possible) > 1 {
		var buf bytes.Buffer
		for i := 0; i < len(possible); i++ {
			buf.WriteString(possible[i].Name)
			buf.WriteString(" ")
		}
		return nil, errors.New(buf.String())
	} else {
		return nil, errors.New(fmt.Sprintf("No compounds found with density within"+
			" %d of %d", eps, d))
	}
}
