# Git Bits

## Setups

```
git config --global user.email "EMAIL"
git config --global user.name "FIRST LAST"
git config --global push.default simple 
```

## Store vs Cache pwds

```
git config --global credential.helper store
git config --global credential.helper 'cache --timeout 9999999'
```

## Not recommended

```
git config --global http.sslVerify false

# or

git -c http.sslVerify=false clone
```



## LFS

```
git lfs installs
git lfs strack "*.psd"
git add .gitattributes
git add file.psd
git commit -m "psd file"
git push origin master
```


## Others

```
git remote set-url origin URL
```

## Git Worktree

https://git-scm.com/docs/git-worktree

A git repository can support multiple working trees, allowing you to check out more than one branch at a time. With git worktree add a new working tree is associated with the repository. This new working tree is called a "linked working tree" as opposed to the "main working tree" prepared by "git init" or "git clone". A repository has one main working tree (if it’s not a bare repository) and zero or more linked working trees.

## Aliases

### done

Checkout master & remove the previous branch (https://stackoverflow.com/a/48705399)

```
git config --global alias.done '!f() { git checkout master && git branch -D @{-1} && git pull origin master; }; f'
```

## Auto complete

https://apple.stackexchange.com/a/55886



