package app

type BetwixtServerStatistics struct {
	requestCount int
}

func (s *BetwixtServerStatistics) IncrementCoapRequestsCount() {
	s.requestCount++
}

func (s *BetwixtServerStatistics) GetRequestsCount() int {
	return s.requestCount
}
