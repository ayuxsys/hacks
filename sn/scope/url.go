package scope

import (
	"net/url"
	"strings"
)

func DetectTypeWithHttpPrefix(u *url.URL) Type {
	switch {
	case isWildCard(u):
		return WildCard
	case isGitURL(u):
		return Git
	case isIOSApp(u):
		return IOS
	case isAndroidAPP(u):
		return Android
	case isNPM(u):
		return NPM
	case isPyPi(u):
		return PyPi
	case isGoPkg(u):
		return GoPackage
	default:
		return URL
	}
}

func isWildCard(parsedU *url.URL) bool {
	return strings.HasPrefix(parsedU.Host, "*")
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
	trailingPath, _, _ := strings.Cut(u.Path, "/")
	return u.Host == "www.npmjs.com" && (trailingPath == "package" || trailingPath == "search")
}

func isGoPkg(u *url.URL) bool {
	return u.Host == "pkg.go.dev"
}

func isPyPi(u *url.URL) bool {
	trailingPath, _, _ := strings.Cut(u.Path, "/")
	return u.Host == "pypi.org" && trailingPath == "project"
}
