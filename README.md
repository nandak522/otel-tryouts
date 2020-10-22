# otel-tryouts
otel-tryouts


# Svc Calls
Check `OTEL.png`

# Instructions
1. Generate Newrelic's `Tracing/Insights` Key and update `NEW_RELIC_API_KEY` in `docker-compose.yaml`
2. docker-compose build && docker-compose up frontend backend-tweets backend-notifications
3. Go to one.newrelic.com and view the traces under `Distributed Tracing`. You should find traces like the one mentioned in `newrelic-tracing.png` screenshot.
