# Alfred Email Address Manager

A simple Alfred workflow to quickly access and copy your email addresses.

## Features

- Quick access to your predefined email addresses
- Search by name or email address
- Copy email address to clipboard with a single click
- Customizable list of email addresses via workflow settings

## Installation

1. Download the latest release from the [releases page](https://github.com/jiaye/alfred-email-addresses/releases)
2. Double-click the downloaded `.alfredworkflow` file to install

## Usage

1. Type `mail` in Alfred
2. Search for your email address by typing part of the name or email
3. Select an email address to copy it to your clipboard

## Customization

You can customize your email addresses in two ways:

### 1. Through Alfred Workflow Settings

1. Open Alfred Preferences
2. Go to Workflows and select "Email Address Manager"
3. Click the [𝓍] icon in the top right to open workflow configuration
4. Edit the `EMAIL_CONFIG` environment variable using this format:
   
   ```
   Name1:email1@example.com,Name2:email2@example.com,Name3:email3@example.com
   ```
   
   Each name-email pair is separated by a colon (`:`), and pairs are separated by commas (`,`).

### 2. By Editing the Code

To add or modify email addresses directly in the code, edit the `GetEmailAddresses()` function in `internal/emailaddresses.go`.

## License

MIT License
