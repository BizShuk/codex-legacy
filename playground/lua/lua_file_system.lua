
-- from https://keplerproject.github.io/luafilesystem/index.html
local lfs = require("lfs")

function attrdir(path)

    for file in lfs.dir(path) do
        if file ~= "." and file ~= ".." and file ~= ".git" then
            local f = path..'/'..file

            local attr = lfs.attributes (f)
            assert (type(attr) == "table")
    
            if attr.mode == "directory" then
                ngx.print ("\t "..f)                    -- show dir name
                attrdir (f)
            else
                ngx.print ("\t "..f)                    -- show file name
                for name, value in pairs(attr) do
                    ngx.print (name, value)         -- show file attribute
                end
            end
        end
    end

end

ngx.print(ngx.var.lua_dir)
attrdir(ngx.var.lua_dir)
