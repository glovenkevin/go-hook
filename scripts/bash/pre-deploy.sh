#!/bin/bash

echo "Pre-deploy script"

# create .env file with port 9994
echo "PORT=9994" > .env

# create ecosystem.config.js file
echo "module.exports = {
	apps : [
	  {
		name   : \"golang-hook\",
		script : \"app\",
		interpreter_args: \"-u\",
		autorestart: false,
		instances: 1,
		env_development: {
		  NODE_ENV: \"development\",
		},
		env_production: {
		  NODE_ENV: \"production\",
		},
	  }
	]
  }" > ecosystem.config.js

echo "done"
