package imager

import (
	"strconv"
	"strings"
	"time"
)

type dateOnly struct {
	Value time.Time
}

func (d *dateOnly) UnmarshalCSV(csv string) (err error) {
	layout := "2006-01-02"
	d.Value, err = time.Parse(layout, csv)
	return err
}

type timeOnly struct {
	Value time.Time
}

func (t *timeOnly) UnmarshalCSV(csv string) (err error) {
	t.Value, err = time.Parse(time.TimeOnly, csv)
	return err
}

type optionalNumber struct {
	Value *int32
}

func (o *optionalNumber) UnmarshalCSV(csv string) (err error) {
	raw, err := strconv.Atoi(csv)
	if err != nil {
		o.Value = nil
		return nil
	}

	value := int32(raw)
	o.Value = &value
	return nil
}

type optionalString struct {
	Value *string
}

func (o *optionalString) UnmarshalCSV(csv string) (err error) {
	if strings.Contains(csv, "\\N") {
		o.Value = nil
	} else {
		o.Value = &csv
	}

	return nil
}

type optionalTimeOnly struct {
	Value *time.Time
}

func (o *optionalTimeOnly) UnmarshalCSV(csv string) (err error) {
	value, err := time.Parse(time.TimeOnly, csv)
	if err != nil {
		o.Value = nil
	} else {
		o.Value = &value
	}

	return nil
}

type optionalDateOnly struct {
	Value *time.Time
}

func (o *optionalDateOnly) UnmarshalCSV(csv string) (err error) {
	layout := "2006-01-02"
	value, err := time.Parse(layout, csv)
	if err != nil {
		o.Value = nil
	} else {
		o.Value = &value
	}

	return nil
}
