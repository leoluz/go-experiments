package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

func main() {

	logs, err := base64.StdEncoding.DecodeString(getB64Logs())
	if err != nil {
		fmt.Printf("error decoding event logs: %v", err)
		return
	}
	//l := strings.ReplaceAll(string(logs), "\\n", "\n")
	//l = strings.ReplaceAll(l, "\\t", "\t")
	str := fmt.Sprintf("\"%s\"", string(logs))
	fmt.Println(str)
	l, err := strconv.Unquote(str)
	if err != nil {
		fmt.Printf("error unquoting logs: %v", err)
		return
	}
	fmt.Printf("> Error logs:\n%s", l)
}

func getB64Logs() string {
	return "TG9ncyBmcm9tIGNvbnRhaW5lcjogaGVsbG8td29ybGQtc3ByaW5nLW1pY3Jvc2VydmljZS1jZGJldGEtc21va2UtdGVzdAp7ImxldmVsIjoiaW5mbyIsIm1zZyI6IlJ1bm5pbmcgdGVzdHMiLCJ0aW1lIjoiMjAxOS0wNC0wOVQwMjoxOTozNVoifQp7ImxldmVsIjoiZmF0YWwiLCJtc2ciOiIxMzo1NTo0Miw3OTIgfC1JTkZPIGluIGNoLnFvcy5sb2diYWNrLmNsYXNzaWMuam9yYW4uYWN0aW9uLkpNWENvbmZpZ3VyYXRvckFjdGlvbiAtIGJlZ2luXG4xMzo1NTo0Myw2ODkgfC1JTkZPIGluIGNoLnFvcy5sb2diYWNrLmNvcmUuam9yYW4uYWN0aW9uLkFwcGVuZGVyQWN0aW9uIC0gQWJvdXQgdG8gaW5zdGFudGlhdGUgYXBwZW5kZXIgb2YgdHlwZSBbY2gucW9zLmxvZ2JhY2suY29yZS5Db25zb2xlQXBwZW5kZXJdXG4xMzo1NTo0Myw3ODYgfC1JTkZPIGluIGNoLnFvcy5sb2diYWNrLmNvcmUuam9yYW4uYWN0aW9uLkFwcGVuZGVyQWN0aW9uIC0gTmFtaW5nIGFwcGVuZGVyIGFzIFtTVERPVVRdXG4xMzo1NTo0Myw3OTMgfC1JTkZPIGluIGNoLnFvcy5sb2diYWNrLmNvcmUuam9yYW4uYWN0aW9uLk5lc3RlZENvbXBsZXhQcm9wZXJ0eUlBIC0gQXNzdW1pbmcgZGVmYXVsdCB0eXBlIFtjaC5xb3MubG9nYmFjay5jbGFzc2ljLmVuY29kZXIuUGF0dGVybkxheW91dEVuY29kZXJdIGZvciBbZW5jb2Rlcl0gcHJvcGVydHlcbjEzOjU1OjQ0LDE5MCB8LUlORk8gaW4gY2gucW9zLmxvZ2JhY2suY2xhc3NpYy5qb3Jhbi5hY3Rpb24uUm9vdExvZ2dlckFjdGlvbiAtIFNldHRpbmcgbGV2ZWwgb2YgUk9PVCBsb2dnZXIgdG8gSU5GT1xuMTM6NTU6NDQsMTkwIHwtSU5GTyBpbiBjaC5xb3MubG9nYmFjay5jb3JlLmpvcmFuLmFjdGlvbi5BcHBlbmRlclJlZkFjdGlvbiAtIEF0dGFjaGluZyBhcHBlbmRlciBuYW1lZCBbU1RET1VUXSB0byBMb2dnZXJbUk9PVF1cbjEzOjU1OjQ0LDE5MSB8LUlORk8gaW4gY2gucW9zLmxvZ2JhY2suY2xhc3NpYy5qb3Jhbi5hY3Rpb24uQ29uZmlndXJhdGlvbkFjdGlvbiAtIEVuZCBvZiBjb25maWd1cmF0aW9uLlxuMTM6NTU6NDQsMTkxIHwtSU5GTyBpbiBjaC5xb3MubG9nYmFjay5jbGFzc2ljLmpvcmFuLkpvcmFuQ29uZmlndXJhdG9yQDZhZDVjMDRlIC0gUmVnaXN0ZXJpbmcgY3VycmVudCBjb25maWd1cmF0aW9uIGFzIHNhZmUgZmFsbGJhY2sgcG9pbnRcbjIwMTktMDMtMjYgMTM6NTU6NDQsMTk0WyBsZXZlbD1JTkZPIF1bIGxvZ2dlcj1jLmEubS5zLnMudC5TZWdtZW50U21va2VUZXN0cyBdOiBCYXNlIFVSTCA9IGh0dHA6Ly9jYW5hcnktMTA3LXNlZ21lbnRcbjIwMTktMDMtMjYgMTM6NTY6MDUsNzg1WyBsZXZlbD1FUlJPUiBdWyBsb2dnZXI9Yy5hLm0ucy5zLnQuU2VnbWVudFNtb2tlVGVzdHMgXTogU21va2UgdGVzdHMgZmFpbGVkXG5qYXZhLm5ldC5Vbmtub3duSG9zdEV4Y2VwdGlvbjogY2FuYXJ5LTEwNy1zZWdtZW50OiBOYW1lIGRvZXMgbm90IHJlc29sdmVcblx0YXQgamF2YS5uZXQuSW5ldDZBZGRyZXNzSW1wbC5sb29rdXBBbGxIb3N0QWRkcihOYXRpdmUgTWV0aG9kKVxuXHRhdCBqYXZhLm5ldC5JbmV0QWRkcmVzcyQyLmxvb2t1cEFsbEhvc3RBZGRyKEluZXRBZGRyZXNzLmphdmE6OTI4KVxuXHRhdCBqYXZhLm5ldC5JbmV0QWRkcmVzcy5nZXRBZGRyZXNzZXNGcm9tTmFtZVNlcnZpY2UoSW5ldEFkZHJlc3MuamF2YToxMzIzKVxuXHRhdCBqYXZhLm5ldC5JbmV0QWRkcmVzcy5nZXRBbGxCeU5hbWUwKEluZXRBZGRyZXNzLmphdmE6MTI3Nilcblx0YXQgamF2YS5uZXQuSW5ldEFkZHJlc3MuZ2V0QWxsQnlOYW1lKEluZXRBZGRyZXNzLmphdmE6MTE5Milcblx0YXQgamF2YS5uZXQuSW5ldEFkZHJlc3MuZ2V0QWxsQnlOYW1lKEluZXRBZGRyZXNzLmphdmE6MTEyNilcblx0Li4uXG4iLCJ0aW1lIjoiMjAxOS0wNC0wOVQwMjoxOTozNVoifQoK"
}

func getLogs() string {
	return `"13:55:42,792 |-INFO in ch.qos.logback.classic.joran.action.JMXConfiguratorAction - begin\n13:55:43,689 |-INFO in ch.qos.logback.core.joran.action.AppenderAction - About to instantiate appender of type [ch.qos.logback.core.ConsoleAppender]\n13:55:43,786 |-INFO in ch.qos.logback.core.joran.action.AppenderAction - Naming appender as [STDOUT]\n13:55:43,793 |-INFO in ch.qos.logback.core.joran.action.NestedComplexPropertyIA - Assuming default type [ch.qos.logback.classic.encoder.PatternLayoutEncoder] for [encoder] property\n13:55:44,190 |-INFO in ch.qos.logback.classic.joran.action.RootLoggerAction - Setting level of ROOT logger to INFO\n13:55:44,190 |-INFO in ch.qos.logback.core.joran.action.AppenderRefAction - Attaching appender named [STDOUT] to Logger[ROOT]\n13:55:44,191 |-INFO in ch.qos.logback.classic.joran.action.ConfigurationAction - End of configuration.\n13:55:44,191 |-INFO in ch.qos.logback.classic.joran.JoranConfigurator@6ad5c04e - Registering current configuration as safe fallback point\n2019-03-26 13:55:44,194[ level=INFO ][ logger=c.a.m.s.s.t.SegmentSmokeTests ]: Base URL = http://canary-107-segment\n2019-03-26 13:56:05,785[ level=ERROR ][ logger=c.a.m.s.s.t.SegmentSmokeTests ]: Smoke tests failed\njava.net.UnknownHostException: canary-107-segment: Name does not resolve\n\tat java.net.Inet6AddressImpl.lookupAllHostAddr(Native Method)\n\tat java.net.InetAddress$2.lookupAllHostAddr(InetAddress.java:928)\n\tat java.net.InetAddress.getAddressesFromNameService(InetAddress.java:1323)\n\tat java.net.InetAddress.getAllByName0(InetAddress.java:1276)\n\tat java.net.InetAddress.getAllByName(InetAddress.java:1192)\n\tat java.net.InetAddress.getAllByName(InetAddress.java:1126)\n\t...\n"`
}
