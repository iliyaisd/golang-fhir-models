// Copyright 2019 - 2021 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// SubscriptionStatusCodes is documented here http://hl7.org/fhir/ValueSet/subscription-status
type SubscriptionStatusCodes int

const (
	SubscriptionStatusCodesRequested SubscriptionStatusCodes = iota
	SubscriptionStatusCodesActive
	SubscriptionStatusCodesError
	SubscriptionStatusCodesOff
)

func (code SubscriptionStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SubscriptionStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "requested":
		*code = SubscriptionStatusCodesRequested
	case "active":
		*code = SubscriptionStatusCodesActive
	case "error":
		*code = SubscriptionStatusCodesError
	case "off":
		*code = SubscriptionStatusCodesOff
	default:
		return fmt.Errorf("unknown SubscriptionStatusCodes code `%s`", s)
	}
	return nil
}
func (code SubscriptionStatusCodes) String() string {
	return code.Code()
}
func (code SubscriptionStatusCodes) Code() string {
	switch code {
	case SubscriptionStatusCodesRequested:
		return "requested"
	case SubscriptionStatusCodesActive:
		return "active"
	case SubscriptionStatusCodesError:
		return "error"
	case SubscriptionStatusCodesOff:
		return "off"
	}
	return "<unknown>"
}
func (code SubscriptionStatusCodes) Display() string {
	switch code {
	case SubscriptionStatusCodesRequested:
		return "Requested"
	case SubscriptionStatusCodesActive:
		return "Active"
	case SubscriptionStatusCodesError:
		return "Error"
	case SubscriptionStatusCodesOff:
		return "Off"
	}
	return "<unknown>"
}
func (code SubscriptionStatusCodes) Definition() string {
	switch code {
	case SubscriptionStatusCodesRequested:
		return "The client has requested the subscription, and the server has not yet set it up."
	case SubscriptionStatusCodesActive:
		return "The subscription is active."
	case SubscriptionStatusCodesError:
		return "The server has an error executing the notification."
	case SubscriptionStatusCodesOff:
		return "Too many errors have occurred or the subscription has expired."
	}
	return "<unknown>"
}
