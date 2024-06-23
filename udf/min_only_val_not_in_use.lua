-- min reducer
local function get_min(val1, val2)
    local min = nil
    if val1 then
        if val2 then
            if val1 < val2 then min = val1 else min = val2 end
        else min = val1
        end
    else
        if val2 then min = val2 end
    end
    return min
end

-- mapper for various single bin aggregates
local function rec_to_bin_value_closure(bin)
    local function rec_to_bin_value(rec)
    -- if a numeric bin exists in record return its value; otherwise return nil
        local val = rec[bin]
        if (type(val) ~= "number") then val = nil end
        return val
    end
    return rec_to_bin_value
end

-- min
function min(stream, bin)
    return stream : map(rec_to_bin_value_closure(bin)) : reduce(get_min)
end