package service

type Service struct {
	UserSrv IUserSrv
	HubSrv  IHubSrv
}

func ProvideServices(
	userSrv IUserSrv,
	hubSrv IHubSrv,
) *Service {
	return &Service{
		UserSrv: userSrv,
		HubSrv:  hubSrv,
	}
}
