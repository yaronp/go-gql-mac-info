// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ApplicationsData struct {
	Name         *string   `json:"Name"`
	ArchKind     *string   `json:"ArchKind"`
	LastModified *string   `json:"LastModified"`
	ObtainedFrom *string   `json:"ObtainedFrom"`
	Path         *string   `json:"Path"`
	SignedBy     []*string `json:"SignedBy"`
	Version      *string   `json:"Version"`
	Info         *string   `json:"Info"`
}

type SoftwareData struct {
	BootMode        *string `json:"BootMode"`
	BootVolume      *string `json:"BootVolume"`
	KernelVersion   *string `json:"KernelVersion"`
	LocalHostname   *string `json:"LocalHostname"`
	OsVersion       *string `json:"OsVersion"`
	SecureVM        *string `json:"SecureVm"`
	SystemIntegrity *string `json:"SystemIntegrity"`
	Uptime          *string `json:"Uptime"`
	UserName        *string `json:"UserName"`
}
