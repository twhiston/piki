# Piki

On server go app written of Piki Raspberry distribution to make managing your Pi Kiosk easier
[PiKiOs](https://github.com/twhiston/PiKiOS)

## Commands

Some commands currently incomplete and more coming soon

### boot

Switch the pi system boot between the app and recovery mode

### dash

control the dashboard with commands in this namespace

#### api
Root of all subcommands that interact with the piki dashboard backend app

##### settings

control server settings

###### set

set a setting

###### get

get a setting value, or all settings

### hostname

Set the hostname of this device (analogous to raspi-config set hostname)

### logs

Look at some app logs easily

#### httpd

Look at the server logs

### net

Add a network setup to the bootfile of the pi

### reboot

Reboot your pi

### server

#### ctrl

Start/Stop/Restart the server

#### edit

edit the server config file with vi

#### httpd

Turn on or off httpd checking when starting up the app. This is independent of the dashboard that you choose to run

### shutdown

shut down your kiosk


## Todo

Add checking that you are on the OS that you should be to execute some commands

replace on server refresh scripts with go tool scripts
logs (httpd, php etc....)
kms graphics setup
Commands for pi control
Commands to interact with the api
add networking features
single command for interactive first time setup
startup screen config