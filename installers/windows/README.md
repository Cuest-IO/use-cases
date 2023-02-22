# Agent
This service deployes k8s components on local machine.
The agent connects to Connector over websockets and http rest API.
The service receives commands from Connector and shares back to Connector metrics & local details.

The service is developed by Go 18.

Run the agent
```shell
./cuest-agent run --config cuest-agent.yaml
```
where cuest-agent.yaml is your configuration file

Check an agent version:
```shell
./cuest-agent version
```

# Config
Setup a configuration file in cuest-agent.yaml
