package v3

import "strconv"

type RegeoResponse struct {
	Status    string     `json:"status"`
	Info      string     `json:"info"`
	Infocode  string     `json:"infocode"`
	Regeocode *regeocode `json:"regeocode"`
}

func (response RegeoResponse) GetAddressName() string {
	if response.Regeocode == nil {
		return ""
	} else {
		regeocode := response.Regeocode
		if len(regeocode.Pois) == 0 && len(regeocode.Aois) == 0 {
			return regeocode.FormattedAddress
		} else {
			// poi
			var min_poi *poi = nil
			for _, poi := range regeocode.Pois {
				if min_poi == nil || poi.GetDistance() < min_poi.GetDistance() {
					min_poi = poi
				}
			}
			var min_aoi *aoi = nil
			for _, aoi := range regeocode.Aois {
				if min_aoi == nil || aoi.GetDistance() < min_aoi.GetDistance() {
					min_aoi = aoi
				}
			}
			if min_poi != nil && min_aoi == nil {
				return min_poi.Name
			} else if min_poi == nil && min_aoi != nil {
				return min_aoi.Name
			} else {
				if min_poi.GetDistance() <= min_aoi.GetDistance() {
					return min_poi.Name
				} else {
					return min_aoi.Name
				}
			}
		}
	}
}

type regeocode struct {
	FormattedAddress string `json:"formatted_address"`
	Pois             []*poi `json:"pois"`
	Aois             []*aoi `json:"aois"`
}
type poi struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	// Tel          string   `json:"tel"`
	// Direction    string   `json:"direction"`
	Distance string `json:"distance"`
	Location string `json:"location"`
	Address  string `json:"address"`
	// Poiweight    string   `json:"poiweight"`
	// Businessarea []string `json:"businessarea"`
}

func (poi poi) GetDistance() float64 {
	distance, _ := strconv.ParseFloat(poi.Distance, 64)
	return distance
}

type aoi struct {
	Area     string `json:"area"`
	Type     string `json:"type"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Adcode   string `json:"adcode"`
	Distance string `json:"distance"`
}

func (aoi aoi) GetDistance() float64 {
	distance, _ := strconv.ParseFloat(aoi.Distance, 64)
	return distance
}
