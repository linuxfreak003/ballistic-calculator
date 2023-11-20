package repository

import (
	"github.com/linuxfreak003/ballistic-calculator/pb"
	"gitlab.com/linuxfreak003/ballistic"
)

// Rifle contains all the variables of the rifle
type Rifle struct {
	RifleId            int64
	Name               string  `json:"name"`
	SightHeight        float64 `json:"sight_height"`         // Sight height in inches (default 1.5)
	BarrelTwist        float64 `json:"barrel_twist"`         // Barrel twist rate (default 7)
	TwistDirectionLeft bool    `json:"twist_direction_left"` // If twist direction is left
	ZeroRange          float64 `json:"zero_range"`           // Zero range of rifle in yrds (default 100)
}

// FromProto ...
func (l *Rifle) FromProto(proto *pb.Rifle) *Rifle {
	if l == nil {
		l = new(Rifle)
	}

	l.RifleId = proto.GetRifleId()
	l.Name = proto.GetName()
	l.SightHeight = proto.GetSightHeight()
	l.BarrelTwist = proto.GetBarrelTwist()
	l.TwistDirectionLeft = proto.GetTwistDirectionLeft()
	l.ZeroRange = proto.GetZeroRange()
	return l
}

// ToProto ...
func (l *Rifle) ToProto() *pb.Rifle {
	if l == nil {
		return nil
	}
	return &pb.Rifle{
		RifleId:            l.RifleId,
		Name:               l.Name,
		SightHeight:        l.SightHeight,
		BarrelTwist:        l.BarrelTwist,
		TwistDirectionLeft: l.TwistDirectionLeft,
		ZeroRange:          l.ZeroRange,
	}
}

// ToBallistic ...
func (l *Rifle) ToBallistic() *ballistic.Rifle {
	if l == nil {
		return nil
	}
	return &ballistic.Rifle{
		SightHeight:        l.SightHeight,
		BarrelTwist:        l.BarrelTwist,
		TwistDirectionLeft: l.TwistDirectionLeft,
		ZeroRange:          l.ZeroRange,
	}
}
