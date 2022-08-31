package service

type User struct {
	Nama string `json:"name"`
}

type UserService struct {
	ListUser []User
}

type UserIface interface {
	Register(u *User)
	Get() []User
}

func NewUserService(user []User) UserIface {
	return &UserService{
		ListUser: user,
	}
}

func (u *UserService) Register(user *User) {
	u.ListUser = append(u.ListUser, *user)
}

func (u *UserService) Get() ([]User){
	// for _, each := range u.ListUser {
	// 	fmt.Println(each.Nama)
	// }
	return u.ListUser
}