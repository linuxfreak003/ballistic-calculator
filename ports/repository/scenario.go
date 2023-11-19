package repository

import "github.com/linuxfreak003/ballistic-calculator/pb"

type Scenario struct {
	ScenarioId    int64  `json:"scenario_id"`
	Name          string `json:"name"`
	EnvironmentId int64  `json:"environment_id"`
	RifleId       int64  `json:"rifle_id"`
	LoadId        int64  `json:"load_id"`
}

// FromProto ...
func (l *Scenario) FromProto(proto *pb.Scenario) *Scenario {
	if l == nil {
		l = new(Scenario)
	}

	l.ScenarioId = proto.GetScenarioId()
	l.Name = proto.GetName()
	l.EnvironmentId = proto.GetEnvironmentId()
	l.RifleId = proto.GetRifleId()
	l.LoadId = proto.GetLoadId()
	return l
}

// ToProto ...
func (l *Scenario) ToProto() *pb.Scenario {
	if l == nil {
		return nil
	}
	return &pb.Scenario{
		ScenarioId:    l.ScenarioId,
		Name:          l.Name,
		EnvironmentId: l.EnvironmentId,
		RifleId:       l.RifleId,
		LoadId:        l.LoadId,
	}
}
