package v3

type amaphelper interface {
	/**
	逆地理编码
	longitude: 经度
	latitude: 纬度
	*/
	Regeo(float64, float64) (*RegeoResponse, error)
	/**
	地理编码
	*/
	Geo() (float64, float64)
}
