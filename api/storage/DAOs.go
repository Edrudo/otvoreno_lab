package storage

import "go.mongodb.org/mongo-driver/bson/primitive"

type MyOpenData struct {
	OpenData []MyOpenDataStruct
}

type MyOpenDataStruct struct {
	VehicleBrand   string       `bson:"vehicleBrand", json:"vehicleBrand"`
	Model          string       `json:"model"`
	Year           int          `json:"year"`
	Engine         EngineStruct `json:"engine"`
	PowerOutput    int          `bson:"powerOutput"`
	Torque         int
	TopSpeed       int `bson:"topSpeed"`
	Acceleration   float64
	Weight         int
	BootSpace      int    `json:"bootSpace", bson:"bootSpace"`
	WikipediaSufix string `bson:"wikipediaSufix", json:"wikipediaSufix"`
	Awards         []Award
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}

type EngineStruct struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Award struct {
	Event string `json:"event"`
	Award string `json:"award"`
}
