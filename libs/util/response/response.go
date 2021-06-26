package response

import (
	"fmt"
	"net/http"
	"simple-api/api/model"

	"github.com/labstack/echo/v4"
)

type MetaResponse struct {
	Success bool                     `json:"success" default:"true"`
	Message string                   `json:"message" default:"true"`
	Info    *model.PaginationInfoDTO `json:"info"`
}

type SuccessResponse struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data"`
}

type ErrorResponse struct {
	Meta  MetaResponse `json:"meta"`
	Error string       `json:"error"`
}

type ErrorConstant struct {
	Response     ErrorResponse
	Code         int
	ErrorMessage error
}

func (r *ErrorConstant) Error() string {
	return fmt.Sprintf("error code %d", r.Code)
}

func (r *ErrorConstant) Builder() *ErrorConstant {
	return r
}

const (
	E_DUPLICATE            = "duplicate"
	E_NOT_FOUND            = "not_found"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
	E_UNAUTHORIZED         = "unauthorized"
	E_BAD_REQUEST          = "bad_request"
	E_SERVER_ERROR         = "server_error"
)

var (
	ErrDuplicate = ErrorConstant{
		Response: ErrorResponse{
			Meta: MetaResponse{
				Success: false,
				Message: "Created value already exists",
			},
			Error: E_DUPLICATE,
		},
		Code: http.StatusConflict,
	}
	ErrDataNotFound = ErrorConstant{
		Response: ErrorResponse{
			Meta: MetaResponse{
				Success: false,
				Message: "Data not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	}
	ErrRouteNotFound = ErrorConstant{
		Response: ErrorResponse{
			Meta: MetaResponse{
				Success: false,
				Message: "Route not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	}
	ErrUnprocessableEntity = ErrorConstant{
		Response: ErrorResponse{
			Meta: MetaResponse{
				Success: false,
				Message: "Invalid parameters or payload",
			},
			Error: E_UNPROCESSABLE_ENTITY,
		},
		Code: http.StatusUnprocessableEntity,
	}
	ErrUnauthorized = ErrorConstant{
		Response: ErrorResponse{
			Meta: MetaResponse{
				Success: false,
				Message: "Unauthorized, please login",
			},
			Error: E_UNAUTHORIZED,
		},
		Code: http.StatusUnauthorized,
	}
	ErrBadRequest = ErrorConstant{
		Response: ErrorResponse{
			Meta: MetaResponse{
				Success: false,
				Message: "Bad Request",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	}
	ErrValidation = ErrorConstant{
		Response: ErrorResponse{
			Meta: MetaResponse{
				Success: false,
				Message: "Invalid parameters or payload",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	}
	ErrServerError = ErrorConstant{
		Response: ErrorResponse{
			Meta: MetaResponse{
				Success: false,
				Message: "Something bad happened",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	}
)

func BuildError(err ErrorConstant, msg error) error {
	err.ErrorMessage = msg
	return &err
}

func BuildErrorCustom(code int, error string, message string) error {
	return &ErrorConstant{
		Response: ErrorResponse{
			Meta: MetaResponse{
				Success: false,
				Message: message,
			},
			Error: error,
		},
		Code: code,
	}
}

func RespSuccess(c echo.Context, message string, data interface{}) error {
	response := SuccessResponse{
		Meta: MetaResponse{
			Success: true,
			Message: message,
		},
		Data: data,
	}
	return c.JSON(http.StatusOK, response)
}

func RespSuccessInfo(c echo.Context, message string, data interface{}, pagination *model.PaginationInfoDTO) error {
	response := SuccessResponse{
		Meta: MetaResponse{
			Success: true,
			Message: message,
			Info:    pagination,
		},
		Data: data,
	}
	return c.JSON(http.StatusOK, response)
}

func RespError(c echo.Context, err error) error {
	re, ok := err.(*ErrorConstant)
	if ok {
		return c.JSON(re.Builder().Code, re.Builder().Response)
	} else {
		return c.JSON(ErrServerError.Code, ErrServerError.Response)
	}
}
