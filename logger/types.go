package logger

type getInfo interface {
	getSource() string
	getCPUs() string
	getMEMs() string
	getNotes() string
}

type infoLog struct {
	CPUs string
	MEMs string
	notes string
}

type infoCurrent struct {
	jobs [][]string
	time string
	sortFlag int
}


func (info infoCurrent) Len() int {
	return len(info.jobs)
}

func (info infoCurrent) Less(i, j int) bool {
	return info.jobs[i][info.sortFlag] > info.jobs[j][info.sortFlag]
}

func (info infoCurrent) Swap(i, j int) {
	temp := info.jobs[i]
	info.jobs[i] = info.jobs[j]
	info.jobs[j] = temp
}