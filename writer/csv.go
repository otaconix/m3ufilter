package writer

import (
	"encoding/csv"
	"github.com/otaconix/m3ufilter/m3u"
	"io"
)

var csvHeaders = []string{
	"tvg-id",
	"group-title",
	"group-title",
	"tvg-name",
	"duration",
	"tvg-logo",
	"uri",
}

func writeCsv(w io.Writer, streams []*m3u.Stream) {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	err := writer.Write(csvHeaders)
	if err != nil {
		log.Errorf("Could not write csv header, err = %s", err)
	}

	for _, stream := range streams {
		printPlaylist(stream, writer)
	}
}

func printPlaylist(pl *m3u.Stream, w *csv.Writer) {
	row := []string{
		pl.Id,
		pl.Group,
		pl.Name,
		pl.Duration,
		pl.Logo,
		pl.Uri,
	}

	err := w.Write(row)
	if err != nil {
		log.Errorf("Could not write csv row, row = %v, err = %s", row, err)
	}
}
