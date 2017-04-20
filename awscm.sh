#!/bin/bash
#
# AWSCM: Command line tool for quickly switching between AWS profiles.

OUTPUT_FORMATS=(
  text
  json
  table
)

REGIONS=(
  ap-northeast-1
  ap-southeast-1
  ap-southeast-2
  eu-central-1
  eu-west-1
  eu-west-2
  sa-east-1
  us-east-1
  us-west-1
  us-west-2
)

function awscm() {
  if [ -z "$1" ]; then
    echo "No command supplied. Use:"
    echo "'awscm add'"
    echo "'awscm configure'"
    echo "'awscm list'"
    echo "'awscm output'"
    echo "'awscm region'"
    echo "'awscm status'"
    echo "'awscm export'"
    echo "'awscm use'"

    return 0
  fi

  case "$1" in
    "add") aws_add "$2" ;;
    "configure") aws_configure ;;
    "list") aws_list "$2" ;;
    "output") aws_output "$2" ;;
    "region") aws_region "$2" ;;
    "status") aws_status ;;
    "export") aws_export_variables ;;
    "use") aws_use "$2" ;;
    *) echo "Unknown command" ;;
  esac
}

function aws_add() {

  if [ -z "$1" ]; then
    echo "No profile name supplied."
  else
    if grep -q "$1" ~/.aws/credentials; then
      echo "Updating the AWS profile [${1}]:"
    else
      echo "Creating the AWS profile [${1}]:"
    fi
    aws configure --profile "${1}"
  fi
}

function aws_configure() {
  echo "Configuring the default AWS profile:"
  aws configure
}

function aws_list() {
  if [ -z "$1" ]; then
    echo "Usage: 'awscm list config', 'aws list credentials'"
  else
    if [[ "$1" == "config" ]]; then
      cat ~/.aws/config
    elif [[ "$1" == "credentials" ]] || [[ "$1" == "creds" ]]; then
      cat ~/.aws/credentials
    else
      echo "Usage: 'awscm list config', 'aws list credentials'"
    fi
  fi
}

function aws_output() {

  if [ -z "$1" ]; then
    echo "No output format supplied"
  else
    if is_output_format_valid "$1"; then
      export AWS_DEFAULT_OUTPUT=${1}
      echo "AWS command line output format set to '${1}'"
    else
      echo "The output format supplied, '${1}', is not supported."
      echo "Please use an output format from:"
      for output_format in "${OUTPUT_FORMATS[@]}"; do
        echo -e "\t $output_format"
      done
    fi
  fi
}

function aws_region() {

  if [[ -z "$1" ]]; then
    echo "No region supplied"
  else
    if is_region_valid "$1"; then
      export AWS_DEFAULT_REGION=${1}
      echo "AWS command line region set to '${1}'"
    else
      echo "The region supplied, '${1}', is not recognised."
      echo "Please use a region from:"
      for region in "${REGIONS[@]}"; do
        echo -e "\t $region"
      done
    fi
  fi
}

function aws_status() {

  if [ -z "$AWS_DEFAULT_PROFILE" ]; then
    echo "AWS profile currently unset."
  else
    echo "AWS profile set to: [$AWS_DEFAULT_PROFILE]."
  fi
  if [ -z "$AWS_DEFAULT_REGION" ]; then
    echo "AWS region currently unset, using profile default."
  else
    echo "AWS region set to: '$AWS_DEFAULT_REGION'."
  fi
  if [ -z "$AWS_DEFAULT_OUTPUT" ]; then
    echo "AWS output format currently unset, using profile default."
  else
    echo "AWS output format set to: '$AWS_DEFAULT_OUTPUT'."
  fi

}

function aws_use() {

  if [ -z "$1" ]; then
    echo "No environment supplied"
  else
    if grep -q "\[$1\]" ~/.aws/credentials; then
      export AWS_DEFAULT_PROFILE=${1}
      echo "AWS command line environment set to [${1}]"
    else
      echo "AWS profile [${1}] not found."
      echo "Please choose from an existing profile:"
      grep "\[" ~/.aws/credentials
      echo "Or create a new one with:"
      echo "'awscm add ${1}'"
    fi
  fi
}

function is_output_format_valid() {

    for output_format in "${OUTPUT_FORMATS[@]}"; do
      if [[ $output_format == "$1" ]]; then
        return 0
      fi
    done
    return 1
}


function is_region_valid() {

  for region in "${REGIONS[@]}"; do
    if [[ $region == "$1" ]]; then
      return 0
    fi
  done
  return 1
}

function aws_export_variables() {

  if [ -z "$1" ]; then
    echo "No environment supplied"
  else
    if grep -q "\[$1\]" ~/.aws/credentials; then
      export AWS_DEFAULT_PROFILE=${1}
      export AWS_PROFILE=${1}
      declare -a env_var_fields=("AWS_ACCESS_KEY_ID" "AWS_SECRET_ACCESS_KEY" "AWS_SESSION_TOKEN")
      for var in "${env_var_fields[@]}"
      do
        lcvar=$(echo $var | tr '[:upper:]' '[:lower:]')
        expval=$(aws configure get "${1}.$lcvar")
        # echo "$var=$expval"
        export $var=$expval
      done
      echo "AWS command line variables exported for environment [${1}]"
    else
      echo "AWS profile [${1}] not found."
      echo "Please choose from an existing profile:"
      grep "\[" ~/.aws/credentials
      echo "Or create a new one with:"
      echo "'awscm add ${1}'"
    fi
  fi
}
