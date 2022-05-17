# SSH

## Ignorant

```
-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no
```

## Removant

```
ssh-keygen -f ~/.ssh/known_hosts -R host.name.com
```

