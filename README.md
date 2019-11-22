![GoTasks Header Image](https://github.com/benjamesfleming/gotasks/blob/master/docs/images/header.gif?raw=true)

# GoTasks

GoTasks is a simple self hosted task / todo list.


## Roadmap To V1

- [ ] Responsive Mobile UI
- [ ] Add Google OAuth Support
- [ ] Add Repeating Tasks
- [ ] Optimize API Calls 
- [ ] Docker Containers


## Installation

>**NOTE**: binaries are only provided for linux, follow the [Build From Source]() section to build for a different platform.

The latest GoTasks binaries can be downloaded at [https://dl.benfleming.nz/gotasks](https://dl.benfleming.nz/gotasks). These are built for linux and will not work on other operation systems.

Run the following commands as **root** user, or use **sudo** where needed.

1. `wget https://dl.benfleming.nz/gotasks/gotasks-latest -O /tmp/gotasks`
2. `chmod +x /tmp/gotasks`
3. `mv /tmp/gotasks /usr/local/bin`
4. `gotasks install --help`

At this stage you have installed the gotasks binary onto your server, but still need to add to configuration. Run the following steps to generate the default config file.

5. `mkdir -p /etc/gotasks`
6. `gotasks install -o /etc/gotasks --with-sqlite3`
7. `cp /etc/gotasks/gotasks.sample.toml /etc/gotasks/gotasks.toml`
8. `vi /etc/gotasks/gotasks.toml`

Now you have the default *gotasks* configuration file, but will not be able to start the server until a **Github**  OAuth2 provider has been set up. Follow [the official **Github** guide](https://developer.github.com/apps/building-oauth-apps/creating-an-oauth-app/) on how to get your application credentials.

9. `gotasks start -c /etc/gotasks/gotasks.toml`

Done! The web server should now have started.


## Build From Source

Building for source is not supported at the moment. Please proceed with caution.

**Requirement:**
 
* Go 1.13+
* Node 10+ w/npm

```bash
#!/bin/bash

# clone repo
git clone https://github.com/benjamesfleming/gotasks.git
cd gotasks

# install deps
npm install
go get github.com/GeertJohan/go.rice/rice
go get

# build frontend code
npm run build

# package into binary
rice embed-go -i ./app -i ./app/commands
go build -o build/gotasks main.go
```

## License

MIT License

Copyright (c) 2019 Ben Fleming

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
