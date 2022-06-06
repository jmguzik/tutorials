## Requests per agent

`rate, watch out for labels (eg ghproxy reset)`

`rate(github_request_duration_count{status=~"403"}[10m])`

### sum, do not care about user agents
`sum(rate(github_request_duration_count{status=~"403"}[10m]))`

### sum, summarize on user_agent
`sum(rate(github_request_duration_count{status=~"403"}[10m])) by (user_agent)`


## Concurrent requests
`pending_outbound_requests`

`concurrent_outbound_requests`

`sum_over_time(pending_outbound_requests{container="ghproxy"}[10m]) / count_over_time(pending_outbound_requests{container="ghproxy"}[10m]) > 10`

## Plugins, took action

### median
`histogram_quantile(0.5, sum(rate(prow_plugin_handle_duration_seconds_bucket{took_action="true"}[30m])) by (le))`

### 95% percent by plugin
`histogram_quantile(0.95, sum(rate(prow_plugin_handle_duration_seconds_bucket{took_action="true"}[30m])) by (le, plugin))`
