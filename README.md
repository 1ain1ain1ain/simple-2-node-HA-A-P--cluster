# Simple HA(Active/Passive) 2 node cluster

### Start server : > *go run master-server.go* (don't forget to change IP address in line 11)

### Start client : > *go run client.go* (don't forget to change IP address in line 56 for your server IP)

### Stop server/client : CTRL + C

### If you want to start it on a single device don't change anything, it will start on a localhost with port 8086

In order to change applications(by default it's Nginx) that starts on the nodes you can add more values in slices on both nodes, for example *CommandList := []string{"systemctl start nginx.service", !!!ADD NEW ACTIVATION COMMANDS HERE!!!}*, after adding more values don't forget to provide deactivation commands on the client *DeactivateCommandList := []string{"systemctl stop nginx", !!!ADD NEW DEACTIVATION COMMANDS HERE!!! }*


Every *Println* line is used for debugging and can be commented or deleted
