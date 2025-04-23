#!/usr/bin/env bash

go build -o alfred-email-addresses
zip alfred-email-addresses.alfredworkflow ./alfred-email-addresses ./info.plist ./icon.png
rm ./alfred-email-addresses
