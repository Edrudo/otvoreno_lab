package storage

type MyOpenData struct {
	OpenData []MyOpenDataStruct
}

type MyOpenDataStruct struct {
	VehicleBrand   string `bson:"vehicleBrand"`
	Model          string
	Year           int
	Engine         EngineStruct
	PowerOutput    int `bson:"powerOutput"`
	Torque         int
	TopSpeed       int `bson:"topSpeed"`
	Acceleration   float64
	Weight         int
	BootSpace      int
	WikipediaSufix string `bson:"wikipediaSufix"`
	Awards         []Award
}

type EngineStruct struct {
	Name string
	Type string
}

type Award struct {
	Event string
	Award string
}
