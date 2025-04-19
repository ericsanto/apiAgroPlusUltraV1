package myerror

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func MessageErrorDuplicateKeyViolatesUniqueConstraint() string {

	return "duplicate key value violates unique constraint"

}

const DuplicateKeyErrorMessage = "duplicate key value violates unique constraint"

func NotFound(errorReturned error) error {

	if errors.Is(errorReturned, gorm.ErrRecordNotFound) {
		return fmt.Errorf("objeto não encontrado")
	}

	return fmt.Errorf("erro ao buscar objeto")
}

func IsUniqueConstraintViolated(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}

func IsViolatedForeingKeyConstraint(err error) bool {
	return strings.Contains(err.Error(), "(SQLSTATE 23503)")
}
