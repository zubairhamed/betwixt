package tests

type MockServerStatistics struct {
}

func (s *MockServerStatistics) IncrementCoapRequestsCount() {

}

func (s *MockServerStatistics) GetRequestsCount() int {
	return 0
}
