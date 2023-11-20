package repository

import (
	"github.com/linuxfreak003/ballistic-calculator/pb"
	"gitlab.com/linuxfreak003/ballistic"
)

// Load includes relevant data about the load
type Load struct {
	LoadId         int64
	Bullet         *Bullet
	MuzzleVelocity float64
}

// Bullet ...
type Bullet struct {
	Caliber float64
	Weight  float64
	BC      *Coefficient
	Length  float64
}

// Coefficient ...
type Coefficient struct {
	Value    float64
	DragFunc int
}

// FromProto ...
func (l *Load) FromProto(proto *pb.Load) *Load {
	if l == nil {
		l = new(Load)
	}

	l.LoadId = proto.GetLoadId()
	l.Bullet = &Bullet{
		Caliber: proto.GetBullet().GetCaliber(),
		Weight:  proto.GetBullet().GetWeight(),
		BC: &Coefficient{
			Value:    proto.GetBullet().GetBc().GetValue(),
			DragFunc: int(proto.GetBullet().GetBc().GetFunction()),
		},
		Length: proto.GetBullet().GetLength(),
	}
	l.MuzzleVelocity = proto.GetMuzzleVelocity()
	return l
}

// ToProto ...
func (l *Load) ToProto() *pb.Load {
	if l == nil {
		return nil
	}
	return &pb.Load{
		LoadId: l.LoadId,
		Bullet: &pb.Bullet{
			Caliber: l.Bullet.Caliber,
			Weight:  l.Bullet.Weight,
			Bc: &pb.BallisticCoefficient{
				Value:    l.Bullet.BC.Value,
				Function: pb.DragFunction(l.Bullet.BC.DragFunc),
			},
			Length: l.Bullet.Length,
		},
		MuzzleVelocity: l.MuzzleVelocity,
	}
}

// ToBallistic ...
func (l *Load) ToBallistic() *ballistic.Load {
	if l == nil {
		return nil
	}
	return &ballistic.Load{
		Bullet: &ballistic.Bullet{
			Caliber: l.Bullet.Caliber,
			Weight:  l.Bullet.Weight,
			BC: &ballistic.Coefficient{
				Value:    l.Bullet.BC.Value,
				DragFunc: ballistic.DragFunction(l.Bullet.BC.DragFunc),
			},
			Length: l.Bullet.Length,
		},
		MuzzleVelocity: l.MuzzleVelocity,
	}
}
