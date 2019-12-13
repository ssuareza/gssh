# gssh
SSH client to connect AWS instances.

## Installation

### MAC:
```
brew tap ssuareza/brew git@github.com:ssuareza/homebrew-brew
brew install ssuareza/brew/gssh -f
```

### Linux:
```
curl -sLo gssh curl -sLo gssh https://github.com/ssuareza/gssh/releases/download/v0.0.2/gssh-v0.0.2-linux-amd64
chmod +x gssh
sudo mv gssh /usr/local/bin/
```

## Upgrade

### MAC:
```
brew update
brew upgrade ssuareza/brew/gssh -f
```

### Linux:
Follow installation steps.

## Configuration

`gssh` it will be configured the first time launched. All configuration will be saved in "$HOME/.gssh/config.yaml".

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

## Profiles
`gssh` supports multiple aws profiles. To use it pass a list of profiles during configuration.

If you want to overwite the values just edit "$HOME/.gssh/config.yaml". Example:
```
aws:
  profile: prod,dev
  region: us-east-1
ssh:
  bastion: bastion.domain.com
  port: 22
  user: myuser
```