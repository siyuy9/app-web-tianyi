package presenter

import (
	"fmt"
	"net/http"

	pkgError "gitlab.com/kongrentian-group/tianyi/v1/pkg/error"
)

func CouldNotMigrateDatabase(err error) error {
	return pkgError.NewWithCode(
		fmt.Errorf("could not migrate database: %w", err),
		http.StatusInternalServerError,
	)
}
