package logger

import (
	"os"
	"reflect"
	"testing"
)

func TestGetInfoLog(t *testing.T){

	tests := []struct {
		name   string
		dir    string
		err    error
		want   []infoLog
	}{
		{
			name: "basic",
			dir: "./tester.txt",
			err:  nil,
			want: []infoLog{
				{
					CPUs: "USER                PID    %MEM %CPU\n" +
						"bytedance           927    1.4  72.3\n" +
						"bytedance           7317   7.3  26.4\n" +
						"_windowserver       221    0.4  9.1\n" +
						"bytedance           4239   1.4  8.1\n" +
						"bytedance           6104   0.4  1.9",
					MEMs: "USER                PID    %MEM %CPU\n" +
						"bytedance           7317   7.3  23.6\n" +
						"bytedance           927    1.4  70.9\n" +
						"bytedance           4239   1.4  9.5\n" +
						"bytedance           29314  1.3  0.0\n" +
						"bytedance           38427  1.3  0.0",

					notes: "2021-02-18 19:47:00 Work load too high current 2.27",
				},
				{
					CPUs: "USER                PID    %MEM %CPU\n" +
						"bytedance           927    1.4  62.3\n" +
						"_windowserver       221    0.4  14.0\n" +
						"bytedance           4239   1.4  8.1\n" +
						"root                305    0.0  6.3\n" +
						"bytedance           38611  0.1  5.3",
					MEMs: "USER                PID    %MEM %CPU\n" +
						"bytedance           7317   6.9  7.3\n" +
						"bytedance           927    1.4  77.4\n" +
						"bytedance           4239   1.4  3.3\n" +
						"bytedance           29314  1.3  0.0\n" +
						"bytedance           38427  1.3  0.1",
					notes: "2021-02-18 19:48:01 Work load too high current 2.29",
				},
			},
		},

		{
			name: "error",
			dir: "./no_such_file.txt",
			err:  nil,
			want: nil,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			logDir = testcase.dir
			got, err := GetInfoLog()
			_, testcase.err = os.Open(logDir)

			if !reflect.DeepEqual(got, testcase.want) || !reflect.DeepEqual(err, testcase.err){
				t.Errorf("GetInfoLog() = %v, %v, want %v, %v", got, err, testcase.want, testcase.err)
			}

		})
	}
}

func TestInfoLog_getSource(t *testing.T) {
	tests := []struct {
		name   string
		want   string
	}{
		{
			name: "basic",
			want: "System log",
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			info := infoLog{}
			got := info.getSource()

			if !reflect.DeepEqual(got, testcase.want){
				t.Errorf("getSource() = %v,  want %v", got,  testcase.want)
			}

		})
	}
}

func TestInfoLog_getCPUs(t *testing.T) {
	tests := []struct {
		name   string
		want   string
	}{
		{
			name: "basic",
			want: "CPU TEST",
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			info := infoLog{
				CPUs: "CPU TEST",
				MEMs: "MEMs TEST",
				notes: "notes TEST",
			}
			got := info.getCPUs()

			if !reflect.DeepEqual(got, testcase.want){
				t.Errorf("getCPUs() = %v,  want %v", got,  testcase.want)
			}
		})
	}
}

func TestInfoLog_getMEMs(t *testing.T) {
	tests := []struct {
		name   string
		want   string
	}{
		{
			name: "basic",
			want: "MEMs TEST",
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			info := infoLog{
				CPUs: "CPU TEST",
				MEMs: "MEMs TEST",
				notes: "notes TEST",
			}
			got := info.getMEMs()

			if !reflect.DeepEqual(got, testcase.want){
				t.Errorf("getMEMs() = %v,  want %v", got,  testcase.want)
			}
		})
	}
}

func TestInfoLog_getNotes(t *testing.T) {
	tests := []struct {
		name   string
		want   string
	}{
		{
			name: "basic",
			want: "notes TEST",
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			info := infoLog{
				CPUs: "CPU TEST",
				MEMs: "MEMs TEST",
				notes: "notes TEST",
			}
			got := info.getNotes()

			if !reflect.DeepEqual(got, testcase.want){
				t.Errorf("getNotes() = %v,  want %v", got,  testcase.want)
			}
		})
	}
}

