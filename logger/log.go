package logger

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	recLength = 20
)
var (
	logDir = "/Users/bytedance/Documents/rookie/linux/logs.txt"
)

func GetInfoLog() ([]infoLog, error){
	file, err := os.Open(logDir)
	if err != nil {
		log.Println("[ERROR] CAN NOT READ FILE! ", err)
		return nil, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	//split
	var infos []infoLog
	records := strings.Split(string(content), "\n\n\n\n")
	for _, record := range records {
		lines := strings.Split(record, "\n")
		if len(lines) < 16 {
			continue
		}
		var info infoLog
		info.notes = lines[0]
		info.CPUs = strings.Join(lines[2:8], "\n")
		info.MEMs = strings.Join(lines[10:16], "\n")
		infos = append(infos, info)
	}
	if len(infos) >= recLength {
		infos = infos[len(infos) - recLength:]
	}
	return infos, nil
}

func (info infoLog) getSource() string {
	return "System log"
}

func (info infoLog) getCPUs() string {
	return info.CPUs
}

func (info infoLog) getMEMs() string {
	return info.MEMs
}

func (info infoLog) getNotes() string {
	return info.notes
}
