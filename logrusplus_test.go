package logrusplus_test

import (
	"github.com/jyiigpgf/logrusplus"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func Test_Main(t *testing.T) {
	lrp := logrusplus.New()
	p2pLogger := lrp.Logger("p2p")
	tvmLogger := lrp.Logger("vm")
	stdLogger := lrp.StandardLogger()
	commonLogger := lrp.CommonLogger()

	for {
		go func() {
			p2pLogger.WithFields(logrus.Fields{
				"test": "p2p",
			}).Info("hello world")
		}()

		go func() {
			tvmLogger.WithFields(logrus.Fields{
				"test": "vm",
			}).Info("hello world")
		}()

		go func() {
			stdLogger.WithFields(logrus.Fields{
				"test": "std",
			}).Info("hello world")
		}()

		go func() {
			commonLogger.WithFields(logrus.Fields{
				"test": "common",
			}).Info("hello world")
		}()

		time.Sleep(1 * time.Second)
	}
}
