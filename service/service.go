package service

type Service struct {
	UserSrv IUserSrv
}

func ProvideService(
	userSrv IUserSrv,
) *Service {
	return &Service{
		UserSrv: userSrv,
	}
}
