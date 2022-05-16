## Golang project with Clean-Architecture

### Description
Create a system to control a rover that explore Mars planet

### A rover information
A rover instructure
- F  ...for moving forward 1 block
- B  ...for moving backward 1 block
- L  ...for turning left
- HL  ...for half turning left
- R  ...for turning right
- HR  ...for turning right

A rover rules
1. A rover always start in position 0,0 with facing north
2. A rover can not be in negative position and also can not go out of map
3. A rover will maintain the direction and position when reaching the edge
4. A rover will alway display current position and direction in format {direction}:{position-x},{position-y}
5. A rover's command have to a spesific format and write down into command.txt file in part /ftp
    - First line: will always be a size of a maps resporesent in interger only
    - Next lines: will be an instruction to move or rotate (can have more than one line)

example a command file
```
12
R F5 B HR F
HR B2 L L HL
F HL F4 R F5
B3 R HR F F 
```

Sample input/output

INPUT
```
24
R F L F L L F R
```
OUTPUT
```
    "N:0,0",
    "E:0,0",
    "E:1,0",
    "N:1,0",
    "N:1,1",
    "W:1,1",
    "S:1,1",
    "S:1,0",
    "W:1,0"
```

## The project decscription
In this project, I used **Golang** to develop and design based on **Clean-Architecture**. And the interface is **REST**.

### The project architecture
```
  - config
  - entities
  - ftp
  - infrastructure
       - input
       - router
  - interfaces
       - controllers
       - presenters
  - registry
  - services
       - interactors
       - modules
       - presenters
  - main.go
```

### 4 Layers of **Clean-Architecture** compared with this project

(sorted by inside to outside)
|Layer|Directory|
|--------------------------------|------------------------------------------------------------|
|Layer 1 Entities|- [entities] To store the rover data structure and instruction|
|Layer 2 User case|- [services] To makes the rover move or rotate follow by the command incoming|
|Layer 3 Interface Adapters|- [interfaces] to convert the rover data to fit with the specification data format and respond to a user interface. For example, The rover will use the degree format( 0 45 90 135 180 225 270 315 ) but the specification wants direction format( E NE N NW W SW S SE)|
|Later 4 Frameworks & Driver|- [infrastructure] to setup router and reads the command file|

### How to use it?
1. Clone this project to your local device
```
$ git clone git@github.com:tanakornwry/mars-exploration-project.git
```
2. Run the project on the project's root path (If you haven't some included packages, please install them follow by what the error says)
```
$ go run main.go
```
or install the project and run with installed version (For more detail: https://go.dev/doc/code)

3. After the project runs you can start by this URL (GET) localhost:8080/mars for greeting Mars and (GET) localhost:8080/mars/explore for exploration
4. So you will see explored path

I have set port 8080 as a default, If want want to run with other ports, you can modify the config at /config/config.json
```
"port": "8080"
```

### How to make your exploration path?
If you would like to explore by your command, you could write your command followed by the rover instruction (descript above). And replace it at /ftp/command.txt
