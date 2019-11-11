# gssh
SSH client to connect AWS instances.

## Installation

Go to the [releases](https://github.com/ssuareza/gssh/releases) page and download the latest binary for your platform. Make sure it is executable (`chmod +x tf`) and move it to a folder on your path (such as `/usr/local/bin` or `/usr/bin`).

## Configuration

`gssh` it will be configured the first time launched. All configuration will be saved in "$HOME/.gssh".

## Usage

Just execute `gssh` and select the InstanceID. Example:

```
gssh
InstanceID              Name            PrivateIP       PublicIP
i-027f3873ebf0b2bC4     http.prod       172.29.18.95
i-0af353e1458d7f5a4     http.prod       172.28.18.129
i-03cccaac566dd2756     app.prod        172.29.17.68
i-0208d936b38ea7e22     app.prod        172.29.17.57
i-0e06e157bc5fa8611     http.dev        172.29.16.48
i-037d57f116d8f8292     app.dev         172.29.19.25

Select InstanceID: i-027f3873ebf0b2bC4
````
