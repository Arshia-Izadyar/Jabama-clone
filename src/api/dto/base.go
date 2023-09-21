package dto

type CreateCityRequest struct {
	Name string `json:"name" binding:"max=30,required"`
}

type CityResponse struct {
	Id        int                `json:"id"`
	Name      string             `json:"name"`
	Provinces []ProvinceResponse `json:"provinces,omitempty"`
}

type UpdateCityRequest struct {
	Name string `json:"name"`
}

type CreateProvinceRequest struct {
	Name   string `json:"name" binding:"max=30,required"`
	CityId int    `json:"cityId" binding:"required"`
}

type UpdateProvinceRequest struct {
	Name string `json:"name" binding:"max=30,required"`
}

type ProvinceResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"max=30,required"`
}

type CreatePropertyCategoryRequest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type UpdatePropertyCategoryRequest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type PropertyCategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type CreatePropertyRequest struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	CategoryId  int    `json:"categoryId"`
}

type UpdatePropertyRequest struct {
	Description string `json:"description"`
	Icon        string `json:"icon"`
	CategoryId  int    `json:"categoryId"`
}

type PropertyResponse struct {
	Id          int                      `json:"id"`
	Description string                   `json:"description"`
	Name        string                   `json:"name"`
	Icon        string                   `json:"icon"`
	Category    PropertyCategoryResponse `json:"category"`
}
