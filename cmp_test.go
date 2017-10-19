package cmp

import (
	"testing"
	"github.com/franela/goblin"
)

func TestCmp(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("cmp int", func() {
		g.It("should test int comparison", func() {
			g.Assert(Int(0, 0)).Equal(0)
			g.Assert(Int(0, 1)).Equal(-1)
			g.Assert(Int(1, 0)).Equal(1)
			g.Assert(Int(1, 3)).Equal(-1)
			g.Assert(Int(3, 3)).Equal(0)
		})
		g.It("should test float comparison", func() {
			g.Assert(F64(0., 0.)).Equal(0)
			g.Assert(F64(0., 1.)).Equal(-1)
			g.Assert(F64(1., 0.)).Equal(1)
			g.Assert(F64(1., 3.)).Equal(-1)
			g.Assert(F64(3., 3.)).Equal(0)
			g.Assert(F64(0.3, 0.1+0.2)).Equal(0)
		})
	})
}
