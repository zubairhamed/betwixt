package fragments

type StatsFragment struct {
}

func (p *StatsFragment) GetContent() []byte {
	return []byte(p.content())
}

func (p *StatsFragment) content() string {
	return ``
}
