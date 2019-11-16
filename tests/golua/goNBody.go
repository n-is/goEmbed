package golua

import (
	"math"
)

func setupNBodyDemo() *nBodySystems {

	return newNBodySystem()
}

func runNBodyDemo(bodies *nBodySystems, n int64) float64 {

	for i := int64(0); i < n; i++ {
		bodies.advance(0.01)
	}
	return bodies.energy()
}

type nBodySystems struct {
	bodies []*body
}

func newNBodySystem() *nBodySystems {
	bodies := []*body{sun(), jupiter(), saturn(), uranus(), neptune()}

	px := float64(0)
	py := float64(0)
	pz := float64(0)

	for i := 0; i < len(bodies); i++ {
		px += bodies[i].vx * bodies[i].mass
		py += bodies[i].vy * bodies[i].mass
		pz += bodies[i].vz * bodies[i].mass
	}
	bodies[0].offsetMomentum(px, py, pz)

	return &nBodySystems{bodies: bodies}
}

func (nb *nBodySystems) advance(dt float64) {
	for i := 0; i < len(nb.bodies); i++ {
		iBody := nb.bodies[i]
		for j := i + 1; j < len(nb.bodies); j++ {
			dx := iBody.x - nb.bodies[j].x
			dy := iBody.y - nb.bodies[j].y
			dz := iBody.z - nb.bodies[j].z

			dSquared := dx*dx + dy*dy + dz*dz
			distance := math.Sqrt(dSquared)
			mag := dt / (dSquared * distance)

			iBody.vx -= dx * nb.bodies[j].mass * mag
			iBody.vy -= dy * nb.bodies[j].mass * mag
			iBody.vz -= dz * nb.bodies[j].mass * mag

			nb.bodies[j].vx += dx * iBody.mass * mag
			nb.bodies[j].vy += dy * iBody.mass * mag
			nb.bodies[j].vz += dz * iBody.mass * mag
		}
	}

	for _, body := range nb.bodies {
		body.x += dt * body.vx
		body.y += dt * body.vy
		body.z += dt * body.vz
	}
}

func (nb *nBodySystems) energy() float64 {
	var dx, dy, dz, distance float64
	e := float64(0)

	for i := 0; i < len(nb.bodies); i++ {
		iBody := nb.bodies[i]
		val := 0.5 * iBody.mass * (iBody.vx*iBody.vx + iBody.vy*iBody.vy + iBody.vz*iBody.vz)
		e += val

		for j := i + 1; j < len(nb.bodies); j++ {
			jBody := nb.bodies[j]
			dx = iBody.x - jBody.x
			dy = iBody.y - jBody.y
			dz = iBody.z - jBody.z

			distance = math.Sqrt(dx*dx + dy*dy + dz*dz)
			val := (iBody.mass * jBody.mass) / distance
			e -= val
		}
	}
	return e
}

const (
	pi            = 3.141592653589793
	solar_mass    = 4 * pi * pi
	days_per_year = 365.24
)

type body struct {
	x, y, z, vx, vy, vz, mass float64
}

func jupiter() *body {
	p := &body{
		x:    4.84143144246472090e+00,
		y:    -1.16032004402742839e+00,
		z:    -1.03622044471123109e-01,
		vx:   1.66007664274403694e-03 * days_per_year,
		vy:   7.69901118419740425e-03 * days_per_year,
		vz:   -6.90460016972063023e-05 * days_per_year,
		mass: 9.54791938424326609e-04 * solar_mass,
	}

	return p
}

func saturn() *body {
	p := &body{
		x:    8.34336671824457987e+00,
		y:    4.12479856412430479e+00,
		z:    -4.03523417114321381e-01,
		vx:   -2.76742510726862411e-03 * days_per_year,
		vy:   4.99852801234917238e-03 * days_per_year,
		vz:   2.30417297573763929e-05 * days_per_year,
		mass: 2.85885980666130812e-04 * solar_mass,
	}

	return p
}

func uranus() *body {
	p := &body{
		x:    1.28943695621391310e+01,
		y:    -1.51111514016986312e+01,
		z:    -2.23307578892655734e-01,
		vx:   2.96460137564761618e-03 * days_per_year,
		vy:   2.37847173959480950e-03 * days_per_year,
		vz:   -2.96589568540237556e-05 * days_per_year,
		mass: 4.36624404335156298e-05 * solar_mass,
	}

	return p
}

func neptune() *body {
	p := &body{
		x:    1.53796971148509165e+01,
		y:    -2.59193146099879641e+01,
		z:    1.79258772950371181e-01,
		vx:   2.68067772490389322e-03 * days_per_year,
		vy:   1.62824170038242295e-03 * days_per_year,
		vz:   -9.51592254519715870e-05 * days_per_year,
		mass: 5.15138902046611451e-05 * solar_mass,
	}

	return p
}

func sun() *body {
	p := &body{
		mass: solar_mass,
	}

	return p
}

func (b *body) offsetMomentum(px, py, pz float64) {
	b.vx = -px / solar_mass
	b.vy = -py / solar_mass
	b.vz = -pz / solar_mass
}
