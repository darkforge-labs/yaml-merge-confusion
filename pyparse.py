import yaml
import sys

f = open('data.yaml', 'r')
doc = yaml.safe_load(f)
print(doc["lang"])
f.close()

