# Github runner

## Actions Runner Controller

## Usage

To use this self-hosted runner in Kubernetes, follow these steps:

1. Create a K8s cluster, if not available.
```shell
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.yaml
```

2. Next, Generate a Personal Access Token (PAT) for ARC to authenticate with GitHub.

- Login to your GitHub account and Navigate to [Create new Token](https://github.com/settings/tokens/new)
- Select repo.
- Click Generate Token and then copy the token locally ( we’ll need it later).

3. Add repository
```shell
helm repo add actions-runner-controller https://actions-runner-controller.github.io/actions-runner-controller
```

4. Install Helm chart
```shell
helm upgrade --install --namespace actions-runner-system --create-namespace\
  --set=authSecret.create=true\
  --set=authSecret.github_token="REPLACE_YOUR_TOKEN_HERE"\
  --wait actions-runner-controller actions-runner-controller/actions-runner-controller
```

5. Create the GitHub self hosted runners and configure to run against your repository. Apply a [runnerdeployment.yaml](https://github.com/Cuest-IO/utilities/blob/main/usecases/github-runner/deployment.yaml) file:
```shell
kubectl apply -f runnerdeployment.yaml
```

## References

1. [Actions Runner Controller Quickstart](https://github.com/actions/actions-runner-controller/blob/master/docs/quickstart.md)
2. [Automatically scaling runners](https://github.com/actions/actions-runner-controller/blob/master/docs/automatically-scaling-runners.md)

## Contributing

If you wish to contribute to this service, please submit a pull request with your proposed changes.

## License

This service is licensed under the [MIT License](https://opensource.org/licenses/MIT).