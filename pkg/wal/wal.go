package wal

type Wal struct {
	// The log file
	LogFile string

	// The name of the file has a timestamp
	FileNameHasTimestamp bool

	// The current offset
	Offset int64

	// The maximum size of the log file
	MaxSize int64

	// The max time
	MaxTime int64
}

func New(logFile string, maxSize int64, maxTime int64, fileNameHasTimestamp bool) *Wal {
	return &Wal{
		LogFile:              logFile,
		Offset:               0,
		MaxSize:              maxSize,
		MaxTime:              maxTime,
		FileNameHasTimestamp: fileNameHasTimestamp,
	}
}

func (w *Wal) Write(data []byte) error {
	// Write the data to the log file
	return nil
}

func (w *Wal) Flush() error {
	// Read the data from the log file
	return nil
}

func (w *Wal) Recover() error {
	// Recover the log file
	return nil
}
