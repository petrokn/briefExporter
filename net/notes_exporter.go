package net

import (
	"briefExporter/common"
	"briefExporter/configuration"
	"briefExporter/ui"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type HistoryRecord struct {
	SerialNumber string    `json:"serial_number"`
	Checksum     [16]byte  `json:"checksum"`
	CreatedOn    time.Time `json:"created_on"`
}

func GetPreviousHistoryRecord(serial string, config *configuration.Config, token *string) (*HistoryRecord, error) {

	resp, err := executeRequest(config.NotesRetrieveUrl+"/"+serial, "GET", nil,
		GetAuthorizationHeaders(nil, token))

	var historyRecord *HistoryRecord

	body, _ := ioutil.ReadAll(resp.Body)
	logResponse(resp, body)

	err = json.Unmarshal(body, &historyRecord)

	return historyRecord, err
}

func SendNotesToServer(notes *[]byte, config *configuration.Config, token *string) {
	headers := make(map[string]string)
	headers["Set-Type"] = "All"
	headers["Content-Type"] = "application/json"

	resp, err := executeRequest(config.NotesSendUrl, "POST", bytes.NewBuffer(*notes),
		GetAuthorizationHeaders(headers, token))
	common.Check(err)

	body, _ := ioutil.ReadAll(resp.Body)

	logResponse(resp, body)
}

func CheckDeviceAvailability(device *ui.Device, config *configuration.Config, token *string) bool {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	deviceJson, err := json.Marshal(*device)
	common.Check(err)

	resp, err := executeRequest(config.DeviceAvailabilityUrl, "GET", bytes.NewBuffer(deviceJson),
		GetAuthorizationHeaders(nil, token))
	common.Check(err)

	if resp.StatusCode == http.StatusOK {
		return true
	}

	return false
}

func CreateUserDevice(deviceId *string, deviceName *string, deviceSerial *string,
	config *configuration.Config, token *string) bool {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	userDevice := map[string]interface{}{"DeviceId": *deviceId, "DeviceName": *deviceName, "DeviceSerial": *deviceSerial}

	userDeviceJson, err := json.Marshal(userDevice)
	common.Check(err)

	resp, err := executeRequest(config.CreateUserDeviceUrl, "POST", bytes.NewBuffer(userDeviceJson),
		GetAuthorizationHeaders(headers, token))
	common.Check(err)

	if resp.StatusCode == http.StatusCreated {
		return true
	}

	return false
}