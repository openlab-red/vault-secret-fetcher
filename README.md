# Hashicorp Vault Agent Token Handler


```
Vault Agent Token Handler sidecar container for kubernetes applications

Usage:
  vault-agent-token-handler [command]

Available Commands:
  handler     handler retrieves credentials managed by the vault agent
  help        Help about any command

Flags:
  -h, --help   help for vault-agent-token-handler

Use "vault-agent-token-handler [command] --help" for more information about a command.

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
| VAULT_TOKEN_HANDLER_CRON    | 1 * * * * * |    Cron Scheduler for the token handler                         |
| PROPERTIES_FILE             |             |    Target properties file to save the decrypted secret          |
| PROPERTIES_TYPE             |     yaml    |    Properties output format                                     |

# Reference

* https://github.com/openlab-red/hashicorp-vault-for-openshift
* https://github.com/raffaelespazzoli/credscontroller
