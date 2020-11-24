package card

import "gmopg-card/client"

// SaveCard is struct for SaveCard API.
// Save card information for payment
type SaveCard struct {
	client.BaseRequest `url:",squash"`

	// REQUIRED
	MemberID string `url:"MemberID"`
	CardNo   string `url:"CardNo"`

	SeqMode     string `url:"SeqMode"`
	CardSeq     string `url:"CardSeq"`
	DefaultFlag string `url:"DefaultFlag"`
	CardName    string `url:"CardName"`
	CardPass    string `url:"CardPass"`

	Expire       string `url:"Expired"`
	HolderName   string `url:"HolderName"`
	UpdateType   string `url:"UpdateType"`
	SecurityCode string `url:"SecurityCode"`
}

// Do excutes SaveCard operation.
func (svc *SaveCard) Do(cli client.Client) (*SaveCardResponse, error) {
	const apiPath = "/payment/SaveCard.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.SiteID = cli.Config.ShopID
	svc.BaseRequest.SitePass = cli.Config.SitePass

	result := &SaveCardResponse{}
	err := cli.Call(apiPath, svc, result)

	return result, err
}

type SaveCardResponse struct {
	CardSeq string
	CardNo  string
	Forward string
	client.BaseResponse
}
