package bitrise

import env "bitbucket-tools/Source/Environment"

func GetToken() string {
	return env.BITRISE_API_ACCESS_TOKEN
}

func GetUrl() string {
	return "https://api.bitrise.io/v0.1/"
}

func GetAppUrl() string {
	return GetUrl() + "apps/" + env.BITRISE_APP_SLUG
}

func GetBuildsUrl() string {
	return GetAppUrl() + "/builds/"
}

func GetBuildUrl() string {
	return "https://app.bitrise.io/build/" + env.BITRISE_BUILD_SLUG
}
