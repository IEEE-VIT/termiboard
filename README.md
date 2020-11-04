# Termiboard
[![Open Source Love svg1](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/IEEE-VIT/termiboard/)
[![GitHub license](https://img.shields.io/github/license/IEEE-VIT/termiboard.svg)](https://github.com/IEEE-VIT/termiboard/blob/master/LICENSE)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/IEEE-VIT/termiboard.svg)](https://github.com/IEEE-VIT/termiboard)
[![GitHub Workflow Build Passing](https://img.shields.io/github/workflow/status/IEEE-VIT/termiboard/termiboard-build)](https://github.com/IEEE-VIT/termiboard)
[![GitHub Open Issues](https://img.shields.io/github/issues-raw/IEEE-VIT/termiboard)](https://github.com/IEEE-VIT/termiboard/issues?q=is%3Aopen+is%3Aissue)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/IEEE-VIT/termiboard/issues/new/choose)

The cross platform CLI Dashboard which shows all your system stats and info in an instant.

## Getting Started
* Fork it.

* Clone your forked repo and move inside it:

`git clone https://github.com/<your-github-username>/termiboard.git && cd termiboard`

* Checkout to a new branch to work on an issue:

`git checkout -b my-amazing-feature`

* Run termiboard
```console
go run .
```

* Build termiboard
```console
go build
./termiboard
```

* Once you're all done coding, it's time to open a PR :)
Run the following commands from the root of the project directory:

`git add .`

`git commit -m "A short description about the feature."`

`git push origin <my-amazing-feature>`

Open your forked repo in your browser and then raise a PR to the `master` branch of this repository!

## Usage
Once the executable file is created, the program can be executed as:

### Without Flag Options
This will by default show all the parameters of the System.
```console
./termiboard
```
Output:   
![Complete Output SS](assets/Termiboard.PNG)
### With Flag Options
If the flag options are passed, then the specific parameters of the System will be displayed.
```console
./termiboard --cpu-usage
```
Output:       
![CPU Info Output SS](assets/CpuInfo.PNG)
#### All Available Options
```console
./termiboard --help
```
| Flag        | Function                                        |
|-------------|-------------------------------------------------|
| --all       | Show all stats                                  |
| --cpu-info  | Show CPU information                            |
| --cpu-usage | Show CPU usage                                  |
| --disk      | Show disk usage                                 |
| --local-ip  | Show local IP address                           |
| --public-ip | Show public IP address                          |
| --ram       | Show RAM usage                                  |
| --top5-ram  | Show top 5 process that consume the most memory |

## Contributing
To start contributing, check out [CONTRIBUTING.md](https://github.com/IEEE-VIT/termiboard/blob/master/CONTRIBUTING.md). New contributors are always welcome to support this project. Kindly consider leaving a :star: if you like the project :)

## License
This project is licensed under [MIT](https://github.com/IEEE-VIT/termiboard/blob/master/LICENSE).
