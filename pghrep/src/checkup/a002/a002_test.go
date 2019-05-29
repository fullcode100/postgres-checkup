package a002

import (
	"fmt"
	"testing"

	checkup ".."
)

func TestGetMajorMinorVersion(t *testing.T) {
	major, minor := getMajorMinorVersion("110003")
	if major != "11" || minor != "3" {
		t.Fatal("TestGetMajorMinorVersion failed")
	}
	major, minor = getMajorMinorVersion("90612")
	if major != "9.6" || minor != "12" {
		t.Fatal("TestGetMajorMinorVersion failed")
	}
}

func TestA002Sucess(t *testing.T) {
	fmt.Println(t.Name())
	var report A002Report
	var hostResult A002ReportHostResult
	hostResult.Data = A002ReportHostResultData{
		Version:          "PostgreSQL 11.3 on x86_64-pc-linux-gnu (Ubuntu 11.22-1.pgdg16.04+1), compiled by gcc (Ubuntu 5.4.0-6ubuntu1~16.04.10) 5.4.0 20160609, 64-bit",
		ServerVersionNum: "110003",
		ServerMajorVer:   "11",
		ServerMinorVer:   "3",
	}
	report.Results = A002ReportHostsResults{"test-host": hostResult}
	result := A002Process(report)
	if result.P1 || result.P2 || result.P3 {
		t.Fatal("TestA002Sucess failed")
	}
	checkup.PrintConclusions(result)
	checkup.PrintReccomendations(result)
}

func TestA002IsSame(t *testing.T) {
	fmt.Println(t.Name())
	var report A002Report
	var host1Result A002ReportHostResult
	var host2Result A002ReportHostResult
	host1Result.Data = A002ReportHostResultData{
		Version:          "PostgreSQL 11.3 on x86_64-pc-linux-gnu (Ubuntu 11.3-1.pgdg16.04+1), compiled by gcc (Ubuntu 5.4.0-6ubuntu1~16.04.10) 5.4.0 20160609, 64-bit",
		ServerVersionNum: "110003",
		ServerMajorVer:   "11",
		ServerMinorVer:   "3",
	}
	host2Result.Data = A002ReportHostResultData{
		Version:          "PostgreSQL 11.3 on x86_64-pc-linux-gnu (Ubuntu 11.3-1.pgdg16.04+1), compiled by gcc (Ubuntu 5.4.0-6ubuntu1~16.04.10) 5.4.0 20160609, 64-bit",
		ServerVersionNum: "110003",
		ServerMajorVer:   "11",
		ServerMinorVer:   "3",
	}
	report.Results = A002ReportHostsResults{"host1": host1Result, "host2": host2Result}
	result := A002Process(report)
	if result.P1 || result.P2 || result.P3 {
		t.Fatal("TestA002IsSame failed")
	}
	checkup.PrintConclusions(result)
	checkup.PrintReccomendations(result)
}

func TestA002IsNotSame(t *testing.T) {
	fmt.Println(t.Name())
	var report A002Report
	var host1Result A002ReportHostResult
	var host2Result A002ReportHostResult
	host1Result.Data = A002ReportHostResultData{
		Version:          "PostgreSQL 9.6.12 on x86_64-pc-linux-gnu (Ubuntu 9.6.11-1.pgdg16.04+1), compiled by gcc (Ubuntu 5.4.0-6ubuntu1~16.04.10) 5.4.0 20160609, 64-bit",
		ServerVersionNum: "90612",
		ServerMajorVer:   "9.6",
		ServerMinorVer:   "12",
	}
	host2Result.Data = A002ReportHostResultData{
		Version:          "PostgreSQL 9.6.11 on x86_64-pc-linux-gnu (Ubuntu 9.6.11-1.pgdg16.04+1), compiled by gcc (Ubuntu 5.4.0-6ubuntu1~16.04.10) 5.4.0 20160609, 64-bit",
		ServerVersionNum: "90611",
		ServerMajorVer:   "9.6",
		ServerMinorVer:   "11",
	}
	report.Results = A002ReportHostsResults{"host1": host1Result, "host2": host2Result}
	result := A002Process(report)
	if !result.P2 {
		t.Fatal("TestA002IsNotSame failed")
	}
	checkup.PrintConclusions(result)
	checkup.PrintReccomendations(result)
}

func TestA002WrongVersion(t *testing.T) {
	fmt.Println(t.Name())
	var report A002Report
	var hostResult A002ReportHostResult
	hostResult.Data = A002ReportHostResultData{
		Version:          "PostgreSQL 99.99.22 on x86_64-pc-linux-gnu (Ubuntu 9.6.11-1.pgdg16.04+1), compiled by gcc (Ubuntu 5.4.0-6ubuntu1~16.04.10) 5.4.0 20160609, 64-bit",
		ServerVersionNum: "990099",
		ServerMajorVer:   "99",
		ServerMinorVer:   "99",
	}
	report.Results = A002ReportHostsResults{"test-host": hostResult}
	result := A002Process(report)
	if !result.P1 {
		t.Fatal("TestA002WrongVersion failed")
	}
	checkup.PrintConclusions(result)
	checkup.PrintReccomendations(result)
}

func TestA002LatestMajor(t *testing.T) {
	fmt.Println(t.Name())
	var report A002Report
	var hostResult A002ReportHostResult
	hostResult.Data = A002ReportHostResultData{
		Version:          "PostgreSQL 9.6.22 on x86_64-pc-linux-gnu (Ubuntu 9.6.22-1.pgdg16.04+1), compiled by gcc (Ubuntu 5.4.0-6ubuntu1~16.04.10) 5.4.0 20160609, 64-bit",
		ServerVersionNum: "90622",
		ServerMajorVer:   "9.6",
		ServerMinorVer:   "22",
	}
	report.Results = A002ReportHostsResults{"test-host": hostResult}
	result := A002Process(report)
	if !result.P3 {
		t.Fatal("TestA002LatestMajor failed")
	}
	checkup.PrintConclusions(result)
	checkup.PrintReccomendations(result)
}
