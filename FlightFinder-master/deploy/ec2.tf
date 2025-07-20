locals {
  # shebang must not be indented!!!
  user_data = <<EOT
#!/usr/bin/env bash

# install docker
curl -fsSL https://get.docker.com | sh

# install systemd service
cat << EOF > /etc/systemd/system/flight-finder.service
[Unit] 
Description=Flight Finder Web Server 
After=network.target 

[Service] 
Type=simple 
Restart=always  
ExecStart=docker run --rm --name=flight-finder -p=80:80 -v $HOME/.aws:/root/.aws mateuszmidor/flight-finder:latest
ExecStop=docker stop flight-finder 

[Install] 
WantedBy=multi-user.target
EOF

# run systemd service
systemctl daemon-reload    
systemctl enable flight-finder
systemctl start flight-finder
systemctl status flight-finder
EOT
}

resource "aws_instance" "flight-finder" {
  ami             = var.ami
  instance_type   = "t2.micro"
  key_name        = var.key_pair_name
  security_groups = [aws_security_group.allow_ssh.name, aws_security_group.allow_http.name, aws_security_group.allow_egress.name]
  user_data       = local.user_data
  tags            = { "Name" : "FlightFinder" }
}

# need sercurity group with incoming ssh allowed; by default all traffic is disallowed
resource "aws_security_group" "allow_ssh" {
  name        = "allow_ssh"
  description = "Allow ssh traffic"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# need sercurity group with incoming http allowed
resource "aws_security_group" "allow_http" {
  name        = "allow_http"
  description = "Allow http traffic"

  # to allow access to flight-finder web server from the web
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# need sercurity group with all traffic outgoing
resource "aws_security_group" "allow_egress" {
  name        = "allow_egress"
  description = "Allow all outgoing traffic"

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}
