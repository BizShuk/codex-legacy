# GenericRundeckTomcatDeployment
A generic tomcat application deployment on Rundeck.

[Sample project](jobs.xml)

- [How to import](#how-to-import)
- [How to config for customized env](#how-to-config-for-customized-env)
- [BackUp](#backup)
    - [rd access and debug configuration](#rd-access-and-debug-configuration)
    - [troubleshooting](#troubleshooting)

### How to import
change basic configuration and use rd tool to import.
`rd jobs load -f /XmlFilePath -p <ProjectName>`

### How to config for customized env
1. Duplicat server job and config with customized server ip addreass and path
2. Duplicate application job with application name, version, and package, and config server job reference to customized server job

### BackUp
- [rd](https://rundeck.github.io/rundeck-cli/) tool, backup tool created by rundeck
- job xml backup cmd: `rd jobs list -f /StoredPath -p <ProjectName>` 

##### rd access and debug configuration
Configuration below can be configured in `~/.rd/rd.conf` (export RD_CONF=[Conf path], but unknown format.


    # Server URL
    export RD_URL=${RD_URL:-http://server:4440}/api/<version>, 4440 for http, 4443 for https.

    # Access credential
    export RD_TOKEN=....
    # or
    export RD_USER=username
    export RD_PASSWORD=password

    # Debug mode
    export RD_DEBUG=1 # basic http request debug
    export RD_DEBUG=2 # http headers
    export RD_DEBUG=3 # http body

    # Security Check
    export RD_INSECURE_SSL=true
    export RD_INSECURE_SSL_HOSTNAME=true

##### troubleshooting

1. IOException: unexpected end of stream on okhttp3.Address@89fa6f1. Resolved by using right linux user `rundeck`

2. PKIX path building failed. unable to find valid certification path to requested target. Resolved by `export RD_INSECURE_SSL=true`

3. api version is not supported. (server:20 rd:21). Resolved by turning on debug more and skipping security check.


