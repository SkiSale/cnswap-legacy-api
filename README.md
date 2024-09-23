# CnSwap Legacy API

Provides a RESTful  API to the legacy CnSwap system for integration with Gearshift and other third party services.

It currently only provides read access to Equipment and Vendor records to support MC9090 lookup app.

# Deployment

This service is intended to be run in an isolated docker container. It provides no authentication and is expected to be run in a private environment

There are two dependencies

- Access to FoxPro database
- ~~Kafka Connection~~

The service expects a readable copy of the CnSwap FoxPro database to be accessible within the container. Mount the full data directory as a volume (suggest at /data) and specify the full path to the dbc file [CnSData.dbc] in the appropriate DBC_PATH environment variable
## Run Locally
To run the project locally you will need access to a full backup of CnSwap's data directory.
```bash
  git clone https://github.com/SkiSale/cnswap-legacy-api
```

Go to the project directory

```bash
  cd github.com/SkiSale/cnswap-legacy-api
```

Install dependencies

```bash
  go install
```

Start the server

```bash
  go run .
```


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file or docker configuration. An .env.example file is provided with further details

`PORT`

`DBC_PATH`

`LOG_LEVEL`

## Authors

- [@brad82](https://www.github.com/brad82)

