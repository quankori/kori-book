# SSH CheatSheet

- ### Generate SSH

```bash
ssh-keygen -f <path>
```

- ### Generate SSH PEM

```bash
ssh-keygen -t rsa -m PEM -f <path>
```

- ### Generate SSH GIT

```bash
ssh-keygen -t ed25519 -C "example_account@gmail.com"
```

- ### Make sure ssh-agent running

```bash
eval "$(ssh-agent -s)"
```

- ### Add ssh key account to ssh-agent

```bash
ssh-add ~/.ssh/id_ed25519   
```