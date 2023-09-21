package dto


type CreateCityRequest struct {
	Name string `json:"name" binding:"size=30,required"`
}

type CityResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
}


type UpdateCityRequest struct {
	Name string `json:"name"`
}
