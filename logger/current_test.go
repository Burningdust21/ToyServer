package logger

import (
	"reflect"
	"testing"
)

func TestInfoCurrent_getNotes(t *testing.T) {
	tests := []struct {
		name   string
		want   string
	}{
		{
			name: "basic",
			want: "Provided By Golang",
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			info := infoCurrent{ }
			got := info.getSource()

			if !reflect.DeepEqual(got, testcase.want){
				t.Errorf("getSource() = %v,  want %v", got,  testcase.want)
			}
		})
	}
}

func TestInfoCurrent_toList(t *testing.T) {
	tests := []struct {
		name   string
		want   []string
	}{
		{
			name: "basic",
			want: []string{
				"USER               PID  %CPU %MEM      VSZ    RSS   TT  STAT STARTED      TIME COMMAND",
				"bytedance         2769   9.8  0.5  6363600  90916   ??     R 10:31AM   0:52.12 3",
				"bytedance         1003  58.1  3.6 14192932 603892   ??     R  9:53AM 254:03.81 -A -u -F vms/0/hyperkit.pid -c",
				"bytedance        27410   4.9  0.6  5141040  95916   ??     S  4:49PM   0:03.83 Monitor.app/Contents/MacOS/Act",
				"bytedance         2769   9.8  0.5  6363600  90916   ??     R 10:31AM   0:52.12 ",
				"bytedance         1003  58.1  3.6 14192932 603892   ??     R  9:53AM 254:03.81 -A -u -F vms/0/hyperkit.pid -c",
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			info := infoCurrent{
				jobs: [][]string{
					{"bytedance", "2769", "9.8", "0.5", "6363600", "90916", "??", "R", "10:31AM", "0:52.12", "3"},
					{"bytedance", "1003", "58.1", "3.6", "14192932", "603892", "??", "R", "9:53AM", "254:03.81", "-A -u -F vms/0/hyperkit.pid -c"},
					{"bytedance", "27410", "4.9", "0.6", "5141040", "95916", "??", "S", "4:49PM", "0:03.83", "Monitor.app/Contents/MacOS/Act"},
					{"bytedance", "2769", "9.8", "0.5", "6363600", "90916", "??", "R", "10:31AM", "0:52.12", ""},
					{"bytedance", "1003", "58.1", "3.6", "14192932", "603892", "??", "R", "9:53AM", "254:03.81", "-A -u -F vms/0/hyperkit.pid -c"},
					{"bytedance", "27410", "4.9", "0.6", "5141040", "95916", "??", "S", "4:49PM", "0:03.83", "Monitor.app/Contents/MacOS/Act"},
				},
			}

			got := info.toList()
			if !reflect.DeepEqual(got, testcase.want){
				t.Errorf("getSource() = %v,  \n want %v", got,  testcase.want)
			}
		})
	}
}