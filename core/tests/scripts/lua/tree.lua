local function BottomUpTree(depth)
        if depth > 0 then
                depth = depth - 1
                local left, right = BottomUpTree(depth), BottomUpTree(depth)
                return {left, right}
        else
                return {}
        end
end

local function ItemCheck(tree)
        if tree[1] then
                return 1 + ItemCheck(tree[1]) + ItemCheck(tree[2])
        else
                return 1
        end
end

local function tree(N)
        if (N < 0) then
                return 0
        end

        local mindepth = 4
        local maxdepth = mindepth + 2
        if maxdepth < N then
                maxdepth = N
        end

        local stretchdepth = maxdepth + 1
        local stretchtree = BottomUpTree(stretchdepth)
        ItemCheck(stretchtree)

        local longlivedtree = BottomUpTree(maxdepth)
        local mmd = maxdepth + mindepth

        for depth = mindepth, maxdepth, 2 do
                local iterations = 2 ^ (mmd - depth)
                local check = 0
                for i = 1, iterations do
                        check = check + ItemCheck(BottomUpTree(depth))
                end
        end
        ItemCheck(longlivedtree)

        return 0
end

function Test(N)
        return tree(N)
end
