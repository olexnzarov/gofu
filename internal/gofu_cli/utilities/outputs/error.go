package outputs

import (
	"fmt"

	"github.com/olexnzarov/gofu/internal/gofu_cli/utilities"
	"github.com/olexnzarov/gofu/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorOutput struct {
	cause       string
	description string
}

func Error(description string, cause interface{}) *ErrorOutput {
	var causeText string
	switch v := cause.(type) {
	case error:
		if status, ok := status.FromError(v); ok && status.Code() == codes.Unavailable {
			description = fmt.Sprintf("this error may indicate that the gofu daemon is not running, %s", description)
		}
		causeText = v.Error()
	case *pb.Error:
		causeText = v.Message
	case fmt.Stringer:
		causeText = v.String()
	case string:
		causeText = v
	default:
		causeText = "unknown"
	}
	utilities.SetExitCode(1)
	return &ErrorOutput{
		cause:       causeText,
		description: description,
	}
}

func Fatal(err interface{}) *ErrorOutput {
	return Error("", err)
}

func (e *ErrorOutput) Error() string {
	if e.description == "" {
		return fmt.Sprintf("Error: %s", e.cause)
	}
	return fmt.Sprintf("Error: %s\nCause: %s", e.description, e.cause)
}

func (e *ErrorOutput) String() string {
	return e.Error()
}

func (e *ErrorOutput) Text() string {
	return e.String()
}

func (e *ErrorOutput) Object() interface{} {
	return e.String()
}
