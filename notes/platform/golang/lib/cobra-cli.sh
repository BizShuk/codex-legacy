#!/bin/bash

MODULE_NAME=""
### Installation of cobra-cli
go install github.com/spf13/cobra-cli@latest


### Init cobra project
go mod init ${MODULE_NAME}
cobra-cli init ${MODULE_NAME} --auther Shuk

### Add Command
COMMAND_NAME=""
cobra-cli add ${COMMAND_NAME} [-p parent cmd: rootCmd] [--viper]



### Access flags
#     Flags().LookUp(key) => bur how to transfer flag to value type
#     embeded with variable

### PersistenFlags
#     flags for current and sub command against local flags
