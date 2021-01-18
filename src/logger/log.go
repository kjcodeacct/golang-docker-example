package logger

import (
	"errors"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLogger *zap.Logger
var logSink *lumberjack.Logger

const LogToFile = 1
const LogToStdOut = 0

func Get() (*zap.Logger, error) {

	if zapLogger == nil {
		return nil, errors.New("Please initilize logging")
	}

	return zapLogger, nil
}

func Setup(logMethod int, logDir string) error {

	logFileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logDir,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"

	encoderCfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	if logMethod == LogToStdOut {
		zapLogger = zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			atom,
		))
	} else {
		zapLogger = zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			logFileWriter,
			atom,
		))
	}

	atom.SetLevel(zap.DebugLevel)

	// flushes zap buffer, if any
	defer zapLogger.Sync()

	zap.RedirectStdLog(zapLogger)

	log.Printf("saving logs to %s\n", logDir)
	return nil
}
