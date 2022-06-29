# iptrack

used as a dns tracker for my dev machines -

iptrack keeps an in-ram listing of my machines IPs then serves
them up as a custom /etc/hosts file

-- compile for pihole

env GOOS=linux GOARCH=arm64 GOARM=5 go build

Usage: iptracker <arg>

Args -

    list    - lists all the current machines
    hosts   - creates the host file
    update  - updates database entry of host with the current assigned
               ip
