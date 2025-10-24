# terraform

## plan targets

```
#!/usr/bin/env python3
import sys
from iterfzf import iterfzf
import subprocess

def execute_command(cmd):
    process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.STDOUT, shell=True, universal_newlines=True)

    output = ""
    for line in iter(process.stdout.readline, ''):
        print(line, end='')  # print in real-time
        output += line  # append to the output string

    process.stdout.close()
    process.wait()

    return output

terraform_plan = execute_command('terraform plan -no-color')

targets = [line.split('#')[1].split('will be')[0].strip() for line in terraform_plan.split('\n') if '# ' in line and ' will be ' in line]

targets = list(set(map(lambda t: '.'.join(t.split('.')[0:2]), targets)))

result = iterfzf(targets, multi=True)

result = '-target=' + ' -target='.join(result)

print('Selected Targets:')
print(result)

```
