package delivery

import (
	"github.com/MeelyNe/assignment4/services/contact/internal/repository/controller"
	"github.com/MeelyNe/assignment4/services/contact/internal/repository/storage/postgres"
	use_case "github.com/MeelyNe/assignment4/services/contact/internal/use-case"
)

func (r *registry) NewContactController() controller.Contact {
	c := use_case.NewContactUseCase(
		postgres.NewContactRepository(r.db),
		postgres.NewDBRepository(r.db),
	)
	return controller.NewContactsController(c)
}
