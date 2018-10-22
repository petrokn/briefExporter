package configuration

import (
	"io/ioutil"
	"encoding/json"
	"briefExporter/common"
)

type Config struct {
	NotesRetrieveUrl           string `json:"retrieve_url"`
	NotesSendUrl               string `json:"send_url"`
	LibraryCheckUrl            string `json:"library_check_url"`
	LibrarySyncUrl             string `json:"library_sync_url"`
	ScanFolder                 string `json:"scan_folder"`
	ScanMountPathScript        string `json:"scan_mount_path_script"`
	DeviceAvailabilityUrl	   string `json:"device_availability_url"`
	TokenRetrieveUrl 		   string `json:"token_retrieve_url"`
}

func GetConfig(path string) (*Config, error) {
	var config *Config

	data, err := ioutil.ReadFile(path)
	common.Check(err)

	err = json.Unmarshal(data, &config)

	return config, err
}