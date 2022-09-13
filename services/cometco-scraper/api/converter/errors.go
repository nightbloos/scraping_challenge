package converter

import (
	"net/http"

	"github.com/pkg/errors"
	"scraping_challenge/common/service/domain"

	apiModel "scraping_challenge/services/cometco-scraper/api/model"
)

func FromError(err error) apiModel.ErrorResponse {
	var errResponse apiModel.ErrorResponse
	errResponse.Code = http.StatusInternalServerError
	errResponse.Message = err.Error()

	var domainErr domain.Error
	if errors.As(err, &domainErr) {
		switch domainErr.GetErrorType() {
		case domain.ErrBadRequest:
			errResponse.Code = http.StatusBadRequest
		case domain.ErrUnauthorized:
			errResponse.Code = http.StatusUnauthorized
		case domain.ErrNotFound:
			errResponse.Code = http.StatusNotFound
		case domain.ErrForbidden:
			errResponse.Code = http.StatusForbidden
		case domain.ErrInternal:
			errResponse.Code = http.StatusInternalServerError
		}
	}

	return errResponse
}
