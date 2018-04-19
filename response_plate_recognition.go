package alpr

// ResponsePlateRecognition :
type ResponsePlateRecognition struct {
	UUID           string `json:"uuid"`
	DataType       string `json:"data_type"`
	EpochTime      int64  `json:"epoch_time"`
	ProcessingTime struct {
		Total    float64 `json:"total"`
		Plates   float64 `json:"plates"`
		Vehicles float64 `json:"vehicles"`
	} `json:"processing_time"`
	ImgHeight int `json:"img_height"`
	ImgWidth  int `json:"img_width"`
	Results   []struct {
		Plate            string  `json:"plate"`
		Confidence       float64 `json:"confidence"`
		RegionConfidence int     `json:"region_confidence"`
		VehicleRegion    struct {
			Y      int `json:"y"`
			X      int `json:"x"`
			Height int `json:"height"`
			Width  int `json:"width"`
		} `json:"vehicle_region"`
		Region           string  `json:"region"`
		PlateIndex       int     `json:"plate_index"`
		ProcessingTimeMs float64 `json:"processing_time_ms"`
		Candidates       []struct {
			MatchesTemplate int     `json:"matches_template"`
			Plate           string  `json:"plate"`
			Confidence      float64 `json:"confidence"`
		} `json:"candidates"`
		Coordinates []struct {
			Y int `json:"y"`
			X int `json:"x"`
		} `json:"coordinates"`
		Vehicle struct {
			Orientation []struct {
				Confidence float64 `json:"confidence"`
				Name       string  `json:"name"`
			} `json:"orientation"`
			Color []struct {
				Confidence float64 `json:"confidence"`
				Name       string  `json:"name"`
			} `json:"color"`
			Make []struct {
				Confidence float64 `json:"confidence"`
				Name       string  `json:"name"`
			} `json:"make"`
			BodyType []struct {
				Confidence float64 `json:"confidence"`
				Name       string  `json:"name"`
			} `json:"body_type"`
			Year []struct {
				Confidence float64 `json:"confidence"`
				Name       string  `json:"name"`
			} `json:"year"`
			MakeModel []struct {
				Confidence float64 `json:"confidence"`
				Name       string  `json:"name"`
			} `json:"make_model"`
		} `json:"vehicle"`
		MatchesTemplate int `json:"matches_template"`
		RequestedTopn   int `json:"requested_topn"`
	} `json:"results"`
	CreditsMonthlyUsed  int  `json:"credits_monthly_used"`
	Version             int  `json:"version"`
	CreditsMonthlyTotal int  `json:"credits_monthly_total"`
	Error               bool `json:"error"`
	RegionsOfInterest   []struct {
		Y      int `json:"y"`
		X      int `json:"x"`
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"regions_of_interest"`
	CreditCost int `json:"credit_cost"`
}
