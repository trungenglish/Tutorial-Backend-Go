package seed

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"tutorial/model"

	"gorm.io/gorm"
)

func SeedMovies(db *gorm.DB) error {
	file, err := os.Open("service/db/seed/title.basics.tsv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	i := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if i == 0 {
			i++
			continue
		}

		if len(record) < 9 {
			continue
		}

		year := 0
		if record[5] != `\N` {
			year, _ = strconv.Atoi(record[5])
		}

		movie := model.Movies{
			Title: record[2],
			Genre: record[8],
			Year:  year,
		}
		db.Create(&movie)

		i++
		if i > 100 {
			break
		}
	}

	return nil
}
