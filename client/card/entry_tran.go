package card

import "gmopg-card/client"

// EntryTran is struct for EntryTran API.
// Entry for payment by card.
type EntryTran struct {
	client.BaseRequest `url:",squash"`

	OrderID string `url:"OrderID"`
	JobCd   string `url:"JobCd"`
	Amount  int64  `url:"Amount"`

	Tax         int64  `url:"Tax,omitempty"`
	PaymentType string `url:"PaymentType,omitempty"`
}

// EntryTranResponse is struct for response of EntryTran API.
type EntryTranResponse struct {
	client.BaseResponse `url:",squash"`

	AccessID   string `url:"AccessID"`
	AccessPass string `url:"AccessPass"`
}

func (svc *EntryTran) Do(cli client.Client) (*EntryTranResponse, error) {
	const apiPath = "/payment/EntryTran.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	if svc.OrderID == "" {
		svc.OrderID = client.GetRandomOrderID()
	}

	result := &EntryTranResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

// IsSuccess checks the request is success or not
func (r *EntryTranResponse) IsSuccess() bool {
	switch {
	case !r.BaseResponse.IsSuccess(),
		r.AccessID == "",
		r.AccessPass == "":
		return false
	}
	return true
}
