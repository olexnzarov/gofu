package utilities

import (
	"context"
	"fmt"

	"github.com/olexnzarov/gofu/pkg/gofu"
)

func Client() (*gofu.Client, error) {
	client, err := gofu.DefaultClient()
	if err != nil {
		return nil, fmt.Errorf("this error may indicate that the gofu daemon is not running, %s", err)
	}
	return client, nil
}

func Timeout() (context.Context, context.CancelFunc) {
	context, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	return context, cancel
}
