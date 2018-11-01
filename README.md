# Hashicorp Vault Secret Fetcher


```
Vault Secret Fetcher sidecar container for kubernetes applications

Usage:
  vault-secret-fetcher [command]

Available Commands:
  start     start retrieves credentials managed by the vault agent
  help      Help about any command

Flags:
  -h, --help   help for vault-secret-fetcher

Use "vault-secret-fetcher [command] --help" for more information about a command.

```

## Environment variable

|     NAME                    |  DEFAULT    |  DESCRIPTION                                                    |
|-----------------------------|-------------|-----------------------------------------------------------------|
| LOG_LEVEL                   |    INFO     |    Log level from [logrus](https://github.com/sirupsen/logrus)  |
| VAULT_ADDR                  |             |    Vault Address                                                |
| VAULT_CAPATH                |             |    Vault CA                                                     |
| VAULT_TOKEN                 |             |    Vault Agent sink file path                                   |
| VAULT_SECRET                |             |    Vault Secret to retrieve                                     |
| VAULT_INSECURE              |    false    |    TLS Skip                                                     |
| VAULT_SECRET_FETCHER_CRON   | 1 * * * * * |    Cron Scheduler for the secret fetcher                        |
| PROPERTIES_FILE             |             |    Target properties file to save the decrypted secret          |

## Make vault-secret-fetcher image available in OpenShift

1. Build

    ```
    oc project openshift

    oc new-build --name vault-secret-fetcher https://github.com/openlab-red/vault-secret-fetcher
    ```

2. Check the Image Stream.

    ```
    oc get is vault-secret-fetcher
    ```

# Reference

* https://github.com/openlab-red/hashicorp-vault-for-openshift
* https://github.com/raffaelespazzoli/credscontroller
