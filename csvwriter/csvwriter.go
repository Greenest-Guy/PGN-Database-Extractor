package csvwriter

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type ChessGame struct {
	WhiteElo    int32
	BlackElo    int32
	TimeControl string
	Moves       string
}

type Writer struct {
	file   *os.File
	writer *csv.Writer
	count  int
}

func New(filename string) (*Writer, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("can't create CSV file: %w", err)
	}

	writer := csv.NewWriter(file)

	header := []string{"WhiteElo", "BlackElo", "TimeControl", "Moves"}
	if err := writer.Write(header); err != nil {
		file.Close()
		return nil, fmt.Errorf("can't write header: %w", err)
	}

	return &Writer{
		file:   file,
		writer: writer,
		count:  0,
	}, nil
}

func (w *Writer) Append(game ChessGame) error {
	row := []string{
		strconv.Itoa(int(game.WhiteElo)),
		strconv.Itoa(int(game.BlackElo)),
		game.TimeControl,
		game.Moves,
	}

	if err := w.writer.Write(row); err != nil {
		return fmt.Errorf("failed to write game: %w", err)
	}

	w.count++
	return nil
}

func (w *Writer) Close() error {
	w.writer.Flush()
	if err := w.writer.Error(); err != nil {
		w.file.Close()
		return fmt.Errorf("error flushing CSV: %w", err)
	}
	if err := w.file.Close(); err != nil {
		return fmt.Errorf("error closing file: %w", err)
	}
	return nil
}

func (w *Writer) GetRowCount() int {
	return w.count
}
