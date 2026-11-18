package scope

import (
	"encoding/json"
	"net/url"
	"sn/log"
	"sn/utils/must"
	"strings"
)

func detectTypeWithHttpPrefix(value string) {
	var sType Type
	u, err := url.Parse(value)
	if err != nil {
		log.Warn.Printf("parse scope value with prefix 'http' %q as url: %v", value, err)
	}
	switch {
	case wildcardRegex.MatchString(u.Host):
		sType = WildCard
	case isGitURL(u):
		sType = Git
	case isIOSApp(u):
		sType = IOS
	case isAndroidAPP(u):
		sType = Android
	case isNPM(u):
		sType = NPM
	case isPyPi(u):
		sType = PyPi
	case isGoPkg(u):
		sType = GoPackage
	default:
		sType = URL
	}

	log.Output.Println(
		string(must.Eval(json.Marshal(New(value, sType)))),
	)
}

func isGitURL(u *url.URL) bool {
	host := u.Host
	lenPath := len(strings.Split(u.Path, "/"))

	switch host {
	// github.com/<owner>/<repo>
	case "github.com":
		return lenPath >= 1

	// gitlab.com/<group>/<project>
	// Supports nested groups.
	case "gitlab.com":
		return lenPath >= 1

	// bitbucket.org/<workspace>/<repo>
	case "bitbucket.org":
		return lenPath >= 1

	// codeberg.org/<owner>/<repo>
	case "codeberg.org":
		return lenPath >= 1

	// git.sr.ht/~user/project
	case "git.sr.ht":
		return lenPath >= 1

	// gitea.com/<owner>/<repo>
	case "gitea.com":
		return lenPath >= 1
	}

	return false
}

func isAndroidAPP(u *url.URL) bool {
	return u.Host == "play.google.com"
}

func isIOSApp(u *url.URL) bool {
	return u.Host == "apps.apple.com"
}

func isNPM(u *url.URL) bool {
	return u.Host == "www.npmjs.com" && strings.Contains(u.Path, "package")
}

func isGoPkg(u *url.URL) bool {
	return u.Host == "pkg.go.dev"
}

func isPyPi(u *url.URL) bool {
	return u.Host == "pypi.org" && strings.Contains(u.Path, "project")
}
