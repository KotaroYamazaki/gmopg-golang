package card

import "gmopg-card/client"

// ExecTran is struct of ExecTran API
// Excution for payment by card.
type ExecTran struct {
	client.BaseRequest `url:",squash"`

	// REQUIRED
	AcceptID   string `url:"OrderID"`
	AccessID   string `url:"AccessID"`
	AccessPass string `url:"AccessPass"`
	Method     string `url:"Method"`
	MemberID   string `url:"MemberID"`
	CardSeq    string `url:"CardSeq"`
	CardNo     string `url:"CardNo"`
	// OPTIONAL
	PayTimes       string `url:"PayTimes"`
	SeqMode        string `url:"SeqMode"`
	CardPass       string `url:"CardPass"`
	HTTPAccept     string `url:"HttpAccept"`
	HTTPUserAgent  string `url:"HttpUserAgent"`
	DeviceCategory string `url:"DeviceCategory"`
	SecurityCode   string `url:"SecurityCode"`
	ClientField1   string `url:"ClientField1,omitempty"`
	ClientField2   string `url:"ClientField2,omitempty"`
	ClientField3   string `url:"ClientField3,omitempty"`
}

// Do excutes ExecTran operation.
func (svc *ExecTran) Do(cli client.Client) (*ExecTranResponse, error) {

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.ShopID = cli.Config.ShopID
	svc.BaseRequest.ShopPass = cli.Config.ShopPass

	return nil, nil
}

// ExecTranResponse is struct for response of ExecTran API.
type ExecTranResponse struct {
	client.BaseResponse `url:",squash"`

	OrderID      string `url:"OrderID"`
	Forward      string `url:"Forward"`
	Method       string `url:"Method"`
	PayTimes     string `url:"PayTimes"`
	Approve      string `url:"Approve"`
	TranID       string `url:"TranID"`
	TranDate     string `url:"TranDate"`
	CheckString  string `url:"CheckString"`
	ClientField1 string `url:"ClientField1"`
	ClientField2 string `url:"ClientField2"`
	ClientField3 string `url:"ClientField3"`
}

// IsSuccess checks the request is success or not
func (r *ExecTranResponse) IsSuccess() bool {
	switch {
	case !r.BaseResponse.IsSuccess():
		return false
	}
	return true
}
