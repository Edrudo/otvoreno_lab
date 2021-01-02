package storage

type MyOpenDataStructDIO struct {
	VehicleBrand   string       `bson:"vehicleBrand", json:"vehicleBrand"`
	Model          string       `json:"model"`
	Year           int          `json:"year"`
	Engine         EngineStruct `json:"engine"`
	PowerOutput    int          `bson:"powerOutput"`
	Torque         int
	TopSpeed       int `bson:"topSpeed"`
	Acceleration   float64
	Weight         int
	BootSpace      int    `bson:"bootSpace", json:"bootSpace"`
	WikipediaSufix string `bson:"wikipediaSufix", json:"wikipediaSufix"`
	Awards         []Award
}

type EngineStructDIO struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AwardDIO struct {
	Event string `json:"event"`
	Award string `json:"award"`
}
