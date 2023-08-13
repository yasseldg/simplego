package sTrading

import (
	"github.com/yasseldg/simplego/sFile"
	"github.com/yasseldg/simplego/sJson"
	"github.com/yasseldg/simplego/sLog"
)

// Backtest Positions

type BacktestPositions struct {
	Positions Positions `bson:"positions" json:"positions"`
}

type Position struct {
	Side       int     `bson:"type" json:"type"` // 0 - Long / 1 - Short
	EntryPrice float64 `bson:"entryPrice" json:"entryPrice"`
	EntryTs    int64   `bson:"entryTS" json:"entryTS"`
	ExitTs     int64   `bson:"exitTS" json:"exitTS"`
	TakeProfit float64 `bson:"tp" json:"tp"`
	StopLoss   float64 `bson:"sl" json:"sl"`
}
type Positions []*Position

// The file_path is the path to the file where the positions will be exported.
// The file will be deleted if it exists.
func (bp BacktestPositions) Export(file_path string) error {
	err := sFile.DeletePath(file_path)
	if err != nil {
		sLog.Error("sFile.Delete(): %s", err)
		return err
	}

	err = sJson.Export(file_path, bp)
	if err != nil {
		sLog.Error("sJson.Export(): %s", err)
		return err
	}
	return nil
}

func (s Side) PositionSide() int {
	switch s {
	case Side_Buy:
		return 0
	case Side_Sell:
		return 1
	default:
		return -1
	}
}
