# Genome Data API Client SDKs

This repository contains the source code for a wrapper around UCSC's Genome Browser REST API. 

By creating a local .NET 8 web API that mirrors requests to https://api.genome.ucsc.edu/, we can leverage .NET 8's [automatic Swagger documentation](https://learn.microsoft.com/en-us/aspnet/core/tutorials/web-api-help-pages-using-swagger?view=aspnetcore-8.0#swagger-ui) to automatically generate an OpenAPI spec (```Swagger/UCSCGenomeAPI.json```) on build that models *most* of the behavior of the Genome Browser's API. 

```bash
Swagger
└── UCSCGenomeAPI.json
```

Then, we can use this spec to automatically generate API SDKs with [Kiota](https://learn.microsoft.com/en-us/openapi/kiota/).
The client SDKs are already built and organized in the `Client` directory, categorized by programming language.

```bash
Client/
├── csharp
├── go
├── java
├── php
├── python
├── ruby
└── typescript
```

Refer to the [Kiota Quickstart](https://learn.microsoft.com/en-us/openapi/kiota/quickstarts/) to learn how to use each SDK with your language of choice.


# Local Development

This section is intended for developers who wish to run and build the C# API locally. Read on if you are setting up the Genome Data API for local development or customization.

## Prerequisites

- [.NET 8 SDK](https://dotnet.microsoft.com/download/dotnet/8.0): Ensure you have the latest version of the .NET 8 SDK installed on your machine. 

## Getting Started

To get started with this project, clone the repository to your local machine and navigate into the project directory:

```bash
git clone https://github.com/mgardos01/UCSC-Genome-API-Wrapper.git
cd UCSC-Genome-API-Wrapper
```

### Restore and Build

Before running the application, restore the NuGet packages and build the project:

```bash
dotnet restore
dotnet build [-p:GenerateSDKs=true]
```

These commands will install all necessary dependencies and compile the project for use. SDKs aren't regenerated on build by default. 

## Running the Application

To run the application locally, use the following command:

```bash
dotnet run
```

This command will start the proxy server hosted at `http://localhost:5099`. You can access the Swagger UI to interact with the API at `http://localhost:5099/swagger`. You can change this configuration in ```Properties/launchSettings.json```. 

