package averill

type Term struct{ Term, Location, Definition string }

var AllTerms []Term = []Term{
	{"chemistry", "1.1",
		"The study of matter and the changes that material substances" +
			" undergo."},
	{"scientific method", "1.2",
		"A procedure that consists of making observations, formulating" +
			" hypotheses, and designing experiments which lead in turn" +
			" to additional observations, hypotheses, and experiements" +
			" in repeated cycles."},
	{"qualitative observation", "1.2",
		"Observations which describe properties or occurrences in ways" +
			" that do not rely on numbers."},
	{"quantitative observation", "1.2",
		"Measurements, which by definition consist of both a number and" +
			" a unit."},
	{"hypothesis", "1.2",
		"A tentative explanation for observation(s)."},
	{"experiment", "1.2",
		"Systematic observations or measurements, perferably made under" +
			" controlled conditions -- that is, conditions in which the" +
			" variable of interest is clearly distinguished from any others."},
	{"law", "1.2",
		"A verbal or mathematical description of a phenomenon that allows" +
			" for general predictions."},
	{"theory", "1.2",
		"An unproven attempt to explain why nature behaves the way it does."},
	{"matter", "1.3",
		"Anything that occupies space and posesses mass."},
	{"mass", "1.3",
		"Quantity of matter."},
	{"weight", "1.3",
		"Force caused by the gravitational attraction that operates" +
			" on the object."},
	{"states of matter", "1.3",
		"Solids, liquids, and gases."},
	{"solid", "1.3",
		"State of matter which is characterized by fixed shape and volume."},
	{"liquid", "1.3",
		"State of matter which is characterized by fixed volume but flow" +
			" to assume the shape of their containers."},
	{"gas", "1.3",
		"State of matter which is characterized by neither fixed shape nor" +
			" fixed volume, expanding to fill their containers completely."},
	{"pressure", "1.3",
		"The amount of force exerted on a given area."},
	{"pure chemical substance", "1.3",
		"Any matter that has a fixed chemical composition and" +
			" characteristic properties."},
	{"mixture", "1.3",
		"Combination of two or more pure substances in variable" +
			" proportions in which the individual substances retain" +
			" their identity."},
	{"homogenous mixture", "1.3",
		"A mixture in which all portions of material are in the" +
			" same state, have no visible boundaries, and are uniform" +
			" throughout."},
	{"alloy", "1.3",
		"Solid solution of two or more metals."},
	{"heterogenous mixture", "1.3",
		"Solution. A mixtre in which the composition of all materials is" +
			" not uniform."},
	{"filtration", "1.3",
		"A technique for separating solids from liquids consisting" +
			" of passing the mixture through a barrier with holes or" +
			" pores that are smaller than the solid particles."},
	{"distillation", "1.3",
		"A technique for separating solutions into their component" +
			" substances by making use of differences in volatility."},
	{"volatility", "1.3",
		"A measure of how easily a substance is converted to a gas" +
			" at a given temperature."},
	{"crystallization", "1.3",
		"A techinique for separting solutions into their component" +
			" substances by making use of differences in solubility."},
	{"solubility", "1.3",
		"A measure of how much of a solid substance remains dissolved" +
			" in a given amount of a specified liquid."},
	{"solvent", "1.3",
		"A liquid in which a solid substance is dissolved into."},
	{"element", "1.3",
		"A substance that cannot be broken down into simpler ones" +
			" by chemical changes."},
	{"compound", "1.3",
		"A substance that contains two or more elements and has" +
			" chemical and physical properties taht are usually" +
			" different from those of the elements of which it is" +
			" composed."},
	{"electrolysis", "1.3",
		"A technique in which electricity provides the energy needed" +
			" to separate the compound into its constituent elements."},
	{"physical properties", "1.3",
		"Properties that are characteristics one can measure without" +
			" changing the composition of the sample."},
	{"chemical properties", "1.3",
		"Properties that describe the characteristic ability of a" +
			" substance to react to form new substances."},
	{"extensive properties", "1.3",
		"Physical properties that vary with the amount of the substance."},
	{"intensive properties", "1.3",
		"Physical properties that do not vary with the amount of the" +
			" substance."},
	{"density", "1.3",
		"Mass per unit volume."},
	{"cathode", "1.5",
		"Negatively charged electrode."},
	{"cathode rays", "1.5",
		"Light emitted from a cathode in a gas discharge tube."},
	{"radioactivity", "1.5",
		"Emission of energy rays by matter."},
	{"alpha particles", "1.5",
		"A readily absorbed, massive compared to electrons, type of radiation."},
	{"beta particles", "1.5",
		"A type of radiation with the same charge and mass-to-charge" +
			" ratio as electrons."},
	{"gamma rays", "1.5",
		"A type of radiation found to be similar to a lower-energy form" +
			" of radiation called X-rays."},
	{"nucleus", "1.5",
		"The mass and positive charge of an atom concentrated in a tiny" +
			" fraction of the volume."}}
