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
package addsvc

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

// Service describes a service that add two thing together
type Service interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

// New returns a basic Service with all of expected middlewares wired in
func NewService(logger log.Logger, ints, chars metrics.Counter) Service {
	var svc Service
	{
		svc = NewBasicService()
		svc = LoggingAddSvcMiddleware(logger)(svc)
		svc = InstrumentingAddSvcMiddleware(ints, chars)(svc)
	}
	return svc
}

var (
	// ErrTwoZeros is an abitrary business rule for Add method
	ErrTwoZeroes = errors.New("can't sum two zeroes")
	// ErrIntOverFlow protects the Add method. We've decided that
	// this error indicates a misbehaving service and should count againts e.g. circuit
	// breakers. So , we return it directly in endpoints, to ilusstrate the
	// difference. In a real service, this probably wouldn't be the case
	ErrIntOverFlow = errors.New("integer overflow")

	// ErrMaxSizeExceed protect the concat method
	ErrMaxSizeExceed = errors.New("result exceed maximum size")
)

func NewBasicService() Service {
	return basicService{}
}

type basicService struct{}

const (
	intMax = 1<<31 - 1
	intMin = -(intMax + 1)
	maxLen = 10
)

func (s basicService) Sum(_ context.Context, a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrTwoZeroes
	}
	if (b > 0 && a > (intMax-b)) || (b < 0 && a < (intMax-b)) {
		return 0, ErrIntOverFlow
	}
	return a + b, nil
}

func (s basicService) Concat(_ context.Context, a, b string) (string, error) {
	if len(a)+len(b) > maxLen {
		return "", ErrMaxSizeExceed
	}
	return a + b, nil
}
