package service

import (
	"github.com/Edrudo/otvoreno_lab/storage"
)

func ModelDataForResponse(data storage.MyOpenData, message string) ApiResponse {
	retValue := ApiResponse{
		Status:   "OK",
		Message:  message,
		Response: []ApiResponseObject{},
	}

	for _, car := range data.OpenData {
		newRetValueObject := ApiResponseObject{
			VehicleBrand:   car.VehicleBrand,
			Model:          car.Model,
			Year:           car.Year,
			Engine:         car.Engine,
			PowerOutput:    car.PowerOutput,
			Torque:         car.Torque,
			TopSpeed:       car.TopSpeed,
			Acceleration:   car.Acceleration,
			Weight:         car.Weight,
			BootSpace:      car.BootSpace,
			WikipediaSufix: car.WikipediaSufix,
			Awards:         car.Awards,
			Links: []Hateoas{
				{
					Href: "/cars?id=" + car.ID.Hex(),
					Rel:  "update",
					Type: "PUT",
				},
				{
					Href: "/cars?id=" + car.ID.Hex(),
					Rel:  "delete",
					Type: "DELETE",
				},
				{
					Href: "/cars",
					Rel:  "create",
					Type: "POST",
				}},
		}

		retValue.Response = append(retValue.Response, newRetValueObject)
	}

	return retValue
}

func ModelOneCarDataForResponse(car storage.MyOpenDataStruct) ApiResponseObject {

	return ApiResponseObject{
		VehicleBrand:   car.VehicleBrand,
		Model:          car.Model,
		Year:           car.Year,
		Engine:         car.Engine,
		PowerOutput:    car.PowerOutput,
		Torque:         car.Torque,
		TopSpeed:       car.TopSpeed,
		Acceleration:   car.Acceleration,
		Weight:         car.Weight,
		BootSpace:      car.BootSpace,
		WikipediaSufix: car.WikipediaSufix,
		Awards:         car.Awards,
		Links: []Hateoas{
			{
				Href: "/cars?id=" + car.ID.Hex(),
				Rel:  "update",
				Type: "PUT",
			},
			{
				Href: "/cars?id=" + car.ID.Hex(),
				Rel:  "delete",
				Type: "DELETE",
			},
			{
				Href: "/cars",
				Rel:  "create",
				Type: "POST",
			}},
	}
}
