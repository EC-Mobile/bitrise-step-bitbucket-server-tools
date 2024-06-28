# Bitbucket And Bitrise Tools

This step allows you to perform different functions on Services like Bitbucket Server, Bitrise, etc..
With time more function will be added to same repo.

## Common Inputs

Bitbucket Inputs:
- BITBUCKET_API_ACCESS_TOKEN
- BITBUCKET_DOMAIN
- BITBUCKET_PROJECT_KEY
- BITBUCKET_REPO_SLUG

Bitrise Inputs:
- BITRISE_API_ACCESS_TOKEN
- BITRISE_APP_SLUG

Function Inputs:
- FUNCTION
- All params required by function. Check functions details below.

## Available Functions

1. Skip Verification

# Skip Verification
This step will check if the PR Title contains "[SV]" tag, it will stop the build on bit rise and add the commit status "FAILED" on the head commit of the PR.

## Required Inputs
- PR_ID:                    Bitrise PR build expose PR id as param: $BITRISE_PULL_REQUEST.

### Outputs
- SKIPPED_VERIFICATION:        True if found [SV] tag and aborted build, otherwise False.


# How to use this Step

Add this in your bitrise.yml file and replace proper variables:

```
- git::https://github.com/KageRiyuu/bitrise-step-bitbucket-server-tools.git@main:
    title: Bitbucket Server Tools
    inputs:
    - PR_ID: $BITRISE_PULL_REQUEST
```
Check above for all needed inputs