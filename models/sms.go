package models

import (
	proto "github.com/garden-raccoon/sms-pkg/protocols/sms-pkg"
	"github.com/gofrs/uuid"
)

type Sms struct {
	SmsUuid       uuid.UUID `json:"sms_uuid"`
	UserUuid      uuid.UUID `json:"user_uuid"`
	IsApproved    bool      `json:"is_approved"`
	IsDisapproved bool      `json:"is_disapproved"`
	CheckPhrase   string    `json:"check_phrase"`
}

// Proto is
func ToProto(s Sms) *proto.Sms {

	sms := &proto.Sms{
		SmsUuid:       s.SmsUuid.Bytes(),
		UserUuid:      s.UserUuid.Bytes(),
		IsApproved:    s.IsApproved,
		IsDisapproved: s.IsDisapproved,
		CheckPhrase:   s.CheckPhrase,
	}
	return sms
}

func FromProto(pb *proto.Sms) *Sms {

	sms := &Sms{
		SmsUuid:       uuid.FromBytesOrNil(pb.SmsUuid),
		UserUuid:      uuid.FromBytesOrNil(pb.UserUuid),
		IsApproved:    pb.IsApproved,
		IsDisapproved: pb.IsDisapproved,
		CheckPhrase:   pb.CheckPhrase,
	}
	return sms
}
