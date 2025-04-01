package models

import (
	proto "github.com/garden-raccoon/sms-pkg/protocols/sms-pkg"
	"github.com/gofrs/uuid"
)

type Sms struct {
	SmsUuid       uuid.UUID `json:"sms_uuid"`
	UserUuid      uuid.UUID `json:"user_uuid"`
	ApiSmsUuid    uuid.UUID `json:"api_sms_uuid"`
	IsApproved    bool      `json:"is_approved"`
	IsDisapproved bool      `json:"is_disapproved"`
	CheckPhrase   string    `json:"check_phrase"`
	Phone         string    `json:"phone"`
}

// Proto is
func ToProto(s Sms) *proto.Sms {

	sms := &proto.Sms{
		SmsUuid:       s.SmsUuid.Bytes(),
		UserUuid:      s.UserUuid.Bytes(),
		ApiSmsUuid:    s.ApiSmsUuid.Bytes(),
		IsApproved:    s.IsApproved,
		IsDisapproved: s.IsDisapproved,
		CheckPhrase:   s.CheckPhrase,
		Phone:         s.Phone,
	}
	return sms
}

func FromProto(pb *proto.Sms) *Sms {

	sms := &Sms{
		SmsUuid:       uuid.FromBytesOrNil(pb.SmsUuid),
		UserUuid:      uuid.FromBytesOrNil(pb.UserUuid),
		ApiSmsUuid:    uuid.FromBytesOrNil(pb.ApiSmsUuid),
		IsApproved:    pb.IsApproved,
		IsDisapproved: pb.IsDisapproved,
		CheckPhrase:   pb.CheckPhrase,
		Phone:         pb.Phone,
	}
	return sms
}
