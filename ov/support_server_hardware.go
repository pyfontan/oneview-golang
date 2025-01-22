/*
(c) Copyright [2021] Hewlett Packard Enterprise Development LP

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package ov -
package ov

import (
	"encoding/json"
	"time"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type SupportServerHardware struct {
	AcquiredProductNumber              string             `json:"acquiredProductNumber"`
	AcquiredSerialNumber               string             `json:"acquiredSerialNumber"`
	Category                           string             `json:"category"`
	Created                            string             `json:"created"`
	CustomDeliveryID                   string             `json:"customDeliveryId"`
	DataCollections                    DataCollections    `json:"dataCollections"`
	DataCollectionURI                  utils.Nstring      `json:"dataCollectionUri"`
	EnteredProductNumber               string             `json:"enteredProductNumber"`
	EnteredSerialNumber                string             `json:"enteredSerialNumber"`
	EnteredObligationID                string             `json:"enteredObligationId"`
	EnteredObligationType              string             `json:"enteredObligationType"`
	Entitlement                        Entitlement        `json:"entitlement"`
	EntitlementDetails                 EntitlementDetails `json:"entitlementDetails"`
	EntitlementURI                     string             `json:"entitlementUri"`
	ETag                               string             `json:"eTag"`
	Modified                           string             `json:"modified"`
	PrimaryContact                     Contact            `json:"primaryContact"`
	PrimaryContactInheritedFrom        string             `json:"primaryContactInheritedFrom"`
	PrimaryContactURI                  utils.Nstring      `json:"primaryContactUri"`
	RsdcEnabled                        bool               `json:"rsdcEnabled"`
	SalesChannelPartner                ChannelPartner     `json:"salesChannelPartner"`
	SecondaryContact                   Contact            `json:"secondaryContact"`
	SecondaryContactInheritedFrom      string             `json:"secondaryContactInheritedFrom"`
	SecondaryContactURI                utils.Nstring      `json:"secondaryContactUri"`
	SupportEnabled                     bool               `json:"supportEnabled"`
	SupportEnabledPreference           string             `json:"supportEnabledPreference"`
	SupportChannelPartner              ChannelPartner     `json:"supportChannelPartner"`
	SupportChannelPartnerInheritedFrom string             `json:"supportChannelPartnerInheritedFrom"`
	SupportChannelPartnerURI           string             `json:"supportChannelPartnerUri"`
	SupportsTestEvents                 bool               `json:"supportsTestEvents"`
	Type                               string             `json:"type"`
	URI                                string             `json:"uri"`

	Client *OVClient `json:"-"`
}

type ChannelPartner struct {
	Category      string `json:"category"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Created       string `json:"created"`
	Default       bool   `json:"default"`
	ETag          string `json:"eTag"`
	ID            string `json:"id"`
	InheritedFrom string `json:"inheritedFrom"`
	Key           string `json:"key"`
	Modified      string `json:"modified"`
	Name          string `json:"name"`
	PartnerType   string `json:"partnerType"`
	Type          string `json:"type"`
	URI           string `json:"uri"`
}

type CollectionMember struct {
	Category       string        `json:"category"`
	CollectionDate string        `json:"collectionDate"`
	CollectionKey  string        `json:"collectionKey"`
	CollectionType string        `json:"collectionType"`
	Created        string        `json:"created"`
	DownloadURI    string        `json:"downloadUri"`
	ETag           string        `json:"eTag"`
	Modified       string        `json:"modified"`
	Status         string        `json:"status"`
	SubmitGUID     string        `json:"submitGuid"`
	Type           string        `json:"type"`
	URI            utils.Nstring `json:"uri"`
}

type Contact struct {
	AdditionalEmails string `json:"additionalEmails"`
	AlternatePhone   string `json:"alternatePhone"`
	Category         string `json:"category"`
	ContactKey       string `json:"contactKey"`
	Created          string `json:"created"`
	Default          bool   `json:"default"`
	DefaultSecondary bool   `json:"defaultSecondary"`
	Email            string `json:"email"`
	ETag             string `json:"eTag"`
	FirstName        string `json:"firstName"`
	InheritedFrom    string `json:"inheritedFrom"`
	Language         string `json:"language"`
	LastName         string `json:"lastName"`
	Modified         string `json:"modified"`
	Notes            string `json:"notes"`
	OosCount         int64  `json:"oosCount"`
	Primary          string `json:"primary"`
	PrimaryPhone     string `json:"primaryPhone"`
	Type             string `json:"type"`
	URI              string `json:"uri"`
}

type DataCollections struct {
	Count       int64              `json:"count"`
	Category    string             `json:"category"`
	Created     string             `json:"created"`
	ETag        string             `json:"eTag"`
	Members     []CollectionMember `json:"members"`
	Modified    string             `json:"modified"`
	NextPageURI utils.Nstring      `json:"nextPageUri"`
	Start       int64              `json:"start"`
	Total       int64              `json:"total"`
	Type        string             `json:"type"`
	URI         utils.Nstring      `json:"uri"`
}

type Entitlement struct {
	Category       string        `json:"category"`
	Created        string        `json:"created"`
	ETag           string        `json:"eTag"`
	Modified       string        `json:"modified"`
	ObligationType string        `json:"obligationType"`
	ObligationID   string        `json:"obligationId"`
	Type           string        `json:"type"`
	URI            utils.Nstring `json:"uri"`
}

type EntitlementDetails struct {
	Category             string  `json:"category,omitempty"`
	CountryCode          string  `json:"countryCode,omitempty"`
	CoverageDays         string  `json:"coverageDays,omitempty"`
	CoverageHoursDay1to5 string  `json:"coverageHoursDay1to5,omitempty"` // "coverageHoursDay1to5": "HoursADay12"
	CoverageHoursDay6    string  `json:"coverageHoursDay6,omitempty"`    // "coverageHoursDay6": "HoursADay24"
	CoverageHoursDay7    string  `json:"coverageHoursDay7,omitempty"`    // "coverageHoursDay7": "StandardOfficeHours"
	CoversHolidays       bool    `json:"coversHoliday,omitempty"`
	Created              string  `json:"created,omitempty"`
	EntitlementKey       string  `json:"entitlementKey"`
	EntitlementPackage   string  `json:"entitlementPackage"`
	EntitlementStatus    string  `json:"entitlementStatus,omitempty"`
	ETag                 string  `json:"eTag"`
	Explanation          string  `json:"explanation"`
	IsEntitled           bool    `json:"isEntitled"`
	LastUpd              string  `json:"lastUpd"`
	Modified             string  `json:"modified"`
	ObligationEndDate    string  `json:"obligationEndDate"`
	ObligationID         string  `json:"obligationId"`
	ObligationStartDate  string  `json:"obligationStartDate"`
	ObligationType       string  `json:"obligationType"`
	OfferEndDate         string  `json:"offerEndDate,omitempty"`
	OfferStartDate       string  `json:"offerStartDate,omitempty"`
	OfferStatus          string  `json:"offerStatus,omitempty"`
	PrimaryContact       Contact `json:"primaryContact,omitempty"`
	RefreshState         string  `json:"refreshState"`
	ResponseHolidays     string  `json:"responseHolidays,omitempty"`
	ResponseTimeDay1to5  string  `json:"responseTimeDay1to5,omitempty"`
	ResponseTimeDay6     string  `json:"responseTimeDay6,omitempty"`
	ResponseTimeDay7     string  `json:"responseTimeDay7,omitempty"`
	ResponseTimeHolidays string  `json:"responseTimeHolidays"`
	SecondaryContact     Contact `json:"secondaryContact,omitempty"`
	Type                 string  `json:"type"`
	URI                  string  `json:"uri"`
}

func OvDate(dateString string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateString)
}

func (ssh SupportServerHardware) GetOfferStartDate() (time.Time, error) {
	return OvDate(ssh.EntitlementDetails.OfferStartDate)
}

func (ssh SupportServerHardware) GetOfferEndDate() (time.Time, error) {
	return OvDate(ssh.EntitlementDetails.OfferEndDate)
}

func (c *OVClient) GetSupportServerHardware(id string) (SupportServerHardware, error) {
	var (
		uri                   = "/rest/support/server-hardware/" + id
		supportServerHardware SupportServerHardware
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil, nil)
	if err != nil {
		return supportServerHardware, err
	}

	log.Debugf("GetSupportServerHardware %s", data)

	if err := json.Unmarshal([]byte(data), &supportServerHardware); err != nil {
		return supportServerHardware, err
	}

	supportServerHardware.Client = c

	return supportServerHardware, nil

}
