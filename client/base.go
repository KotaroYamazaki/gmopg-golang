package client

// BaseRequest is base struct for api request.
type BaseRequest struct {
	Version  string `url:"Version,omitempty" json:"-"`
	ShopID   string `url:"ShopID,omitempty" json:"-"`
	ShopPass string `url:"ShopPass,omitempty" json:"-"`
}

// BaseResponse is base struct for api response.
type BaseResponse struct {
	ErrCode string `url:"ErrCode"`
	ErrInfo string `url:"ErrInfo"`
}
