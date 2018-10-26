# bedrocksrv

### A Go implementation of the Minecraft: Bedrock Edition server protocol

----

## What's working?
- Configuration management (loading `config.json` and saving the default configuration to `config.json` if it does not yet exist)

## What's new compared to the official Minecraft: Bedrock Edition server?
- Nothing!

## What's left?
- Everything!

## What's planned for the future?
- Getting the server working in a clean and stable manner
- Adding support for plugins built to shared object libraries (`*.so` files)
	- Support for these plugins at this point in time will only work on Linux, and thus plugin support will be disabled on Windows and MacOS hosts
- Adding an in-built DNS server to redirect promoted Bedrock servers to the current one (for use cases like the Nintendo Switch and the Xbox One)

----

## Support
For help and support with bedrocksrv, file a previously unfiled issue on the issues page.

## License
The source code for bedrocksrv is released under the MIT License. See LICENSE for more details.

## Documentation
There is currently no documentation available at this time.

## Donations
All donations are highly appreciated. They help me pay for the server costs to keep development builds of bedrocksrv running and even help point my attention span to bedrocksrv to fix issues and implement newer and better features!

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://paypal.me/JoshuaDoes)
