package repositoryimpl

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/opensourceways/app-cla-server/signing/domain"
	"github.com/opensourceways/app-cla-server/signing/domain/dp"
)

const (
	fieldAccount  = "account"
	fieldChanged  = "changed"
	fieldPassword = "password"
)

// userDO
type userDO struct {
	Id             primitive.ObjectID `bson:"_id"       json:"-"`
	Email          string             `bson:"email"     json:"email"     required:"true"`
	LinkId         string             `bson:"link_id"   json:"link_id"   required:"true"`
	Account        string             `bson:"account"   json:"account"   required:"true"`
	Password       string             `bson:"password"  json:"password"  required:"true"`
	CorpSigningId  string             `bson:"cs_id"     json:"cs_id"     required:"true"`
	PasswordChaged bool               `bson:"changed"   json:"changed"`
	Version        int                `bson:"version"   json:"-"`
}

func (do *userDO) toDoc() (bson.M, error) {
	return genDoc(do)
}

func (do *userDO) toUser(u *domain.User) (err error) {
	if u.EmailAddr, err = dp.NewEmailAddr(do.Email); err != nil {
		return
	}

	if u.Account, err = dp.NewAccount(do.Account); err != nil {
		return
	}

	if u.Password, err = dp.NewPassword(do.Password); err != nil {
		return
	}

	u.Id = do.Id.Hex()
	u.LinkId = do.LinkId
	u.CorpSigningId = do.CorpSigningId
	u.PasswordChaged = do.PasswordChaged
	u.Version = do.Version

	return
}

func toUserDO(u *domain.User) userDO {
	return userDO{
		Email:          u.EmailAddr.EmailAddr(),
		LinkId:         u.LinkId,
		Account:        u.Account.Account(),
		Password:       u.Password.Password(),
		CorpSigningId:  u.CorpSigningId,
		PasswordChaged: u.PasswordChaged,
	}
}