package healthcheck

type Status int

const (
	StatusHealthy = Status(iota)
	StatusDegraded
	StatusUnhealthy
)

func (s Status) Int() int {
	return int(s)
}

func (s Status) String() string {
	switch s {
	case 0:
		return "healthy"
	case 1:
		return "degraded"
	case 2:
		return "unhealthy"
	default:
		panic("invalid status")
	}
}
