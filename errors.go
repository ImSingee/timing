package timing

type InvalidRequest struct {
	info string
}

func (e *InvalidRequest) Error() string {
	return "Invalid response " + e.info
}

type InvalidResponse struct {
	info string
}

func (e *InvalidResponse) Error() string {
	return "Invalid response " + e.info
}
