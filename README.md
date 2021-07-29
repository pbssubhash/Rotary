
    █▀▀█ █▀▀█ ▀▀█▀▀ █▀▀█ █▀▀█ █░░█
    █▄▄▀ █░░█ ░░█░░ █▄▄█ █▄▄▀ █▄▄█
    ▀░▀▀ ▀▀▀▀ ░░▀░░ ▀░░▀ ▀░▀▀ ▄▄▄█
    Built by pbssubhash       v0.1
    
## What is this?
Rotary is a simple tool that allows you to run multiple commands concurrently. It allows you to create a template/flow that you can reuse elsewhere.

## Template Structure
All templates need to adhere to the following structure. Explaination of each element is mentioned below.
```
--- 
apiVersion: 0.1
basicInfo: 
  author: pbssubhash
  name: ReconProfile
  targetOs: "Windows" # Accepted value is Windows/Linux for now.
userInputs:
  firstParam: "Enter subcommand for whoami.exe"
concurrent: true
commands: 
  - name: "First Block"
    concurrent: true
    commands:
      - "first_command %{firstParam}"
      - "second_command"
  - name: "Second Block"
    concurrent: false
    commands: 
      - "third_command %{anotherParam}"
```

|S.No | Name of the attribute  | Description of the attribute  | Required/Optional  |  Data Type |
|---|---|---|---|---|
| 1.  |  apiVersion | This is reserved for future backward compatibility. | Required  |  float  |
| 2.  |  basicInfo | Block which will contain basic information  |  Required | YAML object |
| 3.  |  basicInfo.author | Name of the author |  Required | String |
| 4.  |  basicInfo.name | Name of the template  |  Required | String |
| 5.  |  basicInfo.targetOS | Name of the OS that the template is targeted to. This can be used to control if the commands inside is supported inside Linux or Windows; Currently supports Linux or Windows  |  Required | String |
| 6.  |  userInputs | Block which will contain the parameters that are required as an input for the template to run. Advisable to use camelCase, coz it's cute. |  Required | String |
| 7.  |  concurrent | This is true if the sub-blocks should run concurrently/async.  |  Required | bool |
| 8.  |  commands | Block that will contain sub commands.  |  Required | String |
| 9.  |  commands.name | Name of the sub-command  |  Required | String |
| 10.  |  commands.concurrent | This is true if the commands in the sub-block should run concurrently/async  |  Required | String |
| 11.  |  commands.commands | List of all sub commands |  Required | String |

## How to start working!

#### Download and run (The easy way)
1. Download compiled binaries from here: [Download Rotary Compiled](https://github.com/pbssubhash/Rotary/tree/main/build)
#### Building from the source  (The long way)
1. Install Golang [Download and Install Golang](https://golang.org/dl/)
2. Ensure Go (`go`) is in Path
2. Run the following commands 
```
git clone github.com/pbssubhash/Rotary
cd Rotary
go build main.go
./main

    █▀▀█ █▀▀█ ▀▀█▀▀ █▀▀█ █▀▀█ █░░█
    █▄▄▀ █░░█ ░░█░░ █▄▄█ █▄▄▀ █▄▄█
    ▀░▀▀ ▀▀▀▀ ░░▀░░ ▀░░▀ ▀░▀▀ ▄▄▄█
    Built by pbssubhash       v0.1
https://github.com/pbssubhash/rotary

Required argument not present.
Usage:
  -c string
        -c ../webrecon.yml (default "Default")
  -t int
         -t 100 (default 10)
exit status 1
```

## FAQ:
1. I have an issue/doubt/concern? 
A. Let us know about it over here. [Issues](https://github.com/pbssubhash/Rotary/issues/new)


## Request from maintainers:
If you built something cool, please contribute to this repository or let us know, we'll feature it here :-) 
