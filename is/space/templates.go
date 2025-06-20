package space

const (
	// BaseStationMainTfTmpl is the main.tf of the space/base-station terraform module
	// This is a template that is rendered as a part of 'space new'
	BaseStationMainTfTmpl = `#: GENERATED
#: BASE STATION is the remote setup for managing our space/environments

locals {
  space = {
    name       = "{{.SpaceName}}",
    account_id = "{{.SpaceAccountID}}",
    region     = "{{.SpaceRegion}}",
  }
  root = {
    account_id = "{{.RootAccountID}}",
    region     = "{{.RootRegion}}",
  }
}

# Setup state in the root account
module "state" {
  source   = "../../modules/remote-state"
  name     = local.space.name
  region   = local.space.region
}
`

	// BaseStationOutputsTfTmpl is the outputs.tf of the space/base-station terraform module
	BaseStationOutputsTfTmpl = `#: GENERATED
#: BASE STATION is the remote setup for managing our space/environments

output "config" {
  value = {
    region         = module.state.region
    bucket         = module.state.bucket.id
    dynamodb_table = module.state.dynamodb_table.name
    kms_key_id     = module.state.kms_alias.arn
    encrypt        = true
    role_arn       = module.state.role_arn
  }
}
`

	// BaseStationVersionsTfTmpl is the versions.tf of the space/base-station terraform module
	BaseStationVersionsTfTmpl = `#: GENERATED
#: BASE STATION is the remote setup for managing our space/environments

terraform {
  required_version = ">= 1.6"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.25, <= 6.0"
    }
  }
}
`

	// BaseStationProvidersTfTmpl is the providers.tf of the space/base-station terraform module
	BaseStationProvidersTfTmpl = `#: GENERATED
#: BASE STATION is the remote setup for managing our space/environments

# AWS Provider for working with Root Account resources
provider "aws" {
  region              = local.root.region
  allowed_account_ids = [local.root.account_id]
}
`
)
