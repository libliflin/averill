package laws

type Law struct {
	Name           string
	Discoverer     string
	DiscoveredYear string
	Statement      string
}

var AllLaws []Law = []Law{
	{"law of definite proportions",
		"Joseph Proust",
		"1754-1826",
		"A chemical substance always contains the same proportions of" +
			" elements by mass."},
	{"law of multiple proportions",
		"John Dalton",
		"1766-1844",
		"When to elements form a series of compounds, the ratios of" +
			" the masses of the second element that are present per" +
			" gram of the first element can almost always be expressed" +
			" as the ratio of two integers."}}
