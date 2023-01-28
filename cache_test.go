package cache

import (
	"errors"
	"fmt"
	"syscall"
	"testing"
	"time"
)

func MockGenerateData() (map[string]string, error) {
	m := make(map[string]string)
	m["1"] = "one"
	m["2"] = "two"
	fmt.Println("Refreshing")
	return m, nil
}

func MockGenerateDataWithErr() (map[string]string, error) {
	return nil, errors.New("error error error occurring")
}

func TestCache(t *testing.T) {
	//_ = New[string, string](MockGenerateDataWithErr, LogErr, RefreshOnStart[string, string], RefreshWithCron[string, string]("@every 1s"))
	//_ = New[string, string](MockGenerateDataWithErr, PanicErr, RefreshOnStart[string, string], RefreshWithCron[string, string]("@every 1s"))
	//_ = New[string, string](MockGenerateDataWithErr, FatalErr, RefreshOnStart[string, string], RefreshWithCron[string, string]("@every 1s"))

	cache := New[string, string](MockGenerateData, PanicErr, RefreshOnStart[string, string], RefreshWithCron[string, string]("@every 1s"))

	if v, ok := cache.Get("1"); !ok {
		t.Errorf("Got %v", v)
	}

	time.Sleep(time.Second * 2)

	_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}
