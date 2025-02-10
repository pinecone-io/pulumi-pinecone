# Pinecone Pulumi Provider

<img src="img/pinecone.svg" width="50%">

###  ☁ Open in the Cloud 
[![Open in VS Code](https://img.shields.io/badge/Open%20in-VS%20Code-blue?logo=visualstudiocode)](https://vscode.dev/github/pinecone-io/pulumi-pinecone)
[![Open in Glitch](https://img.shields.io/badge/Open%20in-Glitch-blue?logo=glitch)](https://glitch.com/edit/#!/import/github/pinecone-io/pulumi-pinecone)

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/pinecone-io/pulumi-pinecone)
[![Open in CodeSandbox](https://assets.codesandbox.io/github/button-edit-lime.svg)](https://codesandbox.io/s/github/pinecone-io/pulumi-pinecone)
[![Open in StackBlitz](https://developer.stackblitz.com/img/open_in_stackblitz.svg)](https://stackblitz.com/github/pinecone-io/pulumi-pinecone)
[![Open in Repl.it](https://replit.com/badge/github/withastro/astro)](https://replit.com/github/pinecone-io/pulumi-pinecone)

[![Open in Codeanywhere](https://codeanywhere.com/img/open-in-codeanywhere-btn.svg)](https://app.codeanywhere.com/#https://github.com/pinecone-io/pulumi-pinecone)
[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/pinecone-io/pulumi-pinecone)


This Pulumi Pinecone Provider enables you to manage your [Pinecone](https://www.pinecone.io/) collections and indexes using any language of Pulumi Infrastructure as Code.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pinecone-database/pulumi
```

or `yarn`:

```bash
yarn add @pinecone-database/pinecone
```

### Python

To use from Python, install using `pip`:

```bash
pip install pinecone_pulumi
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pinecone-io/pulumi-pinecone/sdk
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package PineconeDatabase.Pinecone
```

## Configuration

The following configuration points are available for the `pinecone` provider:

- `pinecone:APIKey` - This is the Pinecone API key. (environment: `PINECONE_API_KEY`)
