local defaults = {
	enable = true,
	cmd = '/sbin/editorconfig3',
	debug = false,
}

local function set(_, key, value)
	defaults[key] = value
end

local function get(_, key)
	return defaults[key]
end

local function verify()
	if type(defaults.enable) ~= 'boolean' then
		defaults.enable = true
	end
	if type(defaults.cmd) ~= 'string' then
		defaults.cmd = '/sbin/editorconfig3'
	end
	if type(defaults.debug) ~= 'boolean' then
		defaults.debug = false
	end
end

return {
	defaults = defaults,
	get = get,
	set = set,
	verify = verify
}
