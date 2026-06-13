local f = io.open("/home/shuk/lua_server_base/appversion","r")
local content = f:read("*all")
f:close()
--print(content)



for str in string.gmatch(content,"[%w_]+:[%w]+") do
--    print(str)

    local tmp = {}
    local tmp = string.gmatch(str,"([^:]+)")
    print(tmp(),tmp())

end
