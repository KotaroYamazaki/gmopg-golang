package card

import "gmopg-card/client"

// SaveMember is struct for SaveMember API.
type SaveMember struct {
	client.BaseRequest `url:",squash"`

	// REQUIRED
	MemberID string `url:"MemberID"`

	// OPTIONAL
	Version    string `url:"Version"`
	MemberName string `url:"MemberName"`
}

func (svc *SaveMember) Do(cli client.Client) (*SaveMemberResponse, error) {
	const apiPath = "/payment/SaveMember.idPass"

	svc.BaseRequest.Version = cli.Config.Version
	svc.BaseRequest.SiteID = cli.Config.SiteID
	svc.BaseRequest.SitePass = cli.Config.SitePass

	result := &SaveMemberResponse{}
	err := cli.Call(apiPath, svc, result)
	return result, err
}

type SaveMemberResponse struct {
	MemberID            string `url:"MemberID"`
	client.BaseResponse `url:",squash"`
}

// IsSuccess checks the request is success or not
func (r *SaveMemberResponse) IsSuccess() bool {
	switch {
	case !r.BaseResponse.IsSuccess():
		return false
	}
	return true
}
