This is the Docker image for the [Atlas CLI](https://www.mongodb.com/docs/atlas/cli/stable/).

The Atlas CLI is a command line interface built specifically for [MongoDB Atlas](https://www.mongodb.com/docs/atlas/). 
Use the Atlas CLI to interact with your Atlas database deployments and Atlas Search from the terminal with short, 
intuitive commands, so you can accomplish complex database management tasks in seconds.

## Getting Started

1. **Pull the Docker image.** 
<br>
To pull the latest Docker image, run `docker pull mongodb/atlas`.
<br><br>
If you run `docker pull mongodb/atlas` without specifying a version tag, Docker automatically pulls the latest version 
of the Docker image (`mongodb/atlas:latest`). 
<br><br>
To pull a specific version of the Docker image, run the following command, replacing _tag_ with the version tag:
`docker pull mongodb/atlas:tag`.
<br><br>

2. **Get a shell.** 
<br>
Run the command to get a shell in interactive mode: `docker run --rm -it mongodb/atlas bash`
<br> <br>

3. **Authenticate and run Atlas CLI commands.** 
<br>
Follow the steps to [connect from the atlas CLI](https://www.mongodb.com/docs/atlas/cli/stable/connect-atlas-cli/). 
For example, you can run atlas `auth login to authenticate`. 
<br><br>
After you authenticate, you can run [Atlas CLI commands](https://www.mongodb.com/docs/atlas/cli/stable/command/atlas/). 
For example, you can run `atlas --help` to learn about available commands.

## Getting Help

Support is available via the [Atlas support process](https://www.mongodb.com/docs/atlas/support/).
To learn more about the Atlas CLI, see the [Atlas CLI documentation](https://www.mongodb.com/docs/atlas/cli/stable/).

## Contributing

See our [contributing guide](https://github.com/mongodb/mongodb-atlas-cli/blob/master/CONTRIBUTING.md) 
to contribute to the [GitHub repository](https://github.com/mongodb/mongodb-atlas-cli).

## License

MongoDB releases the Atlas CLI under the Apache 2.0 license. See the 
[license file](https://github.com/mongodb/mongodb-atlas-cli/blob/master/LICENSE) for more information.
