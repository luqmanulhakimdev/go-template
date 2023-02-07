package logger

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type severity string

const (
	severityDEBUG     severity = "DEBUG"
	severityINFO      severity = "INFO"
	severityNOTICE    severity = "NOTICE"
	severityWARNING   severity = "WARNING"
	severityERROR     severity = "ERROR"
	severityCRITICAL  severity = "CRITICAL"
	severityALERT     severity = "ALERT"
	severityEMERGENCY severity = "EMERGENCY"
)

var (
	levelsLogrusToGCE = map[logrus.Level]severity{
		logrus.DebugLevel: severityDEBUG,
		logrus.InfoLevel:  severityINFO,
		logrus.WarnLevel:  severityWARNING,
		logrus.ErrorLevel: severityERROR,
		logrus.FatalLevel: severityCRITICAL,
		logrus.PanicLevel: severityALERT,
	}
)

type GCEFormatter struct{}

func NewGCEFormatter() *GCEFormatter {
	return &GCEFormatter{}
}

func (f *GCEFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	data["time"] = entry.Time.Format(time.RFC3339Nano)
	data["severity"] = levelsLogrusToGCE[entry.Level]
	data["msg"] = entry.Message

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}
