package log

import (
	"github.com/AnnonaOrg/osenv"
	"github.com/sirupsen/logrus"
)

func init() {
	logLevel, _ := logrus.ParseLevel(osenv.GetLogLevel())
	logrus.SetLevel(logLevel)
	// logrus.SetReportCaller(true)
	// Set the formatter to include file and line information
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05 MST",
		// CallerPrettyfier: func(f *runtime.Frame) (string, string) {
		// 	return "", fmt.Sprintf("%s:%d", f.File, f.Line)
		// },
	})
}
