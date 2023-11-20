package repository

import (
	"github.com/linuxfreak003/ballistic-calculator/pb"
	"gitlab.com/linuxfreak003/ballistic"
)

// Environment ...
type Environment struct {
	EnvironmentId      int64   `json:"environment_id"`
	Temperature        float64 `json:"temperature"`          // In Fahrenheit
	Altitude           int64   `json:"altitude"`             // in ft
	Pressure           float64 `json:"pressure"`             // in Hg
	Humidity           float64 `json:"humidity"`             // Himidity percentage
	WindAngle          float64 `json:"wind_angle"`           // Wind direction in degrees (0 to 360)
	WindSpeed          float64 `json:"wind_speed"`           // Wind speed in mph
	PressureIsAbsolute bool    `json:"pressure_is_absolute"` // If true, only Pressure will be used
	Latitude           float64 `json:"latitude"`             // Latitude (deg)
	Azimuth            float64 `json:"azimuth"`              // Azimuth of fire (deg)
}

// FromProto ...
func (l *Environment) FromProto(proto *pb.Environment) *Environment {
	if l == nil {
		l = new(Environment)
	}

	l.EnvironmentId = proto.GetEnvironmentId()
	l.Temperature = proto.GetTemperature()
	l.Altitude = proto.GetAltitude()
	l.Pressure = proto.GetPressure()
	l.Humidity = proto.GetHumidity()
	l.WindAngle = proto.GetWindAngle()
	l.WindSpeed = proto.GetWindSpeed()
	l.PressureIsAbsolute = proto.GetPressureIsAbsolute()
	l.Latitude = proto.GetLatitude()
	l.Azimuth = proto.GetAzimuth()
	return l
}

// ToProto ...
func (l *Environment) ToProto() *pb.Environment {
	if l == nil {
		return nil
	}
	return &pb.Environment{
		EnvironmentId:      l.EnvironmentId,
		Temperature:        l.Temperature,
		Altitude:           l.Altitude,
		Pressure:           l.Pressure,
		Humidity:           l.Humidity,
		WindAngle:          l.WindAngle,
		WindSpeed:          l.WindSpeed,
		PressureIsAbsolute: l.PressureIsAbsolute,
		Latitude:           l.Latitude,
		Azimuth:            l.Azimuth,
	}
}

// ToBallistic ...
func (l *Environment) ToBallistic() *ballistic.Environment {
	if l == nil {
		return nil
	}
	return &ballistic.Environment{
		Temperature:        l.Temperature,
		Altitude:           int(l.Altitude),
		Pressure:           l.Pressure,
		Humidity:           l.Humidity,
		WindAngle:          l.WindAngle,
		WindSpeed:          l.WindSpeed,
		PressureIsAbsolute: l.PressureIsAbsolute,
		Latitude:           l.Latitude,
		Azimuth:            l.Azimuth,
	}
}
