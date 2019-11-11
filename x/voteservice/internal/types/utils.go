package types

import (
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto"
)

func GetPointFromSPoint(sp SPoint, base int) crypto.Point {
	return crypto.Point{
		X: common.GetBigInt(sp.X, base),
		Y: common.GetBigInt(sp.Y, base),
	}
}
func GetSPointFromPoint(p crypto.Point) SPoint {
	return SPoint{
		X: p.X.String(),
		Y: p.Y.String(),
	}
}
