package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"
	"github.com/tqtcloud/keyauth/common/utils"
	"net/http"
	"time"
)

const (
	AppName       = "user"
	DefaultDomain = "default"
)

var validate = validator.New()

// 保存Hash过后的Password
func NewUser(req *CreateUserRequest) *User {
	return &User{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Data:     req,
	}
}

// CheckPassword 检查密码
func (u *User) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Data.Password)
}

// Validate 审计输入的json
func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

func (s *UserSet) Add(item *User) {
	s.Items = append(s.Items, item)
}

func NewDefaultUser() *User {
	return NewUser(NewCreateUserRequest())
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Domain: DefaultDomain,
	}
}

func NewQueryUserRequestFromHTTP(r *http.Request) *QueryUserRequest {
	qs := r.URL.Query()

	return &QueryUserRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: qs.Get("keywords"),
	}
}

func NewPutUserRequest(id string) *UpdateUserRequest {
	return &UpdateUserRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateUserRequest(),
	}
}

func NewPatchUserRequest(id string) *UpdateUserRequest {
	return &UpdateUserRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateUserRequest(),
	}
}

func NewDeleteUserRequestWithID(id string) *DeleteUserRequest {
	return &DeleteUserRequest{
		Id: id,
	}
}

func NewDescribeUserRequestById(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DescribeBy_USER_ID,
		UserId:     id,
	}
}

func NewDescribeUserRequestByName(domain, name string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DescribeBy_USER_NAME,
		Domain:     domain,
		UserName:   name,
	}
}
