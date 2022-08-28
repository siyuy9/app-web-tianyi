package presenter

import (
	"fmt"
	"net/http"

	pkgError "gitlab.com/kongrentian-group/tianyi/v1/pkg/error"
)

func InvalidProjectID(err error) error {
	return pkgError.NewWithCode(
		fmt.Errorf("invalid project id: %w", err), http.StatusBadRequest,
	)
}

func CouldNotFindProject(err error) error {
	return pkgError.NewWithCode(
		fmt.Errorf("could not find project(s): %w", err),
		http.StatusNotFound,
	)
}

func CouldNotUpdateProject(err error) error {
	return pkgError.New(fmt.Errorf("could not update project: %w", err))
}

func CouldNotCreateProject(err error) error {
	return pkgError.New(fmt.Errorf("could not create project: %w", err))
}
