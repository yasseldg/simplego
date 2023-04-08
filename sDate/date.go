package sDate

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yasseldg/simplego/sCandle"
	"github.com/yasseldg/simplego/sLog"
)

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

func NextStop(ts int64, interval string) int64 {
	intSec := sCandle.GetIntervalSeconds(interval)
	diff := ts % intSec
	return ts - diff + intSec - intSec/10
}

func Stop(stop time.Time) bool {
	if stop.Before(time.Now()) {
		sLog.Info("Exhausted time \n")
		return true
	}
	return false
}

func FormatD(t time.Time, prec int64) string {
	switch prec {
	case 0:
		return t.Format("2006-01-02")
	case 1:
		return t.Format("2006-01-02 15:04")
	default:
		return t.Format("2006-01-02 15:04:05")
	}
}

func ForLog(value any, prec int64) string {
	t, err := ToTime(value)
	if err != nil {
		return fmt.Sprintf("%v", value)
	}
	return fmt.Sprintf("%d ( %s )", t.Unix(), FormatD(t, prec))
}

func ForWeb(value any, prec int64) string {
	t, err := ToTime(value)
	if err != nil {
		return fmt.Sprintf("%v", value)
	}
	return fmt.Sprintf("%d <br> %s", t.Unix(), FormatD(t, prec))
}

func ToTime(value any) (time.Time, error) {

	sLog.Debug("value: %v ( %T )", value, value)
	switch v := value.(type) {
	case int:
		// Convertir timestamp de tipo int o int64 a time.Time
		return time.Unix(int64(v), 0), nil
	case int64:
		// Convertir timestamp de tipo int o int64 a time.Time
		return time.Unix(v, 0), nil
	case time.Time:
		// Devolver valor time.Time directamente
		return v, nil
	case string:
		// Convertir cadena a time.Time
		// Comprobar si la cadena representa un timestamp
		if t, err := strconv.ParseInt(v, 10, 64); err == nil {
			return time.Unix(t, 0), nil
		}
		// Comprobar si la cadena representa una fecha y hora con segundos
		if t, err := time.Parse("2006-01-02 15:04:05", v); err == nil {
			return t, nil
		}
		// Comprobar si la cadena representa una fecha y hora sin segundos
		if t, err := time.Parse("2006-01-02 15:04", v); err == nil {
			return t, nil
		}
		// Comprobar si la cadena representa una fecha
		if t, err := time.Parse("2006-01-02", v); err == nil {
			return t, nil
		}
		// Si no es un timestamp ni una fecha válida, devolver error
		return time.Time{}, fmt.Errorf("cadena %v no es un timestamp ni una fecha válida", v)
	default:
		// Si el valor no es int, int64, time.Time o string, devolver error
		return time.Time{}, fmt.Errorf("valor %v no es de tipo int, int64, time.Time o string", v)
	}
}
