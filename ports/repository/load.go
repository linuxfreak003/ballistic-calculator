package repository

import "github.com/linuxfreak003/ballistic-calculator/pb"

// Load ...
type Load struct {
	LoadId         int64
	Bullet         Bullet
	MuzzleVelocity float32
}

// Bullet ...
type Bullet struct {
	Caliber float32
	Weight  float32
	BC      Coefficient
	Length  float32
}

// Coefficient ...
type Coefficient struct {
	Value    float32
	DragFunc int
}

// FromProto ...
func (l *Load) FromProto(proto *pb.Load) {
	if l == nil {
		(*l) = Load{}
	}

	l.LoadId = proto.GetLoadId()
	l.Bullet = Bullet{
		Caliber: proto.GetBullet().GetCaliber(),
		Weight:  proto.GetBullet().GetWeight(),
		BC: Coefficient{
			Value:    proto.GetBullet().GetBc().GetValue(),
			DragFunc: int(proto.GetBullet().GetBc().GetFunction()),
		},
		Length: proto.GetBullet().GetLength(),
	}
	l.MuzzleVelocity = proto.GetMuzzleVelocity()
}

// ToProto ...
func (l *Load) ToProto() *pb.Load {
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
