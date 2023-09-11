package impl_test

import (
	"testing"
)

func TestQuerySlaveError(t *testing.T) {
	res, err := svc.QuerySlaveError(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestQuerySlaveMode(t *testing.T) {
	res, err := svc.QuerySlaveMode(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestBasePositionRepair(t *testing.T) {
	err := svc.BasePositionRepair(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuerySlaveExcuteGtid(t *testing.T) {
	res, err := svc.QuerySlaveExcuteGtid(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestBaseGtidRepair(t *testing.T) {
	err := svc.BaseGtidRepair(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
