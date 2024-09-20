# Migrating from BC

⚠️⚠️ **ledger wallet mode is not supported by `mavpay` yet** ⚠️⚠️

`mavpay` is able to build its config from preexisting BC configuration. So all you have to do is to use your old BC config and let `mavpay` to migrate it.

Your configuration gets migrated automatically on `pay` or `generate-payouts`. Old BC configuration will be saved to `config.backup.hjson`

You can run `mavpay` same way as BC - `mavpay pay --cycle=540` or `mavpay pay` for last completed cycle 😉

NOTE: *During BC migration `mavpay` injects 5% donation to your new `config.hjson` to support `mavpay` development. This is entirely optional. Set it as you see fit.*

## If you operate remote signer

`mavpay` does not touch configuration of your signers. To use remote signer with `mavpay` you have to change `public_key` to `pkh` in your `remote_signer.hjson`

For example:
```hjson
public_key: mv1HCXRedE7zVSwmSqxDe3XZcMPLeF7xYqP3
url: http://127.0.0.1:2222
```
becomes:
```hjson
pkh: mv1HCXRedE7zVSwmSqxDe3XZcMPLeF7xYqP3
url: http://127.0.0.1:2222
```