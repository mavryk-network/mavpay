# TEZPAY
<p align="center"><img width="100" src="https://raw.githubusercontent.com/tez-capital/tezpay/main/assets/logo.png" alt="TEZPAY logo"></p>

Hey👋 I am PayBuddy close friend of your BakeBuddy 👨‍🍳 you likely know. I am determined to provide you best experience with your baker payouts.
Because of that I prepared something special for you - TEZPAY.

See [Command Reference](https://docs.tez.capital/tezpay/reference/) for details about commands. 

⚠️ **This repo is in active development and hasn't been security audited. Use at your own risk.** ⚠️

## Contributing

To contribute to TEZPAY please read [CONTRIBUTING.md](docs/CONTRIBUTING.md)

## Setup

1. Create directory where you want to store your `tezpay` configuration and reports
	- e.g. `mk tezpay`
2. Head to [Releases](https://github.com/mavryk-network/mavpay/releases) and download latest release and place it into newly created directory
	- on linux you can just `wget -q https://raw.githubusercontent.com/tez-capital/tezpay/main/install.sh -O /tmp/install.sh && sh /tmp/install.sh`
3. Create and adjust configuration file `config.hjson`  See our configuration examples for all available options.
4. ...
5. Run `tezpay pay` to pay latest cycle

## Container

If you want to use tezpay in container, please refer to [container readme](container/readme.md).

## Credits

- TEZPAY [default data collector](https://github.com/mavryk-network/mavpay/blob/main/engines/colletor/default.go#L39) and [default transactor](https://github.com/mavryk-network/mavpay/blob/main/engines/transactor/default.go#L39) (*only available right now*) are **Powered by [TzKT API](https://api.tzkt.io/)**
