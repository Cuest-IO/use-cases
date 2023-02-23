# Github runner
The service run github runner.

# Config
- Create self-hosted runner in your github project
- Follow github guide
- The command ./config.sh --url https://github.com/company/project --token A4TALX2ZEXORB6RV2AELFATD64YMM creates new files .credentials, .credentials_rsaparams, .runner. You have to put in appropriate environment variables the contents of files 

# Run
Use environment variables to setup the service:
- credentials contains .credentials file.
- credentials_rsaparams contains .credentials_rsaparams file.
- runner contains .runner file.

```shell
export credentials='{"scheme": "OAuth","data": {"clientId": "7b9bb43b-3dc4-412c-a172-e51d67c2ed95","authorizationUrl": "https://vstoken.actions.githubusercontent.com/_apis/oauth2/token/18559c37-5443-4fea-b0f8-72b25418c0e7","requireFipsCryptography": "True"}}'
export credentials_rsaparams='{"d": "QoDD3WD7tIFBvEP7Wweln7f0VDZ4BXS5mD07wDzL3iDMYG9DcgGdEaFWESQAAW/hdm85HcZLtbb1GVXqSO1yCwVCeksCzzRdv2qWpSVdP1dw3ujyM3itty7ZhRyt202kFEJFu97CfessWTt6Vj7nf22R+lxfWfnSSsLZ0UMl3qPdOhajUXBkF2EzGrtm+B/2lsA9YiYAazcyl/8Edd+BgPQmc/BZcrTtXEse9zmpyRTVwRD0qL+vRohWpHKeoOmg+ieE6MCrBsNWcezOQxekmHqaQk1l1ewfRmMXqHiRVLcCibvCBrOKWdCvW9avNndDoIMNMYsakokVA4Geh9taAQ==","dp": "zb9oOmj+x4Bj1aFZBi5sTRPLshiFXICAfJjR1Lj1DetzNM4h2t28VLkb4sygEmvrpCqnPmuq3Rq04gYLDWymjc0XZcVOIL/HXnVId56kqiBAwt7sK2e+dHMV2Nbi/IGRDhapZLsfnSNzdWQ2ZBf7jyoPaT8ouqWzCWnViMy0+LE=","dq": "pGd6Czj8t+uOdo7rlKS4fj4+6Hm3D03jp9SMgwqK4BuKevSBfykdSgwvIqke1BCeS3KEm3kqxQ7ouu9Al4siQsZj6myP3TNuFoFmEBEv2P2cxdwqa9qdXtpUiH/mJamiF+n8wRXYNhtbBHO9nIHLPG8Ak03ZgPUWrUUtlnopJ0E=","exponent": "AQAB","inverseQ": "Isipt60mT49H4e6e9htmDvMz+Ul6YMNIZRTGEAkYXOXIG5MGwGwT2OgkxY+yQSN/ZbdTyzCEMKNj4prZUTQ6a3cuXm9TgxEx3LSpzdClEERRckRFMqwS4YoPfw0kZ+u9tEmAP0gGukPo8F3fFqwG+mpKOlISfkhVBa/hkaqiZUo=","modulus": "3cG3thRIE1eV463iIg+w5uG0JIHkR/YxTOmR+yllWWN30rgf5w6njFBwx5M6MJvszgSG50UvEt6PwF8UfJ94BYN/TcfZXqGhfD10UBR0dKn1Ft6/bQHqyv6GNmwmI1Fz0evCLwB5P6namLTxQHCCcKgzlFZNWVfSAIkwvga1s2OGdhXhkwco2lA81KvS8sRCrgjRpUPe86KtTpfx1uCDQHbB1VGF27PfHgD8/uSNKIQROg0eFCRKQMc/vaJW3fUZfj1H5dlMei0OxB2cG2S6QVjIm/H7WuMOySsIV9m97w9+bK5IeM+IG3hW1nYV2ODMAA721pE0U+3yir5KAfUe1Q==","p": "+C3wjc+uge2LqpwChfZmUYhwDgrHKXjtdaskm7pP1Lq+BH3uD/FVS27ODgZO9v212yCG/OdpHSUeiSUr6UTd/hhDCKKhjE6EfJnmd2k0rmV8f+67RpAmu09DSKPf5opJ111FlKFtS6XWnaKv4emSAR6eq51mvnYVvs/g4IRfFxU=","q": "5L6gRKBqh2gTL5F8z2Anq0uytXPwrdtPDI+oNprOo2s04Po/NL++4qd0pJeC5SSmQ8cQdlaXEdSFCNp1ZetNh0A/6zg0L0IOrIMBgKgAy2G+z5t6R6YBNusDNKZz5DGhUlVt6nAFMgww7018VePQL1xqQQDx0yH4uByf9UJf2ME="}'
export runner='{"agentId": 26,"agentName": "app","poolId": 1,"poolName": "Default","serverUrl": "https://pipelines.actions.githubusercontent.com/55TL8SfXXWI5kAks3RZ4eIU9iISv3eq5O25rn5IQSiPgiaQMdi","gitHubUrl": "https://github.com/Cuest-IO/cuest-agent","workFolder": "_work"}'

go run main.go
```
