package sDate

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yasseldg/simplego/sCandle"
	"github.com/yasseldg/simplego/sEnv"
	"github.com/yasseldg/simplego/sLog"
)

var DateSeparator = "."

func init() {
	DateSeparator = sEnv.Get("DateSeparator", DateSeparator)
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
	df := fmt.Sprintf("2006%s01%s02", DateSeparator, DateSeparator)
	switch prec {
	case 1:
		return t.Format(df)
	case 2:
		return t.Format(fmt.Sprintf("%s 15:04", df))
	case 4:
		return t.Format(fmt.Sprintf("%s 15:04:05.000", df))
	case 5:
		return t.Format(fmt.Sprintf("%s 15:04:05.000000", df))
	default:
		return t.Format(fmt.Sprintf("%s 15:04:05", df))
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
	switch v := value.(type) {
	case int:
		// Convertir timestamp de tipo int a time.Time
		return fromInt64(int64(v)), nil
	case int64:
		// Convertir timestamp de tipo int64 a time.Time
		return fromInt64(v), nil
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

func fromInt64(value int64) time.Time {
	if value < 1e12 {
		// Si el valor es menor a 1e12, se trata de un timestamp en segundos
		return time.Unix(value, 0)
	} else if value < 1e15 {
		// Si el valor es menor a 1e15, se trata de un timestamp en milisegundos
		return time.Unix(value/1e3, (value%1e3)*1e6)
	} else {
		// Si el valor es mayor o igual a 1e15, se trata de un timestamp en microsegundos
		return time.Unix(value/1e6, (value%1e6)*1e3)
	}
}
