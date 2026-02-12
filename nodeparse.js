const fs = require('fs');
const yaml = require('js-yaml');

try {
  // Read the YAML file
  const fileContents = fs.readFileSync('./data.yaml', 'utf8');
  
  // Parse the YAML content
  const data = yaml.load(fileContents);
  
  console.log(data.lang);
} catch (e) {
  console.error('Error parsing YAML file:', e.message);
}