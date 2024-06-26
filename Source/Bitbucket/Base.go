package bitbucket

import env "bitbucket-tools/Source/Environment"

func GetUrl() string {
	return "https://" + env.BITBUCKET_DOMAIN
}

func GetToken() string {
	return "Bearer " + env.BITBUCKET_ACCESS_TOKEN
}
