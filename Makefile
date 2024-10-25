# if not APP_ENV variable present make it e0.
APP_ENV?=e0

# include and export all .env variables.
include .env
.EXPORT_ALL_VARIABLES:

speak:
	@echo $(APP_ENV)
	@echo $(APP_NAME)

run:
	go run main.go