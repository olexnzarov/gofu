package gofu_cli

import (
	"context"
	"fmt"
	"time"

	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/gofu"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InternalError(explanation string, cause error) error {
	if status, ok := status.FromError(cause); ok && status.Code() == codes.Unavailable {
		explanation = fmt.Sprintf("this error may indicate that the gofu daemon is not running, %s", explanation)
	}
	return fmt.Errorf("%s\nCause: %s", explanation, cause.Error())
}

func Error(explanation string, cause *pb.Error) error {
	return fmt.Errorf("%s\nCause: %s", explanation, cause.Message)
}

func Client() (*gofu.Client, error) {
	client, err := gofu.DefaultClient()
	if err != nil {
		return nil, InternalError("this error may indicate that the gofu daemon is not running", err)
	}
	return client, nil
}

func Timeout(timeout time.Duration) context.Context {
	context, _ := context.WithTimeout(context.Background(), timeout)
	return context
}
