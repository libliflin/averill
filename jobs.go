package averill

type Job struct {
	Name  string
	Usage []string
}

var AllJobs []Job = []Job{
	{"Geologist", []string{"Chemical techniques are used to analyze and" +
		" identify rock samples."}}}
