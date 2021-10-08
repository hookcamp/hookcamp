# init sets up git to recognise the .githooks directory as the hooks path for this repo
# it also makes all scripts in the .githooks folder executable
init:
	git config core.hooksPath .githooks
	chmod +x /*