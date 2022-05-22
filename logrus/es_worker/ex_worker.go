package es_worker

import (
	"context"
	"fmt"
	"github.com/leor-w/kid/logger"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"reflect"
	"time"
)

type EsWorker struct {
	opts   Options
	client *elastic.Client
}

func (worker *EsWorker) Exec() logger.ExecFunc {
	return func(docs interface{}) error {
		doc, ok := docs.(map[string]interface{})
		if !ok {
			return fmt.Errorf("logger.EsWorker: failed doc type %v", reflect.TypeOf(doc).Name())
		}
		_, err := worker.client.Index().Index(worker.opts.indexName()).Type("_doc").BodyJson(doc).Do(context.Background())
		fmt.Println(err)
		if err != nil {
			return err
		}
		return nil
	}
}

func (worker *EsWorker) Doc() logger.DocFunc {
	return func(entry *logrus.Entry) interface{} {
		doc := make(map[string]interface{})
		for k, v := range entry.Data {
			doc[k] = v
		}
		doc["time"] = time.Now().Local()
		doc["lvl"] = entry.Level
		doc["message"] = entry.Message
		doc["caller"] = fmt.Sprintf("%s:%d %#v", entry.Caller.File, entry.Caller.Line, entry.Caller.Func)
		return doc
	}
}

func (worker *EsWorker) GetLevels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

type Option func(*Options)

func NewEsWorker(options ...Option) *EsWorker {
	opts := Options{
		logLevel:   logger.InfoLevel.String(),
		esAddress:  []string{},
		esUser:     "",
		esPassword: "",
		cmd:        "",
		indexName: func() string {
			return fmt.Sprintf("%s_doc", time.Now().Format("2006-01-02-15:04:05"))
		},
		health: time.Second * 15,
	}
	for _, o := range options {
		o(&opts)
	}

	if len(opts.esAddress) <= 0 {
		log.Fatal("logger.EsWorker: elasticSearch address is required")
	}

	h := &EsWorker{
		opts: opts,
	}
	es, err := elastic.NewClient(
		elastic.SetURL(h.opts.esAddress...),
		elastic.SetBasicAuth(h.opts.esUser, h.opts.esPassword),
		elastic.SetSniff(false),
		elastic.SetHealthcheckTimeout(h.opts.health),
		elastic.SetErrorLog(log.New(os.Stderr, "ES:", log.LstdFlags)),
	)
	if err != nil {
		log.Fatal("failed to create ElasticSearch v6 client: ", err.Error())
	}
	h.client = es
	return h
}
