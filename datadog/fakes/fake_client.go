// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/blackbox/datadog"
)

type FakeClient struct {
	PublishSeriesStub        func(series datadog.Series) error
	publishSeriesMutex       sync.RWMutex
	publishSeriesArgsForCall []struct {
		series datadog.Series
	}
	publishSeriesReturns struct {
		result1 error
	}
}

func (fake *FakeClient) PublishSeries(series datadog.Series) error {
	fake.publishSeriesMutex.Lock()
	fake.publishSeriesArgsForCall = append(fake.publishSeriesArgsForCall, struct {
		series datadog.Series
	}{series})
	fake.publishSeriesMutex.Unlock()
	if fake.PublishSeriesStub != nil {
		return fake.PublishSeriesStub(series)
	} else {
		return fake.publishSeriesReturns.result1
	}
}

func (fake *FakeClient) PublishSeriesCallCount() int {
	fake.publishSeriesMutex.RLock()
	defer fake.publishSeriesMutex.RUnlock()
	return len(fake.publishSeriesArgsForCall)
}

func (fake *FakeClient) PublishSeriesArgsForCall(i int) datadog.Series {
	fake.publishSeriesMutex.RLock()
	defer fake.publishSeriesMutex.RUnlock()
	return fake.publishSeriesArgsForCall[i].series
}

func (fake *FakeClient) PublishSeriesReturns(result1 error) {
	fake.PublishSeriesStub = nil
	fake.publishSeriesReturns = struct {
		result1 error
	}{result1}
}

var _ datadog.Client = new(FakeClient)