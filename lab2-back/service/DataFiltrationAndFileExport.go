package service

import (
	"encoding/csv"
	"errors"
	"github.com/Edrudo/otvoreno_lab/storage"
	"os"
	"strconv"
	"strings"
)

type DataFilter struct {
	FieldName string
	Value     string
}

func FilterData(data storage.MyOpenData, filter DataFilter) (storage.MyOpenData, error) {
	if filter.FieldName == "vehicleBrand" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strings.Contains(strings.ToUpper(car.VehicleBrand), strings.ToUpper(filter.Value)) {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "model" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strings.Contains(strings.ToUpper(car.Model), strings.ToUpper(filter.Value)) {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "year" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strconv.Itoa(car.Year) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "engineName" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strings.Contains(strings.ToUpper(car.Engine.Name), strings.ToUpper(filter.Value)) {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "engineType" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strings.Contains(strings.ToUpper(car.Engine.Type), strings.ToUpper(filter.Value)) {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "powerOutput" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strconv.Itoa(car.PowerOutput) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "torque" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strconv.Itoa(car.Torque) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "topSpeed" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strconv.Itoa(car.TopSpeed) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "acceleration" {

		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strings.Contains(strconv.FormatFloat(car.Acceleration, 'f', 2, 64), filter.Value) {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "weight" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strconv.Itoa(car.Weight) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "bootSpace" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strconv.Itoa(car.BootSpace) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "wikipediaSufix" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			if strings.Contains(strings.ToUpper(car.WikipediaSufix), strings.ToUpper(filter.Value)) {
				retValue.OpenData = append(retValue.OpenData, car)
			}
		}
		return retValue, nil

	} else if filter.FieldName == "awardEvent" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			for _, award := range car.Awards {
				if strings.Contains(strings.ToUpper(award.Event), strings.ToUpper(filter.Value)) {
					retValue.OpenData = append(retValue.OpenData, car)
					break
				}
			}
		}
		return retValue, nil

	} else if filter.FieldName == "award" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}
		for _, car := range data.OpenData {
			for _, award := range car.Awards {
				if strings.Contains(strings.ToUpper(award.Award), strings.ToUpper(filter.Value)) {
					retValue.OpenData = append(retValue.OpenData, car)
					break
				}
			}
		}
		return retValue, nil
	} else if filter.FieldName == "everyFieldName" {
		retValue := storage.MyOpenData{OpenData: make([]storage.MyOpenDataStruct, 0)}

		for _, car := range data.OpenData {
			added := false

			if strings.Contains(strings.ToUpper(car.VehicleBrand), strings.ToUpper(filter.Value)) {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strings.Contains(strings.ToUpper(car.Model), strings.ToUpper(filter.Value)) {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strconv.Itoa(car.Year) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strings.Contains(strings.ToUpper(car.Engine.Name), strings.ToUpper(filter.Value)) {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strings.Contains(strings.ToUpper(car.Engine.Type), strings.ToUpper(filter.Value)) {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strconv.Itoa(car.PowerOutput) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strconv.Itoa(car.TopSpeed) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strings.Contains(strconv.FormatFloat(car.Acceleration, 'f', 2, 64), filter.Value) {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strconv.Itoa(car.Weight) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strconv.Itoa(car.Torque) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strconv.Itoa(car.BootSpace) == filter.Value {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			if strings.Contains(strings.ToUpper(car.WikipediaSufix), strings.ToUpper(filter.Value)) {
				retValue.OpenData = append(retValue.OpenData, car)
				continue
			}

			for _, award := range car.Awards {
				if strings.Contains(strings.ToUpper(award.Award), strings.ToUpper(filter.Value)) {
					retValue.OpenData = append(retValue.OpenData, car)
					added = true
					break
				}
			}

			if added {
				continue
			}

			for _, award := range car.Awards {
				if strings.Contains(strings.ToUpper(award.Event), strings.ToUpper(filter.Value)) {
					retValue.OpenData = append(retValue.OpenData, car)
					added = true
					break
				}
			}
			if added {
				continue
			}
		}
		return retValue, nil
	}

	return storage.MyOpenData{}, errors.New("Filter is not valid")
}

func ExportDataAsCSV(data storage.MyOpenData) (*os.File, error) {
	csvFile, err := os.Create("cars.csv")
	if err != nil {
		return &os.File{}, err
	}

	csvWriter := csv.NewWriter(csvFile)

	empData := []string{"Vehicle brand", "model", "year", "Engine name", "Engine type", "Power output", "torque", "topSpeed", "acceleration",
		"weight", "bootSpace", "wikipediaSufix", "awards"}
	csvWriter.Write(empData)

	for _, carData := range data.OpenData {

		awards := ""
		for i, award := range carData.Awards {
			awards = awards + award.Award + " from " + award.Event
			if i != len(carData.Awards)-1 {
				awards = awards + ", "
			}
		}

		carRecord := []string{carData.VehicleBrand, carData.Model, strconv.Itoa(carData.Year), carData.Engine.Name, carData.Engine.Type, strconv.Itoa(carData.PowerOutput),
			strconv.Itoa(carData.Torque), strconv.Itoa(carData.TopSpeed), strconv.FormatFloat(carData.Acceleration, 'f', 2, 64),
			strconv.Itoa(carData.Weight), strconv.Itoa(carData.BootSpace), carData.WikipediaSufix, awards}

		err := csvWriter.Write(carRecord)
		if err != nil {
			return &os.File{}, nil
		}
	}
	csvWriter.Flush()
	return csvFile, nil
}
