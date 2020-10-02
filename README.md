<p align="center"><img width="40%" src="https://hacktoberfest.digitalocean.com/assets/HF-full-logo-b05d5eb32b3f3ecc9b2240526104cf4da3187b8b61963dd9042fdc2536e4a76c.svg"/></p>

# Hacktoberfest 2020 with IEEE-VIT :heart:
## Termiboard
The CLI Dashboard to keep check on your CPU, Memory and Network!

Support open source software by participating in [Hacktoberfest](https://hacktoberfest.digitalocean.com) and get goodies and a free t-shirt! :yellow_heart:

> Please check all issues labelled as `hacktoberfest` to start contributing!

Kindly consider leaving a :star: if you like the repository and our organisation.

## Getting Started
* Fork it.

* Clone your forked repo and move inside it:

`git clone https://github.com/<your-github-username>/termiboard.git && cd termiboard`

* Checkout to a new branch to work on an issue:

`git checkout -b my-amazing-feature`

* Run termiboard

```console
go run *.go
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
This will by default show all the parameters of the System. It will provide all the parameters of CPU, Memory and Network Status.
```console
./termiboard
```
Output:   
![Termiboard Output SS](images/Termiboard.PNG)
### With Flag Options
If the flag options are passed, then the specific parameters of CPU, Memory and Network Status will be displayed.
```console
./termiboard --cpu-usage
```
Output:       
![CPU Info Output SS](images/CpuInfo.PNG)
#### All Available Options
```console
./termiboard --help
```
| Flag      | Function                                        |
|-----------|-------------------------------------------------|
| all       | Show all stats                                  |
| cpu-info  | Show CPU information                            |
| cpu-usage | Show CPU usage                                  |
| disk      | Show disk usage                                 |
| local-ip  | Show local IP address                           |
| public-ip | Show public IP address                          |
| ram       | Show RAM usage                                  |
| top5-ram  | Show top 5 process that consume the most memory |

## Contributing
To start contributing, check out [CONTRIBUTING.md](https://github.com/IEEE-VIT/termiboard/blob/master/CONTRIBUTING.md). New contributors are always welcome to support this project. If you want something gentle to start with, check out issues labelled as `easy` or `good-first-issue`. Check out issues labelled as `hacktoberfest` if you are up for some grabs! :)

## License
This project is licensed under [MIT](https://github.com/IEEE-VIT/termiboard/blob/master/LICENSE).
