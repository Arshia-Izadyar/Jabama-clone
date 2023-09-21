package dto


type CreateCityRequest struct {
	Name string `json:"name" binding:"max=30,required"`
}

type CityResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Provinces []ProvinceResponse `json:"provinces,omitempty"`
}


type UpdateCityRequest struct {
	Name string `json:"name"`
}


type CreateProvinceRequest struct {
	Name string `json:"name" binding:"max=30,required"`
	CityId int `json:"cityId" binding:"required"`
}


type UpdateProvinceRequest struct {
	Name string `json:"name" binding:"max=30,required"`
}

type ProvinceResponse struct {
	Id int `json:"id"`
	Name string `json:"name" binding:"max=30,required"`
}