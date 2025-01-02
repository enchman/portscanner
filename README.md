# Simple ports scanner
This doesn't exchange any data with the ports, it just checks if they are open or closed.

## Manually run commands

#### Linux 
```shell
nc -zv localhost 1-65535
```
If ``nc``or ``netcat`` not found, try to install the ``netcat`` package.

#### MacOS 
```shell
nc -zv localhost 1-65535
```

#### Windows (PowerShell 5 or above)
See more details on [Test-NetConnection](https://learn.microsoft.com/en-us/powershell/module/nettcpip/test-netconnection?view=windowsserver2025-ps) documentation
```powershell
Test-NetConnection -ComputerName localhost -DiagnoseRouting -InformationLevel Detailed
```

You can also use ``telnet``, but it's very limited

## Tips
You can show the status of the ports on your machine by these commands.

#### Linux

```shell
netstat -lntu
```
If you which to use netstat, you can install ``net-tools``. Otherwise you can try to run this command below.

```shell
ss -lntu
```

#### MacOS 
```shell
lsof -nP -iTCP -sTCP:LISTEN
```

#### Windows
```powershell
netstat -an
```
