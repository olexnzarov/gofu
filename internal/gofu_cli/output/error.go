package output

import (
	"fmt"

	"github.com/olexnzarov/gofu/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func formatErrorMessage(explanation string, cause string) string {
	if explanation == "" {
		return fmt.Sprintf("Error: %s", cause)
	}
	return fmt.Sprintf("Error: %s\nCause: %s", explanation, cause)
}

func (o *Output) DaemonError(explanation string, cause *pb.Error) *Output {
	return o.Add("error", Text(formatErrorMessage(explanation, cause.Message)))
}

func (o *Output) Error(explanation string, cause error) *Output {
	if status, ok := status.FromError(cause); ok && status.Code() == codes.Unavailable {
		explanation = fmt.Sprintf("this error may indicate that the gofu daemon is not running, %s", explanation)
	}
	return o.Add("error", Text(formatErrorMessage(explanation, cause.Error())))
}

func (o *Output) Fail(cause error) *Output {
	return o.Error("", cause)
}
