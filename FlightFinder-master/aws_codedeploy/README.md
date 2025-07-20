# AWS CodeDeploy

Numbered files in this folder are scripts for Application Lifecycle Hooks, referenced from `appspec.yml`.

## Troubleshooting

### Application Lifecycle Logs (install/start/validate/stop)

- on EC2 instance: `/opt/codedeploy-agent/deployment-root/deployment-logs/codedeploy-agent-deployments.log`

### CodeDeploy Agent (on EC2)

- agent service status: `systemctl status codedeploy-agent`
- agent service logs: `journalctl -u codedeploy-agent`
