--[[ This plugin is made for me in order to lean how Neovim plugins works. Don't use this for anything else than seeing how you can parse an external commands outpout
/Kim ]]

local M = {}
local config = require('editorconfig-lua.config')

local function setup_command()
	vim.cmd(table.concat({
		'command!',
		'LoadEditorConfig',
		'lua require("editorconfig-lua").load()'
	}, ' '))
end

local function print_debug(msg)
    local debug = vim.g.editorconfig_lua_debug
    if debug then
        print("DEBUG: " .. msg)
    end
end

M.setup = function(options)
	setmetatable(M, {
		__newindex = config.set,
		__index = config.get,
	})

	if options ~= nil then
		for k, v1 in pairs(options) do
			config.defaults[k] = v1
		end
	end

	config.verify()

	if M.enable then
		vim.g.editorconfig_lua_enabled = true
		vim.g.editorconfig_lua_cmd = M.cmd
		vim.g.editorconfig_lua_debug = M.debug
	    setup_command()
	end

end

M.load = function()
    local enabled = vim.g.editorconfig_lua_enabled
    if enabled == false then
        print_debug("editorconfig-lua is disabled")
        return
    end

    local loop = vim.loop
	local filename = vim.fn.expand('%:p')
 	local cmd = vim.g.editorconfig_lua_cmd

	-- Create a local structure to keep state
 	local s = {}

	-- Create pipes for receiving data
	s.stdout = loop.new_pipe(false)
	s.stderr = loop.new_pipe(false)

	local function onread(err, data)
		if err then
			error(err)
		end

		if data then
			local lines = vim.split(data, "\n")
			for _, d in pairs(lines) do

				if d ~= "" then
					-- Split on key=value using pattern matching
					for k, v in string.gmatch(d, "([%w%p]+)=([%w%p]+)") do
						if k == "indent_style" and v == "tab" then
							-- We need to schedule since we cannot modify buffer directly
							vim.schedule(function()
								print_debug("setl noexpand")
								vim.api.nvim_buf_set_option(0, 'expandtab', true)
							end)

						elseif k == "indent_style" and v == "space" then
							vim.schedule(function()
								print_debug("setl expand")
								vim.api.nvim_buf_set_option(0, 'expandtab', false)
							end)

						elseif k == "indent_size" then
							vim.schedule(function()
								print_debug("setl tabstop=" .. v)
								print_debug("setl shiftwidth=" .. v)
								local vsize = tonumber(v)
								vim.api.nvim_buf_set_option(0, 'tabstop', vsize)
								vim.api.nvim_buf_set_option(0, 'shiftwidth', vsize)
							end)
						end
					end
				end
			end
		end
	end

	s.handle = loop.spawn(
		cmd, {
			args = { filename },
			stdio = { nil, s.stdout, s.stderr },
		},
		function(code, signal)
			-- Not used but we can pickup return code and signal
			s.code = code
			s.signal = signal

			-- Make sure we close nicely
			if s.stdout then s.stdout:read_stop() end
			if s.stderr then s.stderr:read_stop() end

			if s.stdout and not s.stdout:is_closing() then s.stdout:close() end
			if s.stderr and not s.stderr:is_closing() then s.stderr:close() end

			s.handle:close()
		end)

	if not s.handle then
		local msg = "Unable to run command: " .. cmd .. " " .. filename
		error(msg)
	end
	loop.read_start(s.stdout, onread)
	loop.read_start(s.stderr, onread)
end

return M
