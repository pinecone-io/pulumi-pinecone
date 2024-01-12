## Provider Developers

### Prerequisites

Prerequisites for this repository are already satisfied by the [Pulumi Devcontainer](https://github.com/pulumi/devcontainer) if you are using Github Codespaces, or VSCode.

If you are not using VSCode, you will need to ensure the following tools are installed and present in your `$PATH`:

* [`pulumictl`](https://github.com/pulumi/pulumictl#installation)
* [Go 1.21](https://golang.org/dl/) or 1.latest
* [NodeJS](https://nodejs.org/en/) 14.x.  We recommend using [nvm](https://github.com/nvm-sh/nvm) to manage NodeJS installations.
* [Yarn](https://yarnpkg.com/)
* [TypeScript](https://www.typescriptlang.org/)
* [Python](https://www.python.org/downloads/) (called as `python3`).  For recent versions of MacOS, the system-installed version is fine.
* [.NET](https://dotnet.microsoft.com/download)


### Build & test the boilerplate XYZ provider

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/pinecone-io/pulumi-pinecone)

1. [Open repo in GitHub Codespaces](https://codespaces.new/pinecone-io/pulumi-pinecone).
1. Open a terminal in the CodeSpaces IDE.
1. Run `make build install` to build and install the provider.
1. Export your Pinecone API Token: `export PINECONE_API_KEY="0000b000-86cf-4ecf-a753-00000000000"`
1. Run `make up` to run the example program in `examples/yaml`.
1. Run `make down` to tear down the example program.

This will:

1. Create the SDK codegen binary and place it in a `./bin` folder (gitignored)
2. Create the provider binary and place it in the `./bin` folder (gitignored)
3. Generate the dotnet, Go, Node, and Python SDKs and place them in the `./sdk` folder
4. Install the provider locally.
5. Run the example program in `examples/yaml` using the locally built provider.
6. Tear down the example program.

#### Test against the example

```bash
$ cd examples/simple
$ yarn link @pinecone-database/pinecone
$ yarn install
$ pulumi stack init test
$ pulumi up
```

#### A brief repository overview

You now have:

1. A `provider/` folder containing the building and implementation logic
    1. `cmd/pulumi-resource-pinecone/main.go` - holds the provider's sample implementation logic.
2. `deployment-templates` - a set of files to help you around deployment and publication
3. `sdk` - holds the generated code libraries created by `pulumi-gen-xyz/main.go`
4. `examples` a folder of Pulumi programs to try locally and/or use in CI.
5. A `Makefile` and this `README`.

#### Additional Details

This repository depends on the pulumi-go-provider library. For more details on building providers, please check
the [Pulumi Go Provider docs](https://github.com/pulumi/pulumi-go-provider).
