package logrus

import (
	"github.com/leor-w/kid/logger"
	"github.com/sirupsen/logrus"
)

// KidHook logrus.Hook 的封装
type KidHook struct {
	logger.Worker
}

func (hook *KidHook) Levels() []logrus.Level {
	return hook.GetLevels()
}

func (hook *KidHook) Fire(entry *logrus.Entry) error {
	docs := hook.Doc()(entry)
	var err error
	go func() {
		err = hook.Exec()(docs)
	}()
	return err
}

func NewKidHook(worker logger.Worker) logger.Hook {
	return &KidHook{Worker: worker}
}
