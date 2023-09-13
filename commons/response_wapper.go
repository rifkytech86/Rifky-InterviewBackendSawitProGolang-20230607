package commons

import "github.com/labstack/echo/v4"

type Success struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Errors struct {
	Code    int       `json:"code"`
	Errors  ListError `json:"errors"`
	Message string    `json:"message"`
}
type ListError struct {
	ListError []string `json:"list_error"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	return c.JSON(statusCode, data)
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Errors{
		Code: statusCode,
		Errors: ListError{
			ListError: []string{message},
		},
	})
}

func ErrorResponses(c echo.Context, statusCode int, message string, listErrors []string) error {
	return Response(c, statusCode, Errors{
		Code: statusCode,
		Errors: ListError{
			ListError: listErrors,
		},
		Message: message,
	})
}

func SuccessResponse(c echo.Context, statusCode int, resp interface{}) error {
	return Response(c, statusCode, Success{
		Code:    statusCode,
		Message: "success",
		Data:    resp,
	})
}
