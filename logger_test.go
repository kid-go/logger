package logger

import (
	"github.com/leor-w/logger/logrus/rotate"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	NewLogger(nil)
	// 日志文件分割 hook
	rotateHook, _ := rotate.NewSimpleRotate()
	AddHook(rotateHook)
	// 日志上传 es hook
	//AddWorker(es_worker2.NewEsWorker(
	//	es_worker2.WithEsAddress("http://127.0.0.1:9201"),
	//	es_worker2.WithLogLevel(logger.WarnLevel.String()),
	//))

	Error("error log test")
	time.Sleep(10 * time.Second)
}
