package slacker


type logger interface {
	Output(int, string) error
}

type ilogger interface {
	logger
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

