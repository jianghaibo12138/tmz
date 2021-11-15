package tools

import (
	"fmt"
	"path"
	"stp_go/configs"
	"testing"
)

func TestZipDir(t *testing.T) {
	logDirPath := path.Join(configs.GetHomePath(), "logs")
	zipFileName := path.Join(configs.GetHomePath(), "logs.zip")
	ZipDir(logDirPath, zipFileName)
}

func TestZip(t *testing.T) {
	sourceFilePath := path.Join(configs.GetHomePath(), "logs", "ws.log")
	targetFilePath := path.Join(configs.GetHomePath(), "zipped_logs", "ws.log.zip")
	err := Zip(sourceFilePath, targetFilePath)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetIpv4Address(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{name: "c1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIpv4Address()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIpv4Address() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestIpv42Int(t *testing.T) {
	type args struct {
		ip string
	}
	c1 := args{
		ip: "192.168.7.100",
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "c1", args: c1, want: 3232237412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ipv42Int(tt.args.ip); got != tt.want {
				t.Errorf("Ipv42Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ReadXls(t *testing.T) {
	type args struct {
		xlsFilePath string
	}
	c1 := args{
		xlsFilePath: "/Users/amazing2j/Downloads/MO BOM template-Order header.xls",
	}
	c2 := args{
		xlsFilePath: "/Users/amazing2j/Downloads/MO BOM template.xls",
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{name: "c1", args: c1, wantErr: false},
		{name: "c2", args: c2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadXls(tt.args.xlsFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("readXls() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}
