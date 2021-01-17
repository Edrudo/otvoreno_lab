package service

import (
	"github.com/Edrudo/otvoreno_lab/storage"
)

func GetAllCarsWithLimit(repo storage.Repository, limit int64) (ApiResponse, error) {
	data, err := storage.GetMyOpenDataWithLimit(repo, limit)
	if err != nil {
		return ApiResponse{}, err
	}
	return ModelDataForResponse(data), nil
}

func GetAllCarsFromSomeYear(repo storage.Repository, year int64) (ApiResponse, error) {
	data, err := storage.GetMyOpenDataFromSomeYear(repo, year)
	if err != nil {
		return ApiResponse{}, err
	}
	return ModelDataForResponse(data), nil
}

func GetAllCarsWithSpeedLimit(repo storage.Repository, lowerLimit, upperLimit int) (ApiResponse, error) {
	allData, err := storage.GetMyOpenData(repo)
	if err != nil {
		return ApiResponse{}, err
	}

	data := storage.MyOpenData{[]storage.MyOpenDataStruct{}}

	// filtering all data
	for _, car := range allData.OpenData {
		if car.TopSpeed >= lowerLimit && car.TopSpeed <= upperLimit {
			data.OpenData = append(data.OpenData, car)
		}
	}

	return ModelDataForResponse(data), nil
}
func GetAllCarsWithPowerOutputLimit(repo storage.Repository, lowerLimit, upperLimit int) (ApiResponse, error) {
	allData, err := storage.GetMyOpenData(repo)
	if err != nil {
		return ApiResponse{}, err
	}

	data := storage.MyOpenData{[]storage.MyOpenDataStruct{}}

	// filtering all data
	for _, car := range allData.OpenData {
		if car.PowerOutput >= lowerLimit && car.PowerOutput <= upperLimit {
			data.OpenData = append(data.OpenData, car)
		}
	}

	return ModelDataForResponse(data), nil
}
func GetAllCarsWithBootSpaceLimit(repo storage.Repository, lowerLimit, upperLimit int) (ApiResponse, error) {
	allData, err := storage.GetMyOpenData(repo)
	if err != nil {
		return ApiResponse{}, err
	}

	data := storage.MyOpenData{[]storage.MyOpenDataStruct{}}

	// filtering all data
	for _, car := range allData.OpenData {
		if car.BootSpace >= lowerLimit && car.BootSpace <= upperLimit {
			data.OpenData = append(data.OpenData, car)
		}
	}

	return ModelDataForResponse(data), nil
}
