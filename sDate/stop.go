package sDate

import (
	"time"

	"github.com/yasseldg/simplego/sLog"
)

// NextStop, ts and interval in seconds
func NextStop(ts, interval int64) int64 {
	diff := ts % interval
	return ts - diff + interval - interval/10
}

func Stop(stop time.Time) bool {
	if stop.Before(time.Now()) {
		sLog.Info("Exhausted time \n")
		return true
	}
	return false
}

// StopTs, periodSeconds, afterSeconds, estare procesando hasta (2 * afterSeconds) seg antes del siguiente cierre de vela
func StopTs(stop, periodSeconds, afterSeconds int64) (stopTs time.Time, execute bool) {
	execute = true
	timeProcess := (periodSeconds - (2 * afterSeconds))
	stopTs = time.Now().Add(time.Second * time.Duration(timeProcess))

	sLog.Debug("stopTs: %d  --  timeProcess: %d ", stopTs.Unix(), timeProcess)

	timePeriodSec := (time.Now().Unix() % periodSeconds)
	if timePeriodSec > afterSeconds {
		stopTs = stopTs.Add(time.Second * time.Duration(-timePeriodSec))
		execute = false
	}

	sLog.Debug("stopTs: %d  --  timePeriodSec: %d ", stopTs.Unix(), timePeriodSec)

	if stop < stopTs.Unix() {
		stopTs = time.Unix(stop, 0)
	}

	sLog.Debug("stopTs: %d  --  stop: %d \n", stopTs.Unix(), stop)

	return
}
