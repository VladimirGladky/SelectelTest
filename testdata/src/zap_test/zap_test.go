package zap_test

import "testing"

type Logger struct{}

func (l *Logger) Info(msg string, fields ...interface{})  {}
func (l *Logger) Error(msg string, fields ...interface{}) {}
func (l *Logger) Debug(msg string, fields ...interface{}) {}
func (l *Logger) Warn(msg string, fields ...interface{})  {}
func (l *Logger) Sync() error                             { return nil }

func TestZapLowercaseLetter(t *testing.T) {
	l := &Logger{}
	defer l.Sync()

	l.Info("Starting application")  // want "log message should start with lowercase letter"
	l.Error("Failed to initialize") // want "log message should start with lowercase letter"
	l.Debug("Processing request")   // want "log message should start with lowercase letter"
	l.Warn("Memory usage critical") // want "log message should start with lowercase letter"

	l.Info("starting application")
	l.Error("failed to initialize")
	l.Debug("processing request")
	l.Warn("memory usage critical")
}

func TestZapEnglishOnly(t *testing.T) {
	l := &Logger{}
	defer l.Sync()

	l.Info("начало работы") // want "log message must be in English only"
	l.Error("错误发生了")        // want "log message must be in English only"
	l.Debug("処理中です")        // want "log message must be in English only"
	l.Warn("경고 메시지")        // want "log message must be in English only"

	l.Info("starting work")
	l.Error("error occurred")
	l.Debug("processing")
	l.Warn("warning message")
}

func TestZapNoSpecialChars(t *testing.T) {
	l := &Logger{}
	defer l.Sync()

	l.Info("application started!")      // want "log message must not contain emojis or special characters"
	l.Error("critical error!!!")        // want "log message must not contain emojis or special characters"
	l.Debug("why is this happening?")   // want "log message must not contain emojis or special characters"
	l.Warn("attention: review needed")  // want "log message must not contain emojis or special characters"
	l.Info("task completed; next step") // want "log message must not contain emojis or special characters"
	l.Error("deployment failed🔥")       // want "log message must not contain emojis or special characters"

	l.Info("application started")
	l.Error("critical error")
	l.Debug("investigating issue")
	l.Warn("review needed")
}

func TestZapNoSensitiveData(t *testing.T) {
	l := &Logger{}
	defer l.Sync()

	l.Info("storing password")         // want "log message must not contain sensitive data"
	l.Debug("apikey configured")       // want "log message must not contain sensitive data"
	l.Error("token refresh failed")    // want "log message must not contain sensitive data"
	l.Warn("secret key rotated")       // want "log message must not contain sensitive data"
	l.Info("jwt verification started") // want "log message must not contain sensitive data"
	l.Debug("authorization granted")   // want "log message must not contain sensitive data"
	l.Error("private-key not loaded")  // want "log message must not contain sensitive data"

	l.Info("storing credentials")
	l.Debug("configuration loaded")
	l.Error("refresh operation failed")
	l.Warn("key rotated")
}
