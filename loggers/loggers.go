package loggers

import "log/slog"

type Loggers struct {
	GlobalLogger *slog.Logger
}

var DefaultLogger *Loggers = &Loggers{
	GlobalLogger: slog.Default(),
}
