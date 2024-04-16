package log

import (
	"encoding/json"
	"myapp/libs/masking"
	"reflect"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
}

var L = NewLog()

type Log struct {
	TraceID string
	SpanID  string
}

func NewLog() Logger {
	return &Log{
		TraceID: uuid.NewString(),
		SpanID:  uuid.NewString(),
	}
}

func NewChildLog(log Logger) Logger {
	l := log.(*Log)
	return &Log{
		TraceID: l.TraceID,
		SpanID:  uuid.NewString(),
	}
}

func NewLogForceTraceID(traceID string) Logger {
	return &Log{
		TraceID: traceID,
		SpanID:  uuid.NewString(),
	}
}

func (l *Log) Info(obj any) {
	log.Info().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf("%v", obj)
}

func (l *Log) Infof(format string, obj ...any) {
	log.Info().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf(format, obj...)
}

func (l *Log) Debug(obj any) {
	log.Debug().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf("%v", obj)
}

func (l *Log) Debugf(format string, obj any) {
	log.Debug().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf(format, obj)
}

func (l *Log) Warn(obj any) {
	log.Warn().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf("%v", obj)
}

func (l *Log) Warnf(format string, obj any) {
	log.Warn().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf(format, obj)
}

func (l *Log) Error(obj any) {
	log.Error().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf("%v", obj)
}

func (l *Log) Errorf(format string, obj any) {
	log.Error().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf(format, obj)
}

func (l *Log) Fatal(obj any) {
	log.Fatal().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf("%v", obj)
}

func (l *Log) Fatalf(format string, obj any) {
	log.Fatal().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf(format, obj)
}

func (l *Log) Panic(obj any) {
	log.Panic().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf("%v", obj)
}

func (l *Log) Panicf(format string, obj any) {
	log.Panic().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		Msgf(format, obj)
}

var typeOfArrByte = reflect.TypeOf([]byte(nil))

func objToByte(obj any) []byte {
	var b []byte
	if reflect.ValueOf(obj).Type() == typeOfArrByte {
		b = obj.([]byte)
	} else {
		b, _ = json.Marshal(obj)
	}

	m := map[string]any{}
	json.Unmarshal(b, &m)

	masking.MaskMap(m)
	b, _ = json.Marshal(m)

	return b
}

func (l *Log) LogHttpRequest(msg string, headerObj, bodyObj any) {
	body := objToByte(bodyObj)
	header := objToByte(headerObj)

	log.Info().
		Str("traceId", l.TraceID).
		Str("spanId", l.SpanID).
		RawJSON("header", header).
		RawJSON("body", body).
		Msg(msg)
}
