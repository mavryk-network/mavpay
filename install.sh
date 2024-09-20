#!/bin/sh

TMP_NAME="./$(head -n 1 -c 32 /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32)"
PRERELEASE=false
if [ "$1" = "--prerelease" ]; then
    PRERELEASE=true
fi

if which curl >/dev/null; then
    if curl --help 2>&1 | grep "--progress-bar" >/dev/null 2>&1; then
        PROGRESS="--progress-bar"
    fi

    set -- curl -L $PROGRESS -o "$TMP_NAME"
    if [ "$PRERELEASE" = true ]; then
        LATEST=$(curl -sL https://api.github.com/repos/mavryk-network/mavpay/releases | grep tag_name | sed 's/  "tag_name": "//g' | sed 's/",//g' | head -n 1 | tr -d '[:space:]')
    else
        LATEST=$(curl -sL https://api.github.com/repos/mavryk-network/mavpay/releases/latest | grep tag_name | sed 's/  "tag_name": "//g' | sed 's/",//g')
    fi
else
    if wget --help 2>&1 | grep "--show-progress" >/dev/null 2>&1; then
        PROGRESS="--show-progress"
    fi
    set -- wget -q $PROGRESS -O "$TMP_NAME"
    if [ "$PRERELEASE" = true ]; then
        LATEST=$(wget -qO- https://api.github.com/repos/mavryk-network/mavpay/releases | grep tag_name | sed 's/  "tag_name": "//g' | sed 's/",//g' | head -n 1 | tr -d '[:space:]')
    else
        LATEST=$(wget -qO- https://api.github.com/repos/mavryk-network/mavpay/releases/latest | grep tag_name | sed 's/  "tag_name": "//g' | sed 's/",//g')
    fi
fi

if ./mavpay version | grep "$LATEST"; then
    echo "Latest mavpay already available."
    exit 0
fi

PLATFORM=$(uname -m)
# remap platform
if [ "$PLATFORM" = "x86_64" ]; then
    PLATFORM="amd64"
elif [ "$PLATFORM" = "aarch64" ]; then
    PLATFORM="arm64"
fi
echo "Downloading mavpay-linux-$PLATFORM $LATEST..."

if "$@" "https://github.com/mavryk-network/mavpay/releases/download/$LATEST/mavpay-linux-$PLATFORM" &&
    mv "$TMP_NAME" ./mavpay &&
    chmod +x ./mavpay; then
    echo "mavpay $LATEST for $PLATFORM successfuly installed."
else
    echo "mavpay installation failed!" 1>&2
    exit 1
fi
