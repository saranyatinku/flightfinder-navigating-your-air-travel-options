output "ip" {
  value = aws_instance.flight-finder.public_ip
}

output "dns" {
  value = aws_instance.flight-finder.public_dns
}

output "ssh" {
  value = "ssh ubuntu@${aws_instance.flight-finder.public_ip} -i <key-pair-name.pem"
}
