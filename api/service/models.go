package service

import "github.com/Edrudo/otvoreno_lab/storage"

type ApiResponse struct {
	Response []ApiResponseObject `json:"response"`
	Status   string              `json:"status"`
	Message  string              `json:"message"`
}

type ApiResponsePostAndPut struct {
	Response ApiResponseObject `json:"response"`
	Status   string            `json:"status"`
	Message  string            `json:"message"`
}

type ApiResponseDelete struct {
	Response ApiResponseObjectDelete `json:"response"`
	Status   string                  `json:"status"`
	Message  string                  `json:"message"`
}

type ApiResponseObjectDelete struct {
	Links []Hateoas `json:"links"`
}

type Hateoas struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type"`
}

type ApiResponseObject struct {
	VehicleBrand   string               `json:"vehicleBrand, "bson:"vehicleBrand"`
	Model          string               `json:"model"`
	Year           int                  `json:"year"`
	Engine         storage.EngineStruct `json:"engine"`
	PowerOutput    int                  `json:"powerOutput, "bson:"powerOutput"`
	Torque         int                  `json:"torque"`
	TopSpeed       int                  `json:"topSpeed,"bson:"topSpeed"`
	Acceleration   float64              `json:"acceleration"`
	Weight         int                  `json:"weight"`
	BootSpace      int                  `json:"bootSpace"`
	WikipediaSufix string               `json:"wikipedia_sufix, "bson:"wikipediaSufix"`
	Awards         []storage.Award      `json:"awards"`
	Links          []Hateoas            `json:"links"`
}
