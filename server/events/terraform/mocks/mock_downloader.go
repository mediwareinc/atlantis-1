// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/terraform (interfaces: Downloader)

package mocks

import (
	go_getter "github.com/hashicorp/go-getter"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockDownloader struct {
	fail func(message string, callerSkip ...int)
}

func NewMockDownloader(options ...pegomock.Option) *MockDownloader {
	mock := &MockDownloader{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockDownloader) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockDownloader) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockDownloader) GetFile(dst string, src string, opts ...go_getter.ClientOption) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockDownloader().")
	}
	params := []pegomock.Param{dst, src}
	for _, param := range opts {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetFile", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockDownloader) GetAny(dst string, src string, opts ...go_getter.ClientOption) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockDownloader().")
	}
	params := []pegomock.Param{dst, src}
	for _, param := range opts {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetAny", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockDownloader) VerifyWasCalledOnce() *VerifierMockDownloader {
	return &VerifierMockDownloader{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockDownloader) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockDownloader {
	return &VerifierMockDownloader{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockDownloader) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockDownloader {
	return &VerifierMockDownloader{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockDownloader) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockDownloader {
	return &VerifierMockDownloader{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockDownloader struct {
	mock                   *MockDownloader
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockDownloader) GetFile(dst string, src string, opts ...go_getter.ClientOption) *MockDownloader_GetFile_OngoingVerification {
	params := []pegomock.Param{dst, src}
	for _, param := range opts {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetFile", params, verifier.timeout)
	return &MockDownloader_GetFile_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockDownloader_GetFile_OngoingVerification struct {
	mock              *MockDownloader
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockDownloader_GetFile_OngoingVerification) GetCapturedArguments() (string, string, []go_getter.ClientOption) {
	dst, src, opts := c.GetAllCapturedArguments()
	return dst[len(dst)-1], src[len(src)-1], opts[len(opts)-1]
}

func (c *MockDownloader_GetFile_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 [][]go_getter.ClientOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([][]go_getter.ClientOption, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param2[u] = make([]go_getter.ClientOption, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(go_getter.ClientOption)
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockDownloader) GetAny(dst string, src string, opts ...go_getter.ClientOption) *MockDownloader_GetAny_OngoingVerification {
	params := []pegomock.Param{dst, src}
	for _, param := range opts {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetAny", params, verifier.timeout)
	return &MockDownloader_GetAny_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockDownloader_GetAny_OngoingVerification struct {
	mock              *MockDownloader
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockDownloader_GetAny_OngoingVerification) GetCapturedArguments() (string, string, []go_getter.ClientOption) {
	dst, src, opts := c.GetAllCapturedArguments()
	return dst[len(dst)-1], src[len(src)-1], opts[len(opts)-1]
}

func (c *MockDownloader_GetAny_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 [][]go_getter.ClientOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([][]go_getter.ClientOption, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param2[u] = make([]go_getter.ClientOption, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(go_getter.ClientOption)
				}
			}
		}
	}
	return
}
