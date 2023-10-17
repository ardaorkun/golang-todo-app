package scripts

import "net/http"

const (
	ErrMsgTaskIDRequired         = "Task ID must be provided"
	ErrMsgTitleAndDescriptionReq = "Both title and description must be provided"
	ErrMsgNoValidUpdateData      = "No valid update data provided"
	ErrTaskNotFound              = "Task not found"
)

func HandleError(w http.ResponseWriter, err error, statusCode int) {
	if err != nil {
		http.Error(w, err.Error(), statusCode)
	}
}
