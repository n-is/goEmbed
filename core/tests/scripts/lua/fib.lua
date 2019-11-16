local function fib(n)
        if n < 0 then
                return -1
        end

        if n == 0 then
                return 0
        elseif n == 1 then
                return 1
        end

        local n0, n1 = 0, 1

        for i = 2, n, 1 do
                local tmp = n0 + n1
                n0 = n1
                n1 = tmp
        end

        return n1
end

function Test(n)
        return fib(n)
end
