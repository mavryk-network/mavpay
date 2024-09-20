# MavPay Container Readme
MavPay is a Mavryk reward distributor that simplifies the process of distributing rewards to your stakeholders. This readme provides instructions on how to use the [mavryk-network/mavpay](ghcr.io/mavryk-network/mavpay) container image, which comes with both [mavpay](https://github.com/mavryk-network/mavpay) and [eli](https://github.com/alis-is/eli) preinstalled.

## Prerequisites
Docker installed on your system.
## Usage
1. Pull the MavPay container image:
```bash
docker pull ghcr.io/mavryk-network/mavpay
```
2. Run the MavPay container with the desired command. Replace `[command]` with the desired MavPay command and `[options]` with the corresponding command options:
```bash
docker run --rm -it -v $(pwd):/mavpay ghcr.io/mavryk-network/mavpay [command] [options]
```

Here are some examples of how to use the MavPay commands:

Generate payouts:
```bash
docker run --rm -it -v $(pwd):/mavpay ghcr.io/mavryk-network/mavpay generate-payouts --cycle <cycle_number> [flags]
```
Replace `<cycle_number>` with the desired cycle number for which you want to generate payouts.

Continual payout (executed by default if no commands or arguments are provided):
```bash
docker run --rm -it -v $(pwd):/mavpay ghcr.io/mavryk-network/mavpay continual [flags]
```

Manual payout:
```bash
docker run --rm -it -v $(pwd):/mavpay ghcr.io/mavryk-network/mavpay pay --cycle <cycle_number> [flags]
```

**Note**: When running the container, make sure to mount the current working directory (or the desired directory containing your MavPay configuration) to the `/mavpay` path inside the container. This ensures that the container has access to your configuration files and can write any generated files back to your host system. Your MavPay configuration should be named `config.hjson`. Payout reports will be stored in the mounted directory under the `reports` directory by default.

For more information about available commands and their options, refer to the provided MavPay help:

```bash
docker run --rm -it ghcr.io/mavryk-network/mavpay help
```

## Support
For any questions or issues related to the container or MavPay, please visit the GitHub repository at [mavryk-network/mavpay](https://github.com/mavryk-network/mavpay) or submit an issue there.
