# sys-echo

sys-echo is a Golang web server that listens on port 8080 and returns a JSON object containing system information, including the public and private IP addresses, hostname, root filesystem size, and OS version. The server also has a health check path that returns "ok" if the server is running.

## Usage

To run the sys-echo server, you must have the sys-echo binary file on your system. You can download the binary file from the official repository releases page: 
`https://github.com/demoland/sys-echo/releases`

Once you have the sys-echo binary file, you can start the server by running the following command in your terminal:

```bash
./sys-echo
```

This will start the sys-echo server on port 8080. You can access the server by opening a web browser and navigating to http://localhost:8080/.

By default, the server will return a custom message of "Hello, world!". You can customize this message by passing a command-line argument to the server:

```bash
./sys-echo -msg "Custom message here"
```

This will start the server with a custom message of "Custom message here".

## System Information

The JSON object returned by the sys-echo server contains the following fields:

- `public_ip`: the public IP address of the server, derived by querying the ifconfig.co REST API.
- `private_ip`: the private IP address of the server, derived by querying the network interfaces on the server.
- `hostname`: the hostname of the server.
- `root_size`: the size of the root filesystem in bytes.
- `os_version`: the version of the operating system running on the server.
- `custom_message`: the custom message passed in as a command-line argument to the server.

## License

sys-echo is released under the MIT License. See the LICENSE file for details.

## Acknowledgments

sys-echo was inspired by the [sysdig](https://github.com/draios/sysdig) project.
