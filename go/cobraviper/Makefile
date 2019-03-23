
GOC=go

default: all
all: nooption option config override clean

build:
	@$(GOC) build

nooption: build
	@./cobraviper msg

option: build
	@./cobraviper msg --msg 'We are the keepers of the sacred words: Ni, Ping, and Nee-womm!'

config: build
	@./cobraviper msg --config ./config.yaml

override: build
	@./cobraviper msg --config ./config.yaml --log 'knigts.of.Ekky-ekky-ekky-ekky-zBangzoom.Boingznourrrwringnmmm' --msg 'We are now *no longer* the Knights Who Say \"Ni\"!'

clean:
	@rm ./cobraviper
