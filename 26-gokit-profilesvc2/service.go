/*
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/
package main

import (
	"errors"
	"strings"
)

type StringService interface {
	UpperCase(string) (string, error)
	LowerCase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) UpperCase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) LowerCase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToLower(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

var ErrEmpty = errors.New("Empty String Error")
