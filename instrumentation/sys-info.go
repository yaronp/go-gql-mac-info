package instrumentation

import (
	"encoding/json"
	"os/exec"
	"sync"
	"time"
)

// system_profiler SPSoftwareDataType SPHardwareDataType SPApplicationsDataType -json > mode.json

type SoftwareDataInfo struct {
	Name            string `json:"_name"`
	BootMode        string `json:"boot_mode"`
	BootVolume      string `json:"boot_volume"`
	KernelVersion   string `json:"kernel_version"`
	LocalHostName   string `json:"local_host_name"`
	OsVersion       string `json:"os_version"`
	SecureVm        string `json:"secure_vm"`
	SystemIntegrity string `json:"system_integrity"`
	Uptime          string `json:"uptime"`
	UserName        string `json:"user_name"`
}

type SoftwareDataType struct {
	SPSoftwareDataType []SoftwareDataInfo `json:"SPSoftwareDataType"`
}

func doSoftwareData() []byte {
	args := []string{"SPSoftwareDataType", "-json"}
	cmd, err := exec.Command("system_profiler", args...).Output()
	if err != nil {
		return nil
	}
	return cmd
}

var doSoftwareDataStub = doSoftwareData

func SoftwareData() *SoftwareDataInfo {
	cmd := doSoftwareDataStub()
	if cmd == nil {
		return nil
	}
	var s SoftwareDataType
	err := json.Unmarshal(cmd, &s)
	if err != nil {
		return nil
	}
	return &s.SPSoftwareDataType[0]
}

type SPApplicationsDataType struct {
	Name         string    `json:"_name"`
	ArchKind     string    `json:"arch_kind"`
	LastModified time.Time `json:"lastModified"`
	ObtainedFrom string    `json:"obtained_from"`
	Path         string    `json:"path"`
	SignedBy     []string  `json:"signed_by,omitempty"`
	Version      string    `json:"version,omitempty"`
	Info         string    `json:"info,omitempty"`
}

type ApplicationsData struct {
	SPApplicationsData []SPApplicationsDataType `json:"SPApplicationsDataType"`
}

func doAppData() []byte {
	args := []string{"SPApplicationsDataType", "-json"}
	cmd, err := exec.Command("system_profiler", args...).Output()
	if err != nil {
		return nil
	}
	return cmd
}

var doAppDataStub = doAppData
var s ApplicationsData
var x sync.Once

func ApplicationData() []SPApplicationsDataType {
	x.Do(func() {
		cmd := doAppDataStub()
		_ = json.Unmarshal(cmd, &s)
	})
	return s.SPApplicationsData
}
