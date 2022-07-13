package mocks

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "Repository" can be found in github.com/P2PCloud/miner-search-api/internal/app/test/mocks
*/
import (
	context "context"
	"sync/atomic"
	"time"

	models "github.com/P2PCloud/miner-search-api/internal/app/models"
	"github.com/gojuno/minimock"

	testify_assert "github.com/stretchr/testify/assert"
)

// RepositoryMock implements github.com/P2PCloud/miner-search-api/internal/app/test/mocks.Repository
type RepositoryMock struct {
	t minimock.Tester

	CompareOfferFunc       func(p *models.Offer, p1 *models.VmType) (r bool)
	CompareOfferCounter    uint64
	CompareOfferPreCounter uint64
	CompareOfferMock       mRepositoryMockCompareOffer

	GetMinerReputationByIdFunc       func(p string) (r float32)
	GetMinerReputationByIdCounter    uint64
	GetMinerReputationByIdPreCounter uint64
	GetMinerReputationByIdMock       mRepositoryMockGetMinerReputationById

	GetOffersByParamsFunc       func(p context.Context, p1 *models.VmType) (r models.OfferCollection, r1 error)
	GetOffersByParamsCounter    uint64
	GetOffersByParamsPreCounter uint64
	GetOffersByParamsMock       mRepositoryMockGetOffersByParams

	GetVmTypeByIdFunc       func(p string) (r *models.VmType)
	GetVmTypeByIdCounter    uint64
	GetVmTypeByIdPreCounter uint64
	GetVmTypeByIdMock       mRepositoryMockGetVmTypeById
}

// NewRepositoryMock returns a mock for github.com/P2PCloud/miner-search-api/internal/app/test/mocks.Repository
func NewRepositoryMock(t minimock.Tester) *RepositoryMock {
	m := &RepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CompareOfferMock = mRepositoryMockCompareOffer{mock: m}
	m.GetMinerReputationByIdMock = mRepositoryMockGetMinerReputationById{mock: m}
	m.GetOffersByParamsMock = mRepositoryMockGetOffersByParams{mock: m}
	m.GetVmTypeByIdMock = mRepositoryMockGetVmTypeById{mock: m}

	return m
}

type mRepositoryMockCompareOffer struct {
	mock              *RepositoryMock
	mainExpectation   *RepositoryMockCompareOfferExpectation
	expectationSeries []*RepositoryMockCompareOfferExpectation
}

// RepositoryMockCompareOfferExpectation specifies expectation struct of the Repository.CompareOffer
type RepositoryMockCompareOfferExpectation struct {
	input  *RepositoryMockCompareOfferInput
	result *RepositoryMockCompareOfferResult
}

// RepositoryMockCompareOfferInput represents input parameters of the Repository.CompareOffer
type RepositoryMockCompareOfferInput struct {
	p  *models.Offer
	p1 *models.VmType
}

// RepositoryMockCompareOfferResult represents results of the Repository.CompareOffer
type RepositoryMockCompareOfferResult struct {
	r bool
}

// Expect specifies that invocation of Repository.CompareOffer is expected from 1 to Infinity times
func (m *mRepositoryMockCompareOffer) Expect(p *models.Offer, p1 *models.VmType) *mRepositoryMockCompareOffer {
	m.mock.CompareOfferFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RepositoryMockCompareOfferExpectation{}
	}
	m.mainExpectation.input = &RepositoryMockCompareOfferInput{p, p1}
	return m
}

// Return specifies results of invocation of Repository.CompareOffer
func (m *mRepositoryMockCompareOffer) Return(r bool) *RepositoryMock {
	m.mock.CompareOfferFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RepositoryMockCompareOfferExpectation{}
	}
	m.mainExpectation.result = &RepositoryMockCompareOfferResult{r}
	return m.mock
}

// ExpectOnce specifies that invocation of Repository.CompareOffer is expected once
func (m *mRepositoryMockCompareOffer) ExpectOnce(p *models.Offer, p1 *models.VmType) *RepositoryMockCompareOfferExpectation {
	m.mock.CompareOfferFunc = nil
	m.mainExpectation = nil

	expectation := &RepositoryMockCompareOfferExpectation{}
	expectation.input = &RepositoryMockCompareOfferInput{p, p1}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

// Return sets up return arguments of expectation struct for Repository.CompareOffer
func (e *RepositoryMockCompareOfferExpectation) Return(r bool) {
	e.result = &RepositoryMockCompareOfferResult{r}
}

// Set uses given function f as a mock of Repository.CompareOffer method
func (m *mRepositoryMockCompareOffer) Set(f func(p *models.Offer, p1 *models.VmType) (r bool)) *RepositoryMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.CompareOfferFunc = f
	return m.mock
}

// CompareOffer implements github.com/P2PCloud/miner-search-api/internal/app/test/mocks.Repository interface
func (m *RepositoryMock) CompareOffer(p *models.Offer, p1 *models.VmType) (r bool) {
	counter := atomic.AddUint64(&m.CompareOfferPreCounter, 1)
	defer atomic.AddUint64(&m.CompareOfferCounter, 1)

	if len(m.CompareOfferMock.expectationSeries) > 0 {
		if counter > uint64(len(m.CompareOfferMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to RepositoryMock.CompareOffer. %v %v", p, p1)
			return
		}

		input := m.CompareOfferMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, RepositoryMockCompareOfferInput{p, p1}, "Repository.CompareOffer got unexpected parameters")

		result := m.CompareOfferMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the RepositoryMock.CompareOffer")
			return
		}

		r = result.r

		return
	}

	if m.CompareOfferMock.mainExpectation != nil {

		input := m.CompareOfferMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, RepositoryMockCompareOfferInput{p, p1}, "Repository.CompareOffer got unexpected parameters")
		}

		result := m.CompareOfferMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the RepositoryMock.CompareOffer")
		}

		r = result.r

		return
	}

	if m.CompareOfferFunc == nil {
		m.t.Fatalf("Unexpected call to RepositoryMock.CompareOffer. %v %v", p, p1)
		return
	}

	return m.CompareOfferFunc(p, p1)
}

// CompareOfferMinimockCounter returns a count of RepositoryMock.CompareOfferFunc invocations
func (m *RepositoryMock) CompareOfferMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CompareOfferCounter)
}

// CompareOfferMinimockPreCounter returns the value of RepositoryMock.CompareOffer invocations
func (m *RepositoryMock) CompareOfferMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.CompareOfferPreCounter)
}

// CompareOfferFinished returns true if mock invocations count is ok
func (m *RepositoryMock) CompareOfferFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.CompareOfferMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.CompareOfferCounter) == uint64(len(m.CompareOfferMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.CompareOfferMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.CompareOfferCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.CompareOfferFunc != nil {
		return atomic.LoadUint64(&m.CompareOfferCounter) > 0
	}

	return true
}

type mRepositoryMockGetMinerReputationById struct {
	mock              *RepositoryMock
	mainExpectation   *RepositoryMockGetMinerReputationByIdExpectation
	expectationSeries []*RepositoryMockGetMinerReputationByIdExpectation
}

// RepositoryMockGetMinerReputationByIdExpectation specifies expectation struct of the Repository.GetMinerReputationById
type RepositoryMockGetMinerReputationByIdExpectation struct {
	input  *RepositoryMockGetMinerReputationByIdInput
	result *RepositoryMockGetMinerReputationByIdResult
}

// RepositoryMockGetMinerReputationByIdInput represents input parameters of the Repository.GetMinerReputationById
type RepositoryMockGetMinerReputationByIdInput struct {
	p string
}

// RepositoryMockGetMinerReputationByIdResult represents results of the Repository.GetMinerReputationById
type RepositoryMockGetMinerReputationByIdResult struct {
	r float32
}

// Expect specifies that invocation of Repository.GetMinerReputationById is expected from 1 to Infinity times
func (m *mRepositoryMockGetMinerReputationById) Expect(p string) *mRepositoryMockGetMinerReputationById {
	m.mock.GetMinerReputationByIdFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RepositoryMockGetMinerReputationByIdExpectation{}
	}
	m.mainExpectation.input = &RepositoryMockGetMinerReputationByIdInput{p}
	return m
}

// Return specifies results of invocation of Repository.GetMinerReputationById
func (m *mRepositoryMockGetMinerReputationById) Return(r float32) *RepositoryMock {
	m.mock.GetMinerReputationByIdFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RepositoryMockGetMinerReputationByIdExpectation{}
	}
	m.mainExpectation.result = &RepositoryMockGetMinerReputationByIdResult{r}
	return m.mock
}

// ExpectOnce specifies that invocation of Repository.GetMinerReputationById is expected once
func (m *mRepositoryMockGetMinerReputationById) ExpectOnce(p string) *RepositoryMockGetMinerReputationByIdExpectation {
	m.mock.GetMinerReputationByIdFunc = nil
	m.mainExpectation = nil

	expectation := &RepositoryMockGetMinerReputationByIdExpectation{}
	expectation.input = &RepositoryMockGetMinerReputationByIdInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

// Return sets up return arguments of expectation struct for Repository.GetMinerReputationById
func (e *RepositoryMockGetMinerReputationByIdExpectation) Return(r float32) {
	e.result = &RepositoryMockGetMinerReputationByIdResult{r}
}

// Set uses given function f as a mock of Repository.GetMinerReputationById method
func (m *mRepositoryMockGetMinerReputationById) Set(f func(p string) (r float32)) *RepositoryMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetMinerReputationByIdFunc = f
	return m.mock
}

// GetMinerReputationById implements github.com/P2PCloud/miner-search-api/internal/app/test/mocks.Repository interface
func (m *RepositoryMock) GetMinerReputationById(p string) (r float32) {
	counter := atomic.AddUint64(&m.GetMinerReputationByIdPreCounter, 1)
	defer atomic.AddUint64(&m.GetMinerReputationByIdCounter, 1)

	if len(m.GetMinerReputationByIdMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetMinerReputationByIdMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to RepositoryMock.GetMinerReputationById. %v", p)
			return
		}

		input := m.GetMinerReputationByIdMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, RepositoryMockGetMinerReputationByIdInput{p}, "Repository.GetMinerReputationById got unexpected parameters")

		result := m.GetMinerReputationByIdMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the RepositoryMock.GetMinerReputationById")
			return
		}

		r = result.r

		return
	}

	if m.GetMinerReputationByIdMock.mainExpectation != nil {

		input := m.GetMinerReputationByIdMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, RepositoryMockGetMinerReputationByIdInput{p}, "Repository.GetMinerReputationById got unexpected parameters")
		}

		result := m.GetMinerReputationByIdMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the RepositoryMock.GetMinerReputationById")
		}

		r = result.r

		return
	}

	if m.GetMinerReputationByIdFunc == nil {
		m.t.Fatalf("Unexpected call to RepositoryMock.GetMinerReputationById. %v", p)
		return
	}

	return m.GetMinerReputationByIdFunc(p)
}

// GetMinerReputationByIdMinimockCounter returns a count of RepositoryMock.GetMinerReputationByIdFunc invocations
func (m *RepositoryMock) GetMinerReputationByIdMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetMinerReputationByIdCounter)
}

// GetMinerReputationByIdMinimockPreCounter returns the value of RepositoryMock.GetMinerReputationById invocations
func (m *RepositoryMock) GetMinerReputationByIdMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetMinerReputationByIdPreCounter)
}

// GetMinerReputationByIdFinished returns true if mock invocations count is ok
func (m *RepositoryMock) GetMinerReputationByIdFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetMinerReputationByIdMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetMinerReputationByIdCounter) == uint64(len(m.GetMinerReputationByIdMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetMinerReputationByIdMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetMinerReputationByIdCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetMinerReputationByIdFunc != nil {
		return atomic.LoadUint64(&m.GetMinerReputationByIdCounter) > 0
	}

	return true
}

type mRepositoryMockGetOffersByParams struct {
	mock              *RepositoryMock
	mainExpectation   *RepositoryMockGetOffersByParamsExpectation
	expectationSeries []*RepositoryMockGetOffersByParamsExpectation
}

// RepositoryMockGetOffersByParamsExpectation specifies expectation struct of the Repository.GetOffersByParams
type RepositoryMockGetOffersByParamsExpectation struct {
	input  *RepositoryMockGetOffersByParamsInput
	result *RepositoryMockGetOffersByParamsResult
}

// RepositoryMockGetOffersByParamsInput represents input parameters of the Repository.GetOffersByParams
type RepositoryMockGetOffersByParamsInput struct {
	p  context.Context
	p1 *models.VmType
}

// RepositoryMockGetOffersByParamsResult represents results of the Repository.GetOffersByParams
type RepositoryMockGetOffersByParamsResult struct {
	r  models.OfferCollection
	r1 error
}

// Expect specifies that invocation of Repository.GetOffersByParams is expected from 1 to Infinity times
func (m *mRepositoryMockGetOffersByParams) Expect(p context.Context, p1 *models.VmType) *mRepositoryMockGetOffersByParams {
	m.mock.GetOffersByParamsFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RepositoryMockGetOffersByParamsExpectation{}
	}
	m.mainExpectation.input = &RepositoryMockGetOffersByParamsInput{p, p1}
	return m
}

// Return specifies results of invocation of Repository.GetOffersByParams
func (m *mRepositoryMockGetOffersByParams) Return(r models.OfferCollection, r1 error) *RepositoryMock {
	m.mock.GetOffersByParamsFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RepositoryMockGetOffersByParamsExpectation{}
	}
	m.mainExpectation.result = &RepositoryMockGetOffersByParamsResult{r, r1}
	return m.mock
}

// ExpectOnce specifies that invocation of Repository.GetOffersByParams is expected once
func (m *mRepositoryMockGetOffersByParams) ExpectOnce(p context.Context, p1 *models.VmType) *RepositoryMockGetOffersByParamsExpectation {
	m.mock.GetOffersByParamsFunc = nil
	m.mainExpectation = nil

	expectation := &RepositoryMockGetOffersByParamsExpectation{}
	expectation.input = &RepositoryMockGetOffersByParamsInput{p, p1}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

// Return sets up return arguments of expectation struct for Repository.GetOffersByParams
func (e *RepositoryMockGetOffersByParamsExpectation) Return(r models.OfferCollection, r1 error) {
	e.result = &RepositoryMockGetOffersByParamsResult{r, r1}
}

// Set uses given function f as a mock of Repository.GetOffersByParams method
func (m *mRepositoryMockGetOffersByParams) Set(f func(p context.Context, p1 *models.VmType) (r models.OfferCollection, r1 error)) *RepositoryMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetOffersByParamsFunc = f
	return m.mock
}

// GetOffersByParams implements github.com/P2PCloud/miner-search-api/internal/app/test/mocks.Repository interface
func (m *RepositoryMock) GetOffersByParams(p context.Context, p1 *models.VmType) (r models.OfferCollection, r1 error) {
	counter := atomic.AddUint64(&m.GetOffersByParamsPreCounter, 1)
	defer atomic.AddUint64(&m.GetOffersByParamsCounter, 1)

	if len(m.GetOffersByParamsMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetOffersByParamsMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to RepositoryMock.GetOffersByParams. %v %v", p, p1)
			return
		}

		input := m.GetOffersByParamsMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, RepositoryMockGetOffersByParamsInput{p, p1}, "Repository.GetOffersByParams got unexpected parameters")

		result := m.GetOffersByParamsMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the RepositoryMock.GetOffersByParams")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.GetOffersByParamsMock.mainExpectation != nil {

		input := m.GetOffersByParamsMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, RepositoryMockGetOffersByParamsInput{p, p1}, "Repository.GetOffersByParams got unexpected parameters")
		}

		result := m.GetOffersByParamsMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the RepositoryMock.GetOffersByParams")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.GetOffersByParamsFunc == nil {
		m.t.Fatalf("Unexpected call to RepositoryMock.GetOffersByParams. %v %v", p, p1)
		return
	}

	return m.GetOffersByParamsFunc(p, p1)
}

// GetOffersByParamsMinimockCounter returns a count of RepositoryMock.GetOffersByParamsFunc invocations
func (m *RepositoryMock) GetOffersByParamsMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetOffersByParamsCounter)
}

// GetOffersByParamsMinimockPreCounter returns the value of RepositoryMock.GetOffersByParams invocations
func (m *RepositoryMock) GetOffersByParamsMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetOffersByParamsPreCounter)
}

// GetOffersByParamsFinished returns true if mock invocations count is ok
func (m *RepositoryMock) GetOffersByParamsFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetOffersByParamsMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetOffersByParamsCounter) == uint64(len(m.GetOffersByParamsMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetOffersByParamsMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetOffersByParamsCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetOffersByParamsFunc != nil {
		return atomic.LoadUint64(&m.GetOffersByParamsCounter) > 0
	}

	return true
}

type mRepositoryMockGetVmTypeById struct {
	mock              *RepositoryMock
	mainExpectation   *RepositoryMockGetVmTypeByIdExpectation
	expectationSeries []*RepositoryMockGetVmTypeByIdExpectation
}

// RepositoryMockGetVmTypeByIdExpectation specifies expectation struct of the Repository.GetVmTypeById
type RepositoryMockGetVmTypeByIdExpectation struct {
	input  *RepositoryMockGetVmTypeByIdInput
	result *RepositoryMockGetVmTypeByIdResult
}

// RepositoryMockGetVmTypeByIdInput represents input parameters of the Repository.GetVmTypeById
type RepositoryMockGetVmTypeByIdInput struct {
	p string
}

// RepositoryMockGetVmTypeByIdResult represents results of the Repository.GetVmTypeById
type RepositoryMockGetVmTypeByIdResult struct {
	r *models.VmType
}

// Expect specifies that invocation of Repository.GetVmTypeById is expected from 1 to Infinity times
func (m *mRepositoryMockGetVmTypeById) Expect(p string) *mRepositoryMockGetVmTypeById {
	m.mock.GetVmTypeByIdFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RepositoryMockGetVmTypeByIdExpectation{}
	}
	m.mainExpectation.input = &RepositoryMockGetVmTypeByIdInput{p}
	return m
}

// Return specifies results of invocation of Repository.GetVmTypeById
func (m *mRepositoryMockGetVmTypeById) Return(r *models.VmType) *RepositoryMock {
	m.mock.GetVmTypeByIdFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RepositoryMockGetVmTypeByIdExpectation{}
	}
	m.mainExpectation.result = &RepositoryMockGetVmTypeByIdResult{r}
	return m.mock
}

// ExpectOnce specifies that invocation of Repository.GetVmTypeById is expected once
func (m *mRepositoryMockGetVmTypeById) ExpectOnce(p string) *RepositoryMockGetVmTypeByIdExpectation {
	m.mock.GetVmTypeByIdFunc = nil
	m.mainExpectation = nil

	expectation := &RepositoryMockGetVmTypeByIdExpectation{}
	expectation.input = &RepositoryMockGetVmTypeByIdInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

// Return sets up return arguments of expectation struct for Repository.GetVmTypeById
func (e *RepositoryMockGetVmTypeByIdExpectation) Return(r *models.VmType) {
	e.result = &RepositoryMockGetVmTypeByIdResult{r}
}

// Set uses given function f as a mock of Repository.GetVmTypeById method
func (m *mRepositoryMockGetVmTypeById) Set(f func(p string) (r *models.VmType)) *RepositoryMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetVmTypeByIdFunc = f
	return m.mock
}

// GetVmTypeById implements github.com/P2PCloud/miner-search-api/internal/app/test/mocks.Repository interface
func (m *RepositoryMock) GetVmTypeById(p string) (r *models.VmType) {
	counter := atomic.AddUint64(&m.GetVmTypeByIdPreCounter, 1)
	defer atomic.AddUint64(&m.GetVmTypeByIdCounter, 1)

	if len(m.GetVmTypeByIdMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetVmTypeByIdMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to RepositoryMock.GetVmTypeById. %v", p)
			return
		}

		input := m.GetVmTypeByIdMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, RepositoryMockGetVmTypeByIdInput{p}, "Repository.GetVmTypeById got unexpected parameters")

		result := m.GetVmTypeByIdMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the RepositoryMock.GetVmTypeById")
			return
		}

		r = result.r

		return
	}

	if m.GetVmTypeByIdMock.mainExpectation != nil {

		input := m.GetVmTypeByIdMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, RepositoryMockGetVmTypeByIdInput{p}, "Repository.GetVmTypeById got unexpected parameters")
		}

		result := m.GetVmTypeByIdMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the RepositoryMock.GetVmTypeById")
		}

		r = result.r

		return
	}

	if m.GetVmTypeByIdFunc == nil {
		m.t.Fatalf("Unexpected call to RepositoryMock.GetVmTypeById. %v", p)
		return
	}

	return m.GetVmTypeByIdFunc(p)
}

// GetVmTypeByIdMinimockCounter returns a count of RepositoryMock.GetVmTypeByIdFunc invocations
func (m *RepositoryMock) GetVmTypeByIdMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetVmTypeByIdCounter)
}

// GetVmTypeByIdMinimockPreCounter returns the value of RepositoryMock.GetVmTypeById invocations
func (m *RepositoryMock) GetVmTypeByIdMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetVmTypeByIdPreCounter)
}

// GetVmTypeByIdFinished returns true if mock invocations count is ok
func (m *RepositoryMock) GetVmTypeByIdFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetVmTypeByIdMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetVmTypeByIdCounter) == uint64(len(m.GetVmTypeByIdMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetVmTypeByIdMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetVmTypeByIdCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetVmTypeByIdFunc != nil {
		return atomic.LoadUint64(&m.GetVmTypeByIdCounter) > 0
	}

	return true
}

// ValidateCallCounters checks that all mocked methods of the interface have been called at least once
// Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *RepositoryMock) ValidateCallCounters() {

	if !m.CompareOfferFinished() {
		m.t.Fatal("Expected call to RepositoryMock.CompareOffer")
	}

	if !m.GetMinerReputationByIdFinished() {
		m.t.Fatal("Expected call to RepositoryMock.GetMinerReputationById")
	}

	if !m.GetOffersByParamsFinished() {
		m.t.Fatal("Expected call to RepositoryMock.GetOffersByParams")
	}

	if !m.GetVmTypeByIdFinished() {
		m.t.Fatal("Expected call to RepositoryMock.GetVmTypeById")
	}

}

// CheckMocksCalled checks that all mocked methods of the interface have been called at least once
// Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *RepositoryMock) CheckMocksCalled() {
	m.Finish()
}

// Finish checks that all mocked methods of the interface have been called at least once
// Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *RepositoryMock) Finish() {
	m.MinimockFinish()
}

// MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *RepositoryMock) MinimockFinish() {

	if !m.CompareOfferFinished() {
		m.t.Fatal("Expected call to RepositoryMock.CompareOffer")
	}

	if !m.GetMinerReputationByIdFinished() {
		m.t.Fatal("Expected call to RepositoryMock.GetMinerReputationById")
	}

	if !m.GetOffersByParamsFinished() {
		m.t.Fatal("Expected call to RepositoryMock.GetOffersByParams")
	}

	if !m.GetVmTypeByIdFinished() {
		m.t.Fatal("Expected call to RepositoryMock.GetVmTypeById")
	}

}

// Wait waits for all mocked methods to be called at least once
// Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *RepositoryMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

// MinimockWait waits for all mocked methods to be called at least once
// this method is called by minimock.Controller
func (m *RepositoryMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && m.CompareOfferFinished()
		ok = ok && m.GetMinerReputationByIdFinished()
		ok = ok && m.GetOffersByParamsFinished()
		ok = ok && m.GetVmTypeByIdFinished()

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if !m.CompareOfferFinished() {
				m.t.Error("Expected call to RepositoryMock.CompareOffer")
			}

			if !m.GetMinerReputationByIdFinished() {
				m.t.Error("Expected call to RepositoryMock.GetMinerReputationById")
			}

			if !m.GetOffersByParamsFinished() {
				m.t.Error("Expected call to RepositoryMock.GetOffersByParams")
			}

			if !m.GetVmTypeByIdFinished() {
				m.t.Error("Expected call to RepositoryMock.GetVmTypeById")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

// AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
// it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *RepositoryMock) AllMocksCalled() bool {

	if !m.CompareOfferFinished() {
		return false
	}

	if !m.GetMinerReputationByIdFinished() {
		return false
	}

	if !m.GetOffersByParamsFinished() {
		return false
	}

	if !m.GetVmTypeByIdFinished() {
		return false
	}

	return true
}
