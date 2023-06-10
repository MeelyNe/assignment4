package delivery

import (
	"github.com/MeelyNe/assignment4/services/contact/internal/repository/controller"
	"github.com/MeelyNe/assignment4/services/contact/internal/repository/storage/postgres"
	use_case "github.com/MeelyNe/assignment4/services/contact/internal/use-case"
)

func (r *registry) NewGroupController() controller.Group {
	g := use_case.NewGroupUseCase(
		postgres.NewGroupRepository(r.db),
		postgres.NewDBRepository(r.db),
	)
	return controller.NewGroupsController(g)
}
