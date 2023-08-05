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

func SleepInterval(interval_seconds, seconds_after int64) {
	sleep_seconds := (interval_seconds + seconds_after) - (time.Now().Unix() % interval_seconds)
	sLog.Info("Waiting %d seconds ... ", sleep_seconds)
	time.Sleep(time.Duration(sleep_seconds) * time.Second)
}

func StopInterval(stop_ts, interval_seconds, seconds_after int64) (stop time.Time, execute bool) {
	execute = true
	processing_seconds := (interval_seconds - (2 * seconds_after))
	stop = time.Now().Add(time.Second * time.Duration(processing_seconds))

	sLog.Debug("stop: %s  ..  processing_seconds: %d ", ForLog(stop.Unix(), 0), processing_seconds)

	seconds_elapsed := (time.Now().Unix() % interval_seconds)
	if seconds_elapsed > seconds_after {
		stop = stop.Add(time.Second * time.Duration(-interval_seconds))
		execute = false
	}

	sLog.Debug("stop: %d  ..  seconds_elapsed: %d ", stop.Unix(), seconds_elapsed)

	if stop_ts > 0 && stop_ts < stop.Unix() {
		stop = time.Unix(stop_ts, 0)
	}
	sLog.Debug("stop_ts: %d  ..  stop: %s ", stop_ts, ForLog(stop.Unix(), 0))
	return
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
