package logger

import (
	"bytes"
	"fmt"
	"github.com/themakers/osinfo"
	"os/exec"
)

func GetOsInfo() string {
	osi := osinfo.GetInfo()
	return fmt.Sprintf("Kernel: %s \nCore: %s \nPlatform: %s \nOS: %s\n",
		osi.Kernel, osi.Core, osi.Platform, osi.OS)
}

func execShell(s string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func writeInfo(info getInfo) string{
	var out string
	out += fmt.Sprintln("Data Source:", info.getSource())
	out += fmt.Sprintln("NOTE:", info.getNotes())
	out += fmt.Sprintf("TOP 5 CPU USAGE PROCESSES: \n%s\n", info.getCPUs())
	out += fmt.Sprintf("\nTOP 5 MEM USAGE PROCESSES: \n%s\n\n\n", info.getMEMs())
	return out
}

func FromGo() string{
	info, err := GetInfoGo()
	if err != nil {
		return ""
	}
	out := fmt.Sprintf("Operating System: \n%s\n", GetOsInfo())
	out += writeInfo(info)
	return out
}

func FromLog() string{
	infos, err := GetInfoLog()
	if err != nil {
		return string(err.Error())
	}

	out := fmt.Sprintf("Operating System: \n%s\n", GetOsInfo())

	for _, info := range infos{
		out += writeInfo(info)
	}
	return out
}
