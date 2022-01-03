package domains

type ConnectivityResult struct {
	Result bool
}

func CreateResultObject(res bool) ConnectivityResult {
	return ConnectivityResult{
		Result: res,
	}
}