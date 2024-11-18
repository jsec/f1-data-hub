package imager

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
)

const zipName = "seed.zip"

func FetchSeedData() error {
	out, err := os.Create(zipName)
	if err != nil {
		return fmt.Errorf("Error creating seed file: %w", err)
	}
	defer out.Close()

	resp, err := http.Get("http://ergast.com/downloads/f1db_csv.zip")
	if err != nil {
		return fmt.Errorf("Error fetching seed data: %w", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("Error saving seed data: %w", err)
	}

	return UnzipSeedData()
}

func UnzipSeedData() error {
	r, err := zip.OpenReader(zipName)
	if err != nil {
		return fmt.Errorf("Error opening zipfile: %w", err)
	}
	defer r.Close()

	err = os.MkdirAll("data", 0755)
	if err != nil {
		return fmt.Errorf("Error creating data dir: %w", err)
	}

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return fmt.Errorf("Error extracting file %s: %w", f.Name, err)
		}
		defer rc.Close()

		filePath := fmt.Sprintf("data/%s", f.Name)
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("Error creating uncompressed file %s: %w", f.Name, err)
		}

		_, err = io.Copy(file, rc)
		if err != nil {
			return fmt.Errorf("Error decompressing file %s: %w", f.Name, err)
		}
	}

	return nil
}
