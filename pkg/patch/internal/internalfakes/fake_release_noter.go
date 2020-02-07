/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"sync"

	"k8s.io/release/pkg/patch"
)

type FakeReleaseNoter struct {
	GetMarkdownStub        func() (string, error)
	getMarkdownMutex       sync.RWMutex
	getMarkdownArgsForCall []struct {
	}
	getMarkdownReturns struct {
		result1 string
		result2 error
	}
	getMarkdownReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseNoter) GetMarkdown() (string, error) {
	fake.getMarkdownMutex.Lock()
	ret, specificReturn := fake.getMarkdownReturnsOnCall[len(fake.getMarkdownArgsForCall)]
	fake.getMarkdownArgsForCall = append(fake.getMarkdownArgsForCall, struct {
	}{})
	fake.recordInvocation("GetMarkdown", []interface{}{})
	fake.getMarkdownMutex.Unlock()
	if fake.GetMarkdownStub != nil {
		return fake.GetMarkdownStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getMarkdownReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReleaseNoter) GetMarkdownCallCount() int {
	fake.getMarkdownMutex.RLock()
	defer fake.getMarkdownMutex.RUnlock()
	return len(fake.getMarkdownArgsForCall)
}

func (fake *FakeReleaseNoter) GetMarkdownCalls(stub func() (string, error)) {
	fake.getMarkdownMutex.Lock()
	defer fake.getMarkdownMutex.Unlock()
	fake.GetMarkdownStub = stub
}

func (fake *FakeReleaseNoter) GetMarkdownReturns(result1 string, result2 error) {
	fake.getMarkdownMutex.Lock()
	defer fake.getMarkdownMutex.Unlock()
	fake.GetMarkdownStub = nil
	fake.getMarkdownReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseNoter) GetMarkdownReturnsOnCall(i int, result1 string, result2 error) {
	fake.getMarkdownMutex.Lock()
	defer fake.getMarkdownMutex.Unlock()
	fake.GetMarkdownStub = nil
	if fake.getMarkdownReturnsOnCall == nil {
		fake.getMarkdownReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getMarkdownReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseNoter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMarkdownMutex.RLock()
	defer fake.getMarkdownMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReleaseNoter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ patch.ReleaseNoter = new(FakeReleaseNoter)