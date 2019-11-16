local PI = 3.141592653589793
local SOLAR_MASS = 4 * PI * PI
local DAYS_PER_YEAR = 365.24

local Jupiter = {
        x = 4.84143144246472090e+00,
        y = -1.16032004402742839e+00,
        z = -1.03622044471123109e-01,
        vx = 1.66007664274403694e-03 * DAYS_PER_YEAR,
        vy = 7.69901118419740425e-03 * DAYS_PER_YEAR,
        vz = -6.90460016972063023e-05 * DAYS_PER_YEAR,
        mass = 9.54791938424326609e-04 * SOLAR_MASS
}

local Saturn = {
        x = 8.34336671824457987e+00,
        y = 4.12479856412430479e+00,
        z = -4.03523417114321381e-01,
        vx = -2.76742510726862411e-03 * DAYS_PER_YEAR,
        vy = 4.99852801234917238e-03 * DAYS_PER_YEAR,
        vz = 2.30417297573763929e-05 * DAYS_PER_YEAR,
        mass = 2.85885980666130812e-04 * SOLAR_MASS
}

local Uranus = {
        x = 1.28943695621391310e+01,
        y = -1.51111514016986312e+01,
        z = -2.23307578892655734e-01,
        vx = 2.96460137564761618e-03 * DAYS_PER_YEAR,
        vy = 2.37847173959480950e-03 * DAYS_PER_YEAR,
        vz = -2.96589568540237556e-05 * DAYS_PER_YEAR,
        mass = 4.36624404335156298e-05 * SOLAR_MASS
}

local Neptune = {
        x = 1.53796971148509165e+01,
        y = -2.59193146099879641e+01,
        z = 1.79258772950371181e-01,
        vx = 2.68067772490389322e-03 * DAYS_PER_YEAR,
        vy = 1.62824170038242295e-03 * DAYS_PER_YEAR,
        vz = -9.51592254519715870e-05 * DAYS_PER_YEAR,
        mass = 5.15138902046611451e-05 * SOLAR_MASS
}

local Sun = {
        x = 0,
        y = 0,
        z = 0,
        vx = 0,
        vy = 0,
        vz = 0,
        mass = SOLAR_MASS
}

local function offsetMomentum(b, px, py, pz)
        b.vx = -px / SOLAR_MASS
        b.vy = -py / SOLAR_MASS
        b.vz = -pz / SOLAR_MASS
end

-- NBodySystems represents the system of n-bodies
NBodySystems = {bodies = {}}

function NBodySystems:new(init)
        init = init or {}
        setmetatable(init, self)
        self.__index = self
        self.bodies = {Sun, Jupiter, Saturn, Uranus, Neptune}

        local px, py, pz = 0, 0, 0

        for i = 1, #self.bodies, 1 do
                px = px + self.bodies[i].vx * self.bodies[i].mass
                py = py + self.bodies[i].vy * self.bodies[i].mass
                pz = pz + self.bodies[i].vz * self.bodies[i].mass
        end

        offsetMomentum(self.bodies[1], px, py, pz)
        return init
end

function NBodySystems:advance(dt)
        for i = 1, #self.bodies, 1 do
                local iBody = self.bodies[i]
                for j = i + 1, #self.bodies, 1 do
                        local dx = iBody.x - self.bodies[j].x
                        local dy = iBody.y - self.bodies[j].y
                        local dz = iBody.z - self.bodies[j].z

                        local dSquared = dx * dx + dy * dy + dz * dz
                        local distance = dSquared ^ 0.5
                        local mag = dt / (dSquared * distance)

                        iBody.vx = iBody.vx - dx * self.bodies[j].mass * mag
                        iBody.vy = iBody.vy - dy * self.bodies[j].mass * mag
                        iBody.vz = iBody.vz - dz * self.bodies[j].mass * mag

                        self.bodies[j].vx = self.bodies[j].vx + dx * iBody.mass * mag
                        self.bodies[j].vy = self.bodies[j].vy + dy * iBody.mass * mag
                        self.bodies[j].vz = self.bodies[j].vz + dz * iBody.mass * mag
                end
        end

        for i = 1, #self.bodies, 1 do
                local body = self.bodies[i]
                body.x = body.x + dt * body.vx
                body.y = body.y + dt * body.vy
                body.z = body.z + dt * body.vz
        end
end

function NBodySystems:energy()
        local dx, dy, dz, distance
        local e = 0

        for i = 1, #self.bodies, 1 do
                local iBody = self.bodies[i]
                local val = 0.5 * iBody.mass * (iBody.vx * iBody.vx + iBody.vy * iBody.vy + iBody.vz * iBody.vz)
                e = e + val

                for j = i + 1, #self.bodies, 1 do
                        local jBody = self.bodies[j]
                        dx = iBody.x - jBody.x
                        dy = iBody.y - jBody.y
                        dz = iBody.z - jBody.z

                        distance = (dx * dx + dy * dy + dz * dz) ^ 0.5
                        val = (iBody.mass * jBody.mass) / distance
                        e = e - val
                end
        end
        return e
end

function NBodySystems:run(n)
        if n < 0 then
                return 0
        end
        for i = 1, n, 1 do
                self:advance(0.01)
        end
        return self:energy()
end

function Test(input)
        local nbody = NBodySystems:new(nil)
        local finalEnergy = nbody:run(input)

        -- print(finalEnergy)
        -- Output: -0.16907302171469984, Input: 10

        return finalEnergy
end
