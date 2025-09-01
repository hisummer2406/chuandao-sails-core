package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestNewUserLogic(t *testing.T) {
	ctx := context.Background()
	logx.WithContext(ctx).Infof("hello world")
	logc.Info(ctx, "hello world")
}
