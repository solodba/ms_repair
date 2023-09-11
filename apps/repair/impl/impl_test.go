package impl_test

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/ms_repair/apps/repair"
	"github.com/solodba/ms_repair/test/tools"
)

var (
	svc repair.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(repair.AppName).(repair.Service)
}
