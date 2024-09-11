# MAVPAY
<p align="center"><img width="100" src="https://raw.githubusercontent.com/mavryk-network/mavpay/main/assets/logo.png" alt="MAVPAY logo"></p>

Hey👋 I am PayBuddy close friend of your BakeBuddy 👨‍🍳 you likely know. I am determined to provide you best experience with your baker payouts.
Because of that I prepared something special for you - MAVPAY.

See [Command Reference](https://mavpay.mavryk.org/mavpay/reference/) for details about commands. 

⚠️ **This repo is in active development and hasn't been security audited. Use at your own risk.** ⚠️

## Contributing

To contribute to MAVPAY please read [CONTRIBUTING.md](docs/CONTRIBUTING.md)

## Setup

1. Create directory where you want to store your `mavpay` configuration and reports
	- e.g. `mk mavpay`
2. Head to [Releases](https://github.com/mavryk-network/mavpay/releases) and download latest release and place it into newly created directory
	- on linux you can just `wget -q https://raw.githubusercontent.com/mavryk-network/mavpay/main/install.sh -O /tmp/install.sh && sh /tmp/install.sh`
3. Create and adjust configuration file `config.hjson`  See our configuration examples for all available options.
4. ...
5. Run `mavpay pay` to pay latest cycle

## Container

If you want to use mavpay in container, please refer to [container readme](container/readme.md).

## Credits

- MAVPAY [default data collector](https://github.com/mavryk-network/mavpay/blob/main/engines/colletor/default.go#L39) and [default transactor](https://github.com/mavryk-network/mavpay/blob/main/engines/transactor/default.go#L39) (*only available right now*) are **Powered by [MvKT API](https://api.mavryk.network/)**
