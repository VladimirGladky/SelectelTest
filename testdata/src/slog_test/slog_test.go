package slog_test

import (
	"log/slog"
	"testing"
)

func TestSlogLowercaseLetter(t *testing.T) {
	slog.Info("Starting server")    // want "log message should start with lowercase letter"
	slog.Error("Failed to connect") // want "log message should start with lowercase letter"
	slog.Debug("Processing data")   // want "log message should start with lowercase letter"
	slog.Warn("Memory is high")     // want "log message should start with lowercase letter"

	slog.Info("starting server")
	slog.Error("failed to connect")
	slog.Debug("processing data")
	slog.Warn("memory is high")
}

func TestSlogEnglishOnly(t *testing.T) {
	slog.Info("запуск сервера")     // want "log message must be in English only"
	slog.Error("ошибка соединения") // want "log message must be in English only"
	slog.Debug("处理数据")              // want "log message must be in English only"
	slog.Warn("メモリ不足")              // want "log message must be in English only"

	slog.Info("starting server")
	slog.Error("connection error")
	slog.Debug("processing data")
	slog.Warn("low memory")
}

func TestSlogNoSpecialChars(t *testing.T) {
	slog.Info("server started!")       // want "log message must not contain emojis or special characters"
	slog.Error("connection failed!!!") // want "log message must not contain emojis or special characters"
	slog.Debug("what is this?")        // want "log message must not contain emojis or special characters"
	slog.Warn("warning: check this")   // want "log message must not contain emojis or special characters"
	slog.Info("done; moving forward")  // want "log message must not contain emojis or special characters"
	slog.Error("server crashed🚀")      // want "log message must not contain emojis or special characters"

	slog.Info("server started")
	slog.Error("connection failed")
	slog.Debug("processing complete")
	slog.Warn("check this value")
}

func TestSlogNoSensitiveData(t *testing.T) {
	slog.Info("user password saved")      // want "log message must not contain sensitive data"
	slog.Debug("api_key received")        // want "log message must not contain sensitive data"
	slog.Error("token validation failed") // want "log message must not contain sensitive data"
	slog.Warn("secret not found")         // want "log message must not contain sensitive data"
	slog.Info("jwt token expired")        // want "log message must not contain sensitive data"
	slog.Debug("bearer auth used")        // want "log message must not contain sensitive data"
	slog.Error("private_key missing")     // want "log message must not contain sensitive data"

	slog.Info("user authenticated")
	slog.Debug("request received")
	slog.Error("validation failed")
	slog.Warn("configuration not found")
}
