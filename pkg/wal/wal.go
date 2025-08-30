package wal

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type Wal struct {
	Mu         sync.Mutex
	WriteTimes []time.Time
	pathToWal  string
	StopFlush  chan bool
	Ticker     *time.Ticker
}

func RestoreOrCreateAnew(pathToWal string) (*Wal, bool) {
	wal := &Wal{
		pathToWal: pathToWal,
		StopFlush: make(chan bool),
		Ticker:    time.NewTicker(time.Second),
	}

	// We check if we already have a non-empty file on the directory of the WAL
	if _, err := os.Stat(pathToWal); err == nil {
		return wal, true
	}

	return wal, false
}

func (w *Wal) WriteToWal(content string) error {
	w.Mu.Lock()
	defer w.Mu.Unlock()

	// Check if the WAL file exists
	file, err := fileExistsOrCreate(w.pathToWal)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s\n", content)
	if err != nil {
		return err
	}

	w.WriteTimes = append(w.WriteTimes, time.Now())

	return nil
}

func (w *Wal) ClearWal() error {
	file, err := os.OpenFile(w.pathToWal, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "")
	if err != nil {
		return err
	}

	return nil
}

func (w *Wal) ReplayWal() ([]string, error) {
	file, err := os.Open(w.pathToWal)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, scanner.Err()
}

func (w *Wal) FlushOldEntries(maxSize, maxTime int64) {
	for {
		select {
		case <-w.StopFlush:
			return
		case t := <-w.Ticker.C:
			// If there are no times put in, we can just move forward and ignore the tick
			if len(w.WriteTimes) == 0 {
				continue
			}

			// TODO: Consider doing a WAL shifting instead of a whole clearing?
			// If the last entry is older than the threshold, we have to clear the WAL
			// If t + maxTime < now -> flush
			if w.WriteTimes[len(w.WriteTimes)-1].Add(time.Second * time.Duration(maxTime)).Before(t) {
				if err := w.ClearWal(); err != nil {
					fmt.Println("Error clearing WAL:", err)
				}
			}

			// If the WAL is bigger than the size threshold, we will clear it
			file, err := os.Stat(w.pathToWal)
			if err == nil && file.Size() > maxSize {
				if err := w.ClearWal(); err != nil {
					fmt.Println("Error clearing WAL:", err)
				}
			}
		}
	}
}

func fileExistsOrCreate(pathToWal string) (*os.File, error) {
	if _, err := os.Stat(pathToWal); os.IsNotExist(err) {
		// Create the WAL file if it doesn't exist
		file, err := os.Create(pathToWal)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return nil, nil
}
