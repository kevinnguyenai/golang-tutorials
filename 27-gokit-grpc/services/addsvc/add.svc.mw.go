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

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

// Middleware describes a service (as apposed to endpoint) middleware
type AddSvcMiddleware func(Service) Service

// LoggingSvcMiddleware takes a logger as a dependence
// and returns a service Middleware
func LoggingAddSvcMiddleware(logger log.Logger) AddSvcMiddleware {
	return func(next Service) Service {
		return loggingAddSvcMiddleware{logger, next}
	}
}

type loggingAddSvcMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingAddSvcMiddleware) Sum(ctx context.Context, a, b int) (v int, err error) {
	defer func() {
		mw.logger.Log("method", "Sum", "a", a, "b", b, "v", v, "err", err)
	}()
	return mw.next.Sum(ctx, a, b)
}

func (mw loggingAddSvcMiddleware) Concat(ctx context.Context, a, b string) (v string, err error) {
	defer func() {
		mw.logger.Log("method", "Concat", "a", a, "b", b, "v", v, "err", err)
	}()
	return mw.next.Concat(ctx, a, b)
}

func InstrumentingAddSvcMiddleware(ints, chars metrics.Counter) AddSvcMiddleware {
	return func(next Service) Service {
		return instrumentingAddSvcMiddleware{
			ints:  ints,
			chars: chars,
			next:  next,
		}
	}
}

type instrumentingAddSvcMiddleware struct {
	ints  metrics.Counter
	chars metrics.Counter
	next  Service
}

func (mw instrumentingAddSvcMiddleware) Sum(ctx context.Context, a, b int) (v int, err error) {
	v, err = mw.next.Sum(ctx, a, b)
	mw.ints.Add(float64(v))
	return v, err
}

func (mw instrumentingAddSvcMiddleware) Concat(ctx context.Context, a, b string) (v string, err error) {
	v, err = mw.next.Concat(ctx, a, b)
	mw.chars.Add(float64(len(v)))
	return v, err
}
