lua_note

##### # lib include
相對路徑以執行位置當路徑

for nginx
base on nginx執行的位置
=> set lua_package_path  with lua_web_root_path in nginx.conf
and all lib path related with the location


### # syntax

property exist

`var[property]`

`ngx.exit(http_status)`


foreach key
```
	for k,v in pairs(obj) do
	end
```

for  	// init count =1 to 3  , add 1 for each time
```

	for counter = 1 ,3 ,1 do
	end

	for i = 1, #myArray do
	end
```




##### # object oriented
local oo = {}

local var = setmetatable({}, [table])
	it assing [table](value and method) to a empty table {}
	=> simular instance

function oo.none_pass_self_method
end

function oo:pass_self_method
end


```
	local class = { v1 = 1 }
	oo.__index = oo 			-- simular constructor , every object could be class if you assign __index

	function oo.new()
		local self = setmetatable( {} , oo )
		return self
	end

	-- oo.m1 oo:m2 will interfere each other  , should notice?????
	function oo:m1(v)				-- allow self , call by obj.m1( obj , v ) or obj:m1( v )
		self.p1 = v or self.p1 or 0
		return self.p1
	end

	function oo.m2()				-- not allow self , call by obj.m2() or obj:m2()
		return 3
	end

	local oo_ins = oo.new() or setmetatable( {} , oo )  -- they are the same
	oo_ins.v2 =2
	oo_ins:m1()
	oo_ins.m2()

```






### # profiler

##### # installation
sudo apt-get install luarocks
luarocks install luaprofiler

##### # usage
add on top of file
```
	local profiler = require "profiler"
	profiler.start("log_filename")
```
and on the end of file
`profiler.stop()`

output columns:
	`stack_level     file_defined    function_name   line_defined    current_line    local_time      total_time`


##### # summary profiler log
	a lua script (summary.lua) of luaprofiler package under /usr/local/bin/summary.lua

exec:
	`summary.lua -v logfile`

output columns:
	` Node name           Calls       Average per call        Total time              %Time`





### # debug lib

##### # debug.getinfo([thread],f,[,what])
f = stack level
	level 0 is the current function (getinfo itself); level 1 is the function that called getinfo (except for tail calls, which do not count on the stack); and so on. If f is a number larger than the number of active functions, then getinfo returns nil.

what:
	- If present, the option 'f' adds a field named func with the function itself.
	- If present, the option 'L' adds a field named activelines with the table of valid lines.

	debug.getinfo(1,"n").name returns a table with a name for the current function
	debug.getinfo(print) returns a table with all available information about the print function.
	??? what is for S

```
	function d1() 		0: getinfo , 1: d1 , 2 : nil 	, ......
	function d2()		0: getinfo , 1: d2 , 2 : d1 		, ......
	function d3()		0: getinfo , 1: d3 , 2 : d2 		, ......
```

debug.getinfo( param1 , param2 )
	param1
		most of time , it's called by debug.getinfo(    2     ,"xxxxx")  when using sethook to detect profiler information
		=> because 0 is debug.getinfo self  ,  1 is hook function  , 2 is function that trigger hook

	param2
		none
			source,what,func,nups,short_src,name,currentline,namewhat,linedefined,lastlinedefined
		S
			source,what,lastlinedefined,linedefined,short_src
		n
			name,namewhat
		f
			func
		L
			activelines
		u
			for nup
		l
			currentline

debug.getinfo(1,"nS")	return
```
	lastlinedefined 38				-- -1 for native function , end line of function
	linedefined     31 					-- -1 for native function , start line of function
	source  @lua_oo.lua 				-- =[C] for native function
	what    Lua
	name    d1							-- 0 for  getinfo  , 1 for function1 which call debug.getinfo , 2 for function2 which call function1
	namewhat        global				-- global or field
	short_src       lua_oo.lua  			-- [C] for native file
```

##### # debug.sethook([thread],hook,mask[,count])
	- hook
		called when matching mask described.
		function hook( hooktype , line_number)
			- hooktype = "call" , "return" , "line" and "count"
	- mask
		'c': the hook is called every time Lua calls a function;
		'r': the hook is called every time Lua returns from a function;
		'l': the hook is called every time Lua enters a new line of code.
		With a count different from zero, the hook is called after every count instructions.



##### # debug.getlocal([thread,] f , local )
	f = stack level as


##### # debug.getupvalue(func , i)
	func: stack_level
	i : index of upvalue







config with lua[http://wulijun.github.io/2012/09/04/nginx-lua.html]
luajit
	Not Yet Implemented [http://wiki.luajit.org/NYI]

	luajit -bg xxx.lua xxx.luac

lua
	lua5.1 manual [http://www.lua.org/manual/5.1/manual.html]
	lua5.2 manual [http://www.lua.org/manual/5.2/manual.html]
	lua unofficial FAQ [http://www.luafaq.org/]

	lua io [http://www.tutorialspoint.com/lua/lua_file_io.htm]
	lua simple json [http://regex.info/blog/lua/json]



	lua userdata and c
http://www.jellythink.com/archives/587

luajit not implemented yet
http://wiki.luajit.org/NYI




