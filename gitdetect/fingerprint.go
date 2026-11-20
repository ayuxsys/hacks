package main

var fingerPrints = []fingerPrint{
	{Pattern: "https://github.com/{org}", Cloud: Github},
	{Pattern: "https://github.com/{org}", Cloud: Gitlab},
}

type fingerPrint struct {
	Pattern string
	Cloud   GitCloudType
}
