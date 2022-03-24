package instrumentation

import "testing"

const outputData = `{"SPSoftwareDataType" : [{
	"_name" : "os_overview",
	"boot_mode" : "normal_boot",
	"boot_volume" : "Macintosh HD",
	"kernel_version" : "Darwin 21.4.0",
	"local_host_name" : "C02FD3J4MD6T",
	"os_version" : "macOS 12.3 (21E230)",
	"secure_vm" : "secure_vm_enabled",
	"system_integrity" : "integrity_enabled",
	"uptime" : "up 7:21:56:33",
	"user_name" : "Yaron Pdut (i355599)"
}]}`

var outputDataStruct = SoftwareDataInfo{
	Name:            "os_overview",
	BootMode:        "normal_boot",
	BootVolume:      "Macintosh HD",
	KernelVersion:   "Darwin 21.4.0",
	LocalHostName:   "C02FD3J4MD6T",
	OsVersion:       "macOS 12.3 (21E230)",
	SecureVm:        "secure_vm_enabled",
	SystemIntegrity: "integrity_enabled",
	Uptime:          "up 7:21:56:33",
	UserName:        "Yaron Pdut (i355599)",
}

func TestSystemProfiler(t *testing.T) {
	doSoftwareDataStub = func() []byte { return []byte(outputData) }
	r := SoftwareData()
	if *r != outputDataStruct {
		t.Error("output is not the same")
	}
}
