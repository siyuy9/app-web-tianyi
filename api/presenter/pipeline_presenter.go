package presenter

import (
	"fmt"
	"net/http"

	pkgError "gitlab.com/kongrentian-group/tianyi/v1/pkg/error"
)

func CouldNotFindPileline(err error) error {
	return pkgError.NewWithCode(
		fmt.Errorf("could not find project pipeline(s): %w", err),
		http.StatusNotFound,
	)
}

func CouldNotSchedulePipelines(err error) error {
	return pkgError.NewWithCode(
		fmt.Errorf("could not schedule pipeline(s): %w", err),
		http.StatusInternalServerError,
	)
}
