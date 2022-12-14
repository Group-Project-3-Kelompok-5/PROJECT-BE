package factory

import (
	deliveryAuth "be13/project/features/auth/delivery"
	repoAuth "be13/project/features/auth/repository"
	serviceAuth "be13/project/features/auth/service"
	deliveryCheck "be13/project/features/check/delivery"
	repoCheck "be13/project/features/check/repository"
	serviceCheck "be13/project/features/check/service"
	deliveryComment "be13/project/features/comment/delivery"
	repoComment "be13/project/features/comment/repository"
	serviceComment "be13/project/features/comment/service"
	deliveryHome "be13/project/features/homestay/delivery"
	repoHome "be13/project/features/homestay/repository"
	serviceHome "be13/project/features/homestay/service"
	deliveryReserve "be13/project/features/reservation/delivery"
	repoReserve "be13/project/features/reservation/repository"
	serviceReserve "be13/project/features/reservation/service"
	userDelivery "be13/project/features/user/delivery"
	userRepo "be13/project/features/user/repository"
	userService "be13/project/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepofaktory := userRepo.New(db) //menginiasialisasi func new yang ada di repository
	userServiceFaktory := userService.New(userRepofaktory)
	userDelivery.New(userServiceFaktory, e)

	authRepoFactory := repoAuth.NewAuth(db)
	authServiceFactory := serviceAuth.NewAuth(authRepoFactory)
	deliveryAuth.NewAuth(authServiceFactory, e)

	homeRepoFactory := repoHome.NewHome(db)
	homeServiceFactory := serviceHome.NewHome(homeRepoFactory)
	deliveryHome.NewHome(homeServiceFactory, e)

	commentRepoFactory := repoComment.NewComment(db)
	commentServiceFactory := serviceComment.NewComment(commentRepoFactory)
	deliveryComment.NewComment(commentServiceFactory, e)

	checkRepoFactory := repoCheck.NewCheck(db)
	checkServiceFactory := serviceCheck.NewCheck(checkRepoFactory)
	deliveryCheck.NewCheck(checkServiceFactory, e)

	reserveRepoFactory := repoReserve.NewRes(db)
	reserveServiceFactory := serviceReserve.NewRes(reserveRepoFactory)
	deliveryReserve.NewRes(reserveServiceFactory, e)

}
