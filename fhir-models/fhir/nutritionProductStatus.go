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

// NutritionProductStatus is documented here http://hl7.org/fhir/ValueSet/nutritionproduct-status
type NutritionProductStatus int

const (
	NutritionProductStatusActive NutritionProductStatus = iota
	NutritionProductStatusInactive
	NutritionProductStatusEnteredInError
)

func (code NutritionProductStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *NutritionProductStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = NutritionProductStatusActive
	case "inactive":
		*code = NutritionProductStatusInactive
	case "entered-in-error":
		*code = NutritionProductStatusEnteredInError
	default:
		return fmt.Errorf("unknown NutritionProductStatus code `%s`", s)
	}
	return nil
}
func (code NutritionProductStatus) String() string {
	return code.Code()
}
func (code NutritionProductStatus) Code() string {
	switch code {
	case NutritionProductStatusActive:
		return "active"
	case NutritionProductStatusInactive:
		return "inactive"
	case NutritionProductStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code NutritionProductStatus) Display() string {
	switch code {
	case NutritionProductStatusActive:
		return "Active"
	case NutritionProductStatusInactive:
		return "Inactive"
	case NutritionProductStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code NutritionProductStatus) Definition() string {
	switch code {
	case NutritionProductStatusActive:
		return "The product can be used."
	case NutritionProductStatusInactive:
		return "The product is not expected or allowed to be used."
	case NutritionProductStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	}
	return "<unknown>"
}