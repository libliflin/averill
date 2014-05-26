package jobs

type Job struct {
	Name  string
	Usage []string
}

var AllJobs []Job = []Job{
	{"Geologist",
		{"Chemical techniques are used to analyze and identify rock samples."}}}
