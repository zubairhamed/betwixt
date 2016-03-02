package webadmin

type DefaultServerStatistics struct {
	requestsCount int
}

func (s *DefaultServerStatistics) IncrementCoapRequestsCount() {
	s.requestsCount++
}

func (s *DefaultServerStatistics) GetRequestsCount() int {
	return s.requestsCount
}
