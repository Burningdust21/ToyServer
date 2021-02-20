package logger

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	header = []string{"USER", "PID", "%CPU", "%MEM", "VSZ", "RSS", "TT", "STAT", "STARTED", "TIME", "COMMAND"}
	format = "%-16s%6s%6s%5s%9s%7s%5s%6s%8s%10s %s"
	recNum = 5
	cmdLength = 30
)

func GetInfoGo() (infoCurrent, error){
	var info infoCurrent
	//Get raw processes states from system
	raw, err := execShell("ps aux")
	if err != nil {
		log.Fatalln("[ERROR] `ps aus` FAILED:", err)
		return info, err
	}

	processes := strings.Split(raw, "\n")
	for _, process := range processes[1:]  {
		line := strings.Fields(process)
		if len(line) < 11 {
			break
		}
		line[10] = strings.Join(line[11:], " ")
		if len(line[10]) > cmdLength {
			line[10] = line[10][:cmdLength]
		}
		line = line[:11]
		info.jobs = append(info.jobs, line)
	}

	info.time = time.Now().Format(time.RFC3339)
	return info, nil
}


func (info infoCurrent) getSource() string {
	return "Provided By Golang"
}

func (info infoCurrent) getCPUs() string {
	//sort according to CPU usage
	info.sortFlag = 2
	sort.Sort(info)

	return strings.Join(info.toList(), "\n")
}

func (info infoCurrent) getMEMs() string {
	info.sortFlag = 3
	sort.Sort(info)
	return strings.Join(info.toList(), "\n")
}

func (info infoCurrent) getNotes() string {
	return info.time
}

func (info infoCurrent) toList() []string {
	//set up header
	inter := make([]interface{}, len(header))
	for i, v := range header {
		inter[i] = v
	}
	out := []string{fmt.Sprintf(format, inter...)}

	//combine to one string
	for _, record := range info.jobs {
		inter := make([]interface{}, len(record))
		for i, v := range record {
			inter[i] = v
		}

		s := fmt.Sprintf(format, inter...)
		out = append(out, s)
	}
	fmt.Println("Number of processes: ", len(out)-1)
	if len(out) > 5{
		out = out[:recNum+1]
	}

	return out
}