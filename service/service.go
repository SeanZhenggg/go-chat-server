package service

type Service struct {
	UserSrv IUserSrv
}

func ProvideServices(
	userSrv IUserSrv,
) *Service {
	return &Service{
		UserSrv: userSrv,
	}
}
