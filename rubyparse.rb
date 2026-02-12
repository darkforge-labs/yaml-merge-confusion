require 'yaml'

data = YAML.load_file('./data.yaml', aliases: true)

puts data['lang']
