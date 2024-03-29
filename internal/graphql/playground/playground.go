package playground

// This page is a copy of the GQLGen playground wrapper with client id & access token support:
// https://github.com/99designs/gqlgen/blob/0f016df3ae7ee4898358dc67a491689164297df6/graphql/playground/playground.go
//
// To update to the latest code, compare this file against (ignoring whitespace):
// https://github.com/99designs/gqlgen/blob/master/graphql/playground/playground.go

import (
	"html/template"
	"net/http"
	"net/url"
)

var page = template.Must(template.New("graphiql").Parse(`<!DOCTYPE html>
<html>
	<head>
		<title>{{.title}}</title>
		<link
			rel="stylesheet"
			href="https://cdn.jsdelivr.net/npm/graphiql@{{.version}}/graphiql.min.css"
			integrity="{{.cssSRI}}"
			crossorigin="anonymous"
		/>
	</head>
	<body style="margin: 0;">
		<div id="graphiql" style="height: 100vh;"></div>
		<script
			src="https://cdn.jsdelivr.net/npm/react@17.0.2/umd/react.production.min.js"
			integrity="{{.reactSRI}}"
			crossorigin="anonymous"
		></script>
		<script
			src="https://cdn.jsdelivr.net/npm/react-dom@17.0.2/umd/react-dom.production.min.js"
			integrity="{{.reactDOMSRI}}"
			crossorigin="anonymous"
		></script>
		<script
			src="https://cdn.jsdelivr.net/npm/graphiql@{{.version}}/graphiql.min.js"
			integrity="{{.jsSRI}}"
			crossorigin="anonymous"
		></script>
		<script>
{{- if .endpointIsAbsolute}}
		const url = {{.endpoint}};
		const subscriptionUrl = {{.subscriptionEndpoint}};
{{- else}}
		const url = location.protocol + '//' + location.host + {{.endpoint}};
		const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:';
		const subscriptionUrl = wsProto + '//' + location.host + {{.endpoint}};
{{- end}}
		const fetcher = GraphiQL.createFetcher({ url, subscriptionUrl });
		const headers = {
			"X-Client-ID": {{.clientID}},
		};
		const accessToken = new URL(window.location).searchParams.get("accessToken");
		if (accessToken) {
			headers["Authentication"] = "Bearer " + accessToken;
		}
		ReactDOM.render(
			React.createElement(GraphiQL, {
			fetcher: fetcher,
			tabs: true,
			headerEditorEnabled: true,
			shouldPersistHeaders: true,
			headers: JSON.stringify(headers, null, 2),
			defaultQuery: ` + "`" + `query AttackOnTitanShows {
	searchShows(search:"Attack on Titan", limit:5){
		id
		name
		createdAt
		createdBy {
			username
		}
		updatedAt
		seasonCount
		episodeCount
	}
}` + "`" + `
			}),
			document.getElementById('graphiql'),
		);
		</script>
	</body>
</html>
`))

// Handler responsible for setting up the playground
func Handler(title string, endpoint string, clientID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		err := page.Execute(w, map[string]interface{}{
			"title":                title,
			"endpoint":             endpoint,
			"endpointIsAbsolute":   endpointHasScheme(endpoint),
			"subscriptionEndpoint": getSubscriptionEndpoint(endpoint),
			"version":              "1.8.2",
			"cssSRI":               "sha256-CDHiHbYkDSUc3+DS2TU89I9e2W3sJRUOqSmp7JC+LBw=",
			"jsSRI":                "sha256-X8vqrqZ6Rvvoq4tvRVM3LoMZCQH8jwW92tnX0iPiHPc=",
			"reactSRI":             "sha256-Ipu/TQ50iCCVZBUsZyNJfxrDk0E2yhaEIz0vqI+kFG8=",
			"reactDOMSRI":          "sha256-nbMykgB6tsOFJ7OdVmPpdqMFVk4ZsqWocT6issAPUF0=",
			"clientID":             clientID,
		})
		if err != nil {
			panic(err)
		}
	}
}

// endpointHasScheme checks if the endpoint has a scheme.
func endpointHasScheme(endpoint string) bool {
	u, err := url.Parse(endpoint)
	return err == nil && u.Scheme != ""
}

// getSubscriptionEndpoint returns the subscription endpoint for the given
// endpoint if it is parsable as a URL, or an empty string.
func getSubscriptionEndpoint(endpoint string) string {
	u, err := url.Parse(endpoint)
	if err != nil {
		return ""
	}

	switch u.Scheme {
	case "https":
		u.Scheme = "wss"
	default:
		u.Scheme = "ws"
	}

	return u.String()
}
