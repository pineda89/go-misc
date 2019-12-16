package lib_gc_fts_helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func init() {
	H_FtsMessageMaker = &h_ftsmessagemaker{}
}

type Requesttype string

const (
	ACCOMODATION_AVAIL   Requesttype = "ACCOMODATION_AVAIL"   // Avail request
	ACCOMODATION_BOOKING Requesttype = "ACCOMODATION_BOOKING" // Booking request
	CONFIRM              Requesttype = "CONFIRM"              // Confirm request
	CANCEL               Requesttype = "CANCEL"               // Cancel request
)

var FTSOperation map[Requesttype]string = map[Requesttype]string{
	ACCOMODATION_AVAIL:   "AVAIL",
	ACCOMODATION_BOOKING: "BOOKING",
	CONFIRM:              "CONFIRM",
	CANCEL:               "CANCEL",
}

const (
	AVAILREQUESTID = iota // 0
	REQUESTXML            // 1
	RESPONSEXML           // 2
	PRELOCATOR            // 3
	LOCATOR               // 4
	CUSTOMERID            // 5
	RESPONSESTATUS        // 6
	SUPPLIERID            // 7
	LINKREQUESTID         // 8

	RTYPE       // 9
	ATTEMPS     // 10
	TARGETQUEUE // 11
	TOPIC       // 12
)

var H_FtsMessageMaker H_FtsMessageMaker_I

type AuditTrack struct {
	Attempts            int    `json:"attempts"`       // Always set to zero
	Application         string `json:"application"`    // Application to be tracked. This name is unique for app
	TargetQueue         string `json:"targetQueue"`    // It will be 'auditXmlQueue' always
	AvailRequestId      string `json:"availRequestId"` // Booking flux unique ID
	PreLocator          string `json:"preLocator"`     // Booking action locator
	Locator             string `json:"locator"`        // Confirm action locator
	LinkRequestId       string `json:"linkRequestId"`
	Type                string `json:"type"`                // AVAIL,BOOKING,CONFIRM,CANCEL segun el caso
	Origin              string `json:"origin"`              // INTEGRATION
	CustomerId          string `json:"customerId"`          // ClientId
	SupplierId          string `json:"supplierId"`          // Supplier ID
	XmlRequest          string `json:"xmlRequest"`          // Leave empty
	XmlResponse         string `json:"xmlResponse"`         // Leave empty
	ExternalXmlRequest  string `json:"externalXmlRequest"`  // Request XML serialized as string
	ExternalXmlResponse string `json:"externalXmlResponse"` // Response XML serialized as string
	ResponseStatus      string `json:"responseStatus"`      // Operation result status element value (0,0009, ...)
	DateTime            string `json:"dateTime"`            // Timestamp
	Backup              string `json:"backup"`              // Always set to false
	External            string `json:"external"`            // Always set to false
}

type FTSTrackMessage struct {
	Rtype          Requesttype
	Attempts       int
	Application    string
	Targetqueue    string
	Topic          string
	RequestXML     string
	ResponseXML    string
	AvailRequestId string
	PreLocator     string
	Locator        string
	CustomerId     string
	ResponseStatus string
	LinkRequestId  string
	SupplierID     string
}

type H_FtsMessageMaker_I interface {
	MakeFTSTrackMessage(track *FTSTrackMessage) ([]byte, error)
}

type h_ftsmessagemaker struct{}

func (hftsmm *h_ftsmessagemaker) MakeFTSTrackMessage(track *FTSTrackMessage) ([]byte, error) {

	message := new(AuditTrack)
	message.Attempts = track.Attempts
	message.Application = track.Application
	message.TargetQueue = track.Targetqueue
	message.AvailRequestId = track.AvailRequestId
	message.PreLocator = track.PreLocator
	message.Locator = track.Locator
	message.LinkRequestId = track.LinkRequestId

	if t, ok := FTSOperation[track.Rtype]; !ok {
		return nil, errors.New(fmt.Sprintf("The operation %s is not supported", track.Rtype))
	} else {
		message.Type = t
	}

	message.Origin = "INTEGRATION"
	message.CustomerId = track.CustomerId
	message.SupplierId = track.SupplierID

	message.ResponseStatus = track.ResponseStatus

	message.ExternalXmlRequest = track.RequestXML
	message.ExternalXmlResponse = track.ResponseXML

	message.DateTime = strconv.FormatInt(time.Now().UnixNano()/1e6, 10) //Format("20060102150405") //timestamp del momento de tiempo actual
	message.Backup = "false"
	message.External = "false"

	// Se da formato al mensaje en formato JSON
	if messageJSON, err := json.Marshal(&message); err != nil {
		return nil, err
	} else {
		return messageJSON, nil
	}
}
