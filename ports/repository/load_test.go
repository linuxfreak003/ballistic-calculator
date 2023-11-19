package repository_test

import (
	"testing"

	"github.com/linuxfreak003/ballistic-calculator/pb"
	"github.com/linuxfreak003/ballistic-calculator/ports/repository"
)

func TestLoad(t *testing.T) {
	t.Run("FromProto", func(t *testing.T) {
		var l *repository.Load
		pbLoad := &pb.Load{
			LoadId: 1,
		}
		l = l.FromProto(pbLoad)
		if l.LoadId != 1 {
			t.Fatal("bad error")
		}
	})
}
