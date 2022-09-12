package git

type Action struct {
	Label   string
	Process func(string, chan *Result)
}

type Result struct {
	Output string
	Err    error
}
