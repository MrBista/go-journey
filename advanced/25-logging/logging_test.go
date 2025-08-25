package logging

import (
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogDasar(t *testing.T) {

	logger := logrus.New()

	// logger.Println("Hello dunia")

	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Trace")
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")

}

func TestOut(t *testing.T) {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	multiWriter := io.MultiWriter(os.Stdout, file)

	logger.SetOutput(multiWriter)

	logger.Info("Masuk file")
}

func TestFormater(t *testing.T) {
	logger := logrus.New()
	// logger.SetFormatter(&logrus.TextFormatter{}) // default
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello World")
}

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("field.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	multiWriter := io.MultiWriter(os.Stdout, file)

	logger.SetOutput(multiWriter)

	logger.WithField("username", "Gusti Bisman Taka").Info("Logger info message")
}

func TestFields(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	multiWriter := io.MultiWriter(os.Stdout, file)

	logger.SetOutput(multiWriter)

	logger.SetFormatter(&logrus.JSONFormatter{})

	// fields itu alias untuk map[string]interface{}
	logger.WithFields(logrus.Fields{
		"nama":  "Bismoy",
		"umur":  24,
		"hobby": "Lari",
	},
	).Infof("Logger with fields")
}
