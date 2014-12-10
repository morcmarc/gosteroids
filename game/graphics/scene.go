package graphics

import (
	o "github.com/morcmarc/gosteroids/game/objects"
	"github.com/satori/go.uuid"
)

type Scene struct {
	ObjectManager *o.ObjectManager
	Background    *Background
	Spaceship     *Spaceship
	Asteroids     map[uuid.UUID]*Asteroid
	Projectiles   map[uuid.UUID]*Projectile
	Score         *Score
	gameOver      bool
}

type SceneObject interface {
	Draw(ct float32)
	Delete()
}

func NewScene(om *o.ObjectManager, w, h, bgQuality int) *Scene {
	s := &Scene{
		ObjectManager: om,
		Spaceship:     NewSpaceship(om.Spaceship),
		Background:    NewBackground(w, h, bgQuality),
		Score:         NewScore(w, h),
		Asteroids:     map[uuid.UUID]*Asteroid{},
		Projectiles:   map[uuid.UUID]*Projectile{},
		gameOver:      false,
	}

	s.Reset()

	return s
}

func (s *Scene) Fire() {
	po := s.ObjectManager.FireProjectile()
	p := NewProjectile(po)
	s.AddProjectile(p)
}

func (s *Scene) AddAsteroid(a *Asteroid) {
	s.Asteroids[a.AObject.Id] = a
}

func (s *Scene) AddProjectile(p *Projectile) {
	s.Projectiles[p.PSObject.Id] = p
}

func (s *Scene) RemoveAsteroid(id uuid.UUID) {
	// s.Asteroids[id].Delete()
	s.ObjectManager.RemoveAsteroid(id)
	delete(s.Asteroids, id)
}

func (s *Scene) RemoveProjectile(id uuid.UUID) {
	// s.Projectiles[id].Delete()
	s.ObjectManager.RemoveProjectile(id)
	delete(s.Projectiles, id)
}

func (s *Scene) Update(ct float32) {
	s.ObjectManager.Update()

	if !s.gameOver {
		s.Score.Points += 1
	}

	for _, p := range s.Projectiles {
		if p.PSObject.IsOffScreen() {
			s.RemoveProjectile(p.PSObject.Id)
		}
	}

	hitP, hitA := s.CheckHits()
	if hitP != [16]byte{} && hitA != [16]byte{} {
		pointVal := int(s.Asteroids[hitA].AObject.Radius * float64(1000))
		s.Score.Points += pointVal
		s.RemoveAsteroid(hitA)
		s.RemoveProjectile(hitP)
	}
}

func (s *Scene) GameOver() {
	s.gameOver = true
}

func (s *Scene) CheckCollision() bool {
	return s.ObjectManager.CheckCollision()
}

func (s *Scene) CheckHits() (uuid.UUID, uuid.UUID) {
	return s.ObjectManager.CheckHits()
}

func (s *Scene) Reset() {
	s.gameOver = false
	s.Score.Points = 0
	s.ObjectManager.Reset()

	for _, a := range s.Asteroids {
		s.RemoveAsteroid(a.AObject.Id)
	}

	for _, ao := range s.ObjectManager.Asteroids {
		a := NewAsteroid(ao)
		s.AddAsteroid(a)
	}
}

func (s *Scene) Draw(ct float32) {
	s.Background.Draw(ct)
	s.Spaceship.Draw(ct)
	for _, a := range s.Asteroids {
		a.Draw(ct)
	}
	for _, p := range s.Projectiles {
		p.Draw(ct)
	}
	s.Score.Draw(ct)
}

func (s *Scene) Delete() {
	s.Background.Delete()
	s.Spaceship.Delete()
	for _, a := range s.Asteroids {
		a.Delete()
	}
	for _, p := range s.Projectiles {
		p.Delete()
	}
	s.Score.Delete()
}
