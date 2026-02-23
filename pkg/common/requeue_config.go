package common

import (
	"os"
	"strconv"
	"time"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var configLog = logf.Log.WithName("requeue_config")

const (
	// Environment variable names for configuring requeue delays
	EnvRequeueDelay      = "REQUEUE_DELAY_SECONDS"
	EnvRequeueDelayError = "REQUEUE_DELAY_ERROR_SECONDS"

	// Default values
	DefaultRequeueDelaySeconds      = 300 // 5 minutes
	DefaultRequeueDelayErrorSeconds = 30
)

// GetRequeueDelay returns the normal requeue delay from the REQUEUE_DELAY_SECONDS
// env var, or the provided default if not set.
func GetRequeueDelay(defaultSeconds int) time.Duration {
	return getDurationFromEnv(EnvRequeueDelay, defaultSeconds)
}

// GetRequeueDelayError returns the error requeue delay from the REQUEUE_DELAY_ERROR_SECONDS
// env var, or the provided default if not set.
func GetRequeueDelayError(defaultSeconds int) time.Duration {
	return getDurationFromEnv(EnvRequeueDelayError, defaultSeconds)
}

func getDurationFromEnv(envVar string, defaultSeconds int) time.Duration {
	val := os.Getenv(envVar)
	if val == "" {
		return time.Duration(defaultSeconds) * time.Second
	}

	seconds, err := strconv.Atoi(val)
	if err != nil {
		configLog.Error(err, "invalid value for env var, using default", "env", envVar, "value", val, "default", defaultSeconds)
		return time.Duration(defaultSeconds) * time.Second
	}

	configLog.Info("using configured requeue delay", "env", envVar, "seconds", seconds)
	return time.Duration(seconds) * time.Second
}
