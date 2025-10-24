# env config
## log on
- export ENABLE_LOGGING=true

## discord
>default type is discord

- export NOTIFIER_TYPE=discord
- export DISCORD_WEBHOOK=""

## telegram
- export NOTIFIER_TYPE=telegram
- export TELEGRAM_BOT_TOKEN="83921057"
- export TELEGRAM_CHAT_ID="-482263"

# recived json

```json
{"receiver":"monitoring/dengfeng-alertmanager-config/discordNotify","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"RedisDown","endpoint":"metrics","instance":"10.0.42.54:9000","job":"external-redis-exporter","namespace":"monitoring","prometheus":"monitoring/k8s","service":"external-redis-exporter","severity":"critical"},"annotations":{"description":"Redis instance is down\n  VALUE = 1\n  LABELS = map[__name__:redis_up endpoint:metrics instance:10.0.42.54:9000 job:external-redis-exporter namespace:monitoring service:external-redis-exporter]","summary":"Redis down (instance 10.0.42.54:9000)"},"startsAt":"2025-10-24T03:32:02.839Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://prometheus-k8s-0:9090/graph?g0.expr=redis_up+%3D%3D+1\u0026g0.tab=1","fingerprint":"c4da66c7c14de159"}],"groupLabels":{"alertname":"RedisDown","severity":"critical"},"commonLabels":{"alertname":"RedisDown","endpoint":"metrics","instance":"10.0.42.54:9000","job":"external-redis-exporter","namespace":"monitoring","prometheus":"monitoring/k8s","service":"external-redis-exporter","severity":"critical"},"commonAnnotations":{"description":"Redis instance is down\n  VALUE = 1\n  LABELS = map[__name__:redis_up endpoint:metrics instance:10.0.42.54:9000 job:external-redis-exporter namespace:monitoring service:external-redis-exporter]","summary":"Redis down (instance 10.0.42.54:9000)"},"externalURL":"http://alertmanager-main-0:9093","version":"4","groupKey":"{}/{namespace=\"monitoring\"}/{severity=\"critical\"}:{alertname=\"RedisDown\", severity=\"critical\"}","truncatedAlerts":0}
```