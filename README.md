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

|     NAME                    |  DEFAULT    |
|-----------------------------|-------------|
| LOG_LEVEL                   |    INFO     |
| VAULT_ADDR                  |             |
| VAULT_CAPATH                |             |
| VAULT_TOKEN                 |             |
| VAULT_SECRET                |             |
| PROPERTIES_FILE             |             |
| VAULT_INSECURE              |    false    |
| VAULT_TOKEN_HANDLER_CRON    | 1 * * * * * |

# Reference

* https://github.com/openlab-red/hashicorp-vault-for-openshift
* https://github.com/raffaelespazzoli/credscontroller
