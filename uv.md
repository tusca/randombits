# uv (python)

## init

```
uv init
echo ".venv" >>.gitignore
```

then change the epython version in `.python-version`


## folders

### appname

- appname
  - __init__.py
  - app.py (with def main)

+pyproject.toml:

```
[project.scripts]
myapp = "appname.app:main"

[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[tool.hatch.build.targets.wheel]
packages = ["appname"]
```

then

```
uv sync # for scripts/.venv/...

uv run myapp [parameters]
```

### src folder with multiple apps

- pyproject.toml
- src
  - app1
    - pyproject.toml
    - README.md
    - src/app1
    - tests
      - __init__.py
      - my_test.py
  - app2
    - pyproject.toml
    - README.md
    - src/app2
    - tests
      - __init__.py
      - my_test.py
- libs
  - common
    - pyporject.toml
    - README.md
    - src/common 
  - dev-dependencies
    - pyproject.toml
    - README.md 

#### root pyproject

```
[project]
name = "my-applications"
version = "1.0.0"
description = "multi apps"
readme = "README.md"
requires-python = ">=3.14"
dependencies = [
    "app1",
    "app2,
    "common"
]

[dependency-groups]
dev = [
    "dev-dependencies",
]

[tool.uv.workspace]
members = ["libs/*", "apps/*"]

[tool.uv.sources]
app1 = { workspace = true }
app2 = { workspace = true }
common = { workspace = true }
dev-dependencies = { workspace = true }
```

#### app1 pyproject

```
[project]
name = "app1"
version = "1.0.0"
requires-python = ">=3.14"
authors = [
    {name = "Me"},
]
readme = "README.md"
description = "app1"
dependencies = [
    ...
]
[dependency-groups]
dev = [
  "dev-dependencies",
]

[tool.uv.sources]
common = { workspace = true }

[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[tool.hatch.build.targets.wheel]
packages = ["src/app1"]
only-include = ["src/app1"]

[tool.hatch.build.targets.sdist]
only-include = [
    "src/app1",
    "pyproject.toml",
    "README.md",
]
```

#### common pyproject

```
[project]
name = "common"
requires-python = ">=3.14"
dependencies = [

]
version = "1.0.0"

[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[tool.hatch.build.targets.wheel]
packages = ["src/common"]

[tool.hatch.build.targets.sdist]
only-include = [
    "src/common",
    "pyproject.toml",
    "README.md",
]
```

#### dev-dependencies pyproject

```
[project]
name = "dev-dependencies"
version = "1.0.0"
requires-python = ">=3.14"

dependencies = [
    "pytest>=8.4.2",
]

[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[tool.hatch.build.targets.wheel]
packages = ["src/foo"] # dummy package to allow installation

[tool.hatch.build.targets.sdist]
only-include = [
    "pyproject.toml",
    "README.md",
]
```


