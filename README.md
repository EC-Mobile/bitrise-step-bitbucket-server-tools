# Bitbucket And Bitrise Tools

This step allows you to perform different functions on Services like Bitbucket Server, Bitrise, etc..
With time more function will be added to same repo.

## Common Inputs

Bitbucket Inputs:
- bitbucket_api_access_token
- bitbucket_domain
- bitbucket_project_key
- bitbucket_repo_slug

Bitrise Inputs:
- bitrise_api_access_token
- bitrise_app_slug

Function Inputs:
- function
- All params required by function. Check functions details below.

## Available Functions

1. Skip Verification

# Skip Verification
This step will check if the PR Title contains "[SV]" tag, it will stop the build on bit rise and add the commit status "FAILED" on the head commit of the PR.

## Required Inputs
- pr_id:                    Bitbucket PR id, e.g: $BITRISE_PULL_REQUEST.
- bitrise_build_slug:       Build id on bitrise, e.g: $BITRISE_BUILD_SLUG

### Outputs
- SKIPPED_VERIFICATION:        True if found [SV] tag and aborted build, otherwise False.


# How to use this Step

Add this in your bitrise.yml file and replace proper variables:

```
- git::https://github.com/KageRiyuu/bitrise-step-bitbucket-server-tools.git@main:
    title: Bitbucket Server Tools
    inputs:
    - pr_id: $BITRISE_PULL_REQUEST
```
Check above for all needed inputs