class NBodySystem(object):

        class Body(object):
                PI            = 3.141592653589793
                SOLAR_MASS    = 4 * PI * PI
                DAYS_PER_YEAR = 365.24
                def __init__(self,x,y,z,vx,vy,vz,mass):
                        self.x = x
                        self.y = y
                        self.z = z
                        self.vx = vx
                        self.vy = vy
                        self.vz = vz
                        self.mass = mass

                @classmethod
                def Jupiter(cls):
                        return cls(
                                x =    4.84143144246472090e+00,
                                y =    -1.16032004402742839e+00,
                                z =    -1.03622044471123109e-01,
                                vx =   1.66007664274403694e-03 * cls.DAYS_PER_YEAR,
                                vy =   7.69901118419740425e-03 * cls.DAYS_PER_YEAR,
                                vz =   -6.90460016972063023e-05 * cls.DAYS_PER_YEAR,
                                mass = 9.54791938424326609e-04 * cls.SOLAR_MASS
                        )

                @classmethod
                def Saturn(cls):
                        return cls(
                                x =    8.34336671824457987e+00,
                                y =    4.12479856412430479e+00,
                                z =    -4.03523417114321381e-01,
                                vx =   -2.76742510726862411e-03 * cls.DAYS_PER_YEAR,
                                vy =   4.99852801234917238e-03 * cls.DAYS_PER_YEAR,
                                vz =   2.30417297573763929e-05 * cls.DAYS_PER_YEAR,
                                mass = 2.85885980666130812e-04 * cls.SOLAR_MASS,
                        )

                @classmethod
                def Uranus(cls):
                        return cls(
                                x =    1.28943695621391310e+01,
                                y =    -1.51111514016986312e+01,
                                z =    -2.23307578892655734e-01,
                                vx =   2.96460137564761618e-03 * cls.DAYS_PER_YEAR,
                                vy =   2.37847173959480950e-03 * cls.DAYS_PER_YEAR,
                                vz =   -2.96589568540237556e-05 * cls.DAYS_PER_YEAR,
                                mass = 4.36624404335156298e-05 * cls.SOLAR_MASS,
                        )

                @classmethod
                def Neptune(cls):
                        return cls(
                                x =    1.53796971148509165e+01,
                                y =    -2.59193146099879641e+01,
                                z =    1.79258772950371181e-01,
                                vx =   2.68067772490389322e-03 * cls.DAYS_PER_YEAR,
                                vy =   1.62824170038242295e-03 * cls.DAYS_PER_YEAR,
                                vz =   -9.51592254519715870e-05 * cls.DAYS_PER_YEAR,
                                mass = 5.15138902046611451e-05 * cls.SOLAR_MASS,
                        )

                @classmethod
                def Sun(cls):
                        return cls(
                                x =    0,
                                y =    0,
                                z =    0,
                                vx =   0,
                                vy =   0,
                                vz =   0,
                                mass = cls.SOLAR_MASS,
                        )
                
                
                def offsetMomentum(self, px, py, pz):
                        self.vx = -px / self.SOLAR_MASS
                        self.vy = -py / self.SOLAR_MASS
                        self.vz = -pz / self.SOLAR_MASS


        def __init__(self):
                self.bodies = [ self.Body.Sun(),
                                self.Body.Jupiter(),
                                self.Body.Saturn(),
                                self.Body.Uranus(),
                                self.Body.Neptune()]
                
                px, py, pz = 0.0, 0.0, 0.0

                for i in range(len(self.bodies)):
                        px += self.bodies[i].vx * self.bodies[i].mass
                        py += self.bodies[i].vy * self.bodies[i].mass
                        pz += self.bodies[i].vz * self.bodies[i].mass

                self.bodies[0].offsetMomentum(px, py, pz)
                        
        def advance(self, dt):
                for i in range(len(self.bodies)):
                        iBody = self.bodies[i]
                        for j in range(i + 1, len(self.bodies)):
                                dx = iBody.x - self.bodies[j].x
                                dy = iBody.y - self.bodies[j].y
                                dz = iBody.z - self.bodies[j].z

                                dSquared = dx*dx + dy*dy + dz*dz
                                distance = dSquared ** 0.5
                                mag = dt / (dSquared * distance)

                                iBody.vx -= dx * self.bodies[j].mass * mag
                                iBody.vy -= dy * self.bodies[j].mass * mag
                                iBody.vz -= dz * self.bodies[j].mass * mag

                                self.bodies[j].vx += dx * iBody.mass * mag
                                self.bodies[j].vy += dy * iBody.mass * mag
                                self.bodies[j].vz += dz * iBody.mass * mag

                for i in range(len(self.bodies)):
                        body = self.bodies[i]
                        body.x = body.x + dt * body.vx
                        body.y = body.y + dt * body.vy
                        body.z = body.z + dt * body.vz

        def energy(self):
                e = 0

                for i in range(len(self.bodies)):
                        iBody = self.bodies[i]
                        val = 0.5 * iBody.mass * (iBody.vx*iBody.vx + iBody.vy*iBody.vy + iBody.vz*iBody.vz)
                        e += val

                        for j in range(i + 1, len(self.bodies)):
                                jBody = self.bodies[j]
                                dx = iBody.x - jBody.x
                                dy = iBody.y - jBody.y
                                dz = iBody.z - jBody.z

                                distance = (dx*dx + dy*dy + dz*dz) ** 0.5
                                val = (iBody.mass * jBody.mass) / distance
                                e -= val
                return e

        def run(self, n):
                if n < 0:
                        return 0
                for _ in range(n):
                        self.advance(0.01)
                return self.energy()
