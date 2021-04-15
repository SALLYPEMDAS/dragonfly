package item

import (
	"github.com/df-mc/dragonfly/dragonfly/block/cube"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/df-mc/dragonfly/dragonfly/world/particle"
	"github.com/go-gl/mathgl/mgl64"
)

// BoneMeal is an item used to force growth in plants & crops.
type BoneMeal struct{}

// BoneMealAffected represents a block that is affected when bone meal is used on it.
type BoneMealAffected interface {
	// BoneMeal attempts to affect the block using a bone meal item.
	BoneMeal(pos cube.Pos, w *world.World) bool
}

// UseOnBlock ...
func (b BoneMeal) UseOnBlock(pos cube.Pos, _ cube.Face, _ mgl64.Vec3, w *world.World, _ User, ctx *UseContext) bool {
	if bm, ok := w.Block(pos).(BoneMealAffected); ok && bm.BoneMeal(pos, w) {
		ctx.CountSub = 1
		w.AddParticle(pos.Vec3(), particle.Bonemeal{})
		return true
	}
	return false
}

// EncodeItem ...
func (b BoneMeal) EncodeItem() (id int32, meta int16) {
	return 351, 15
}
