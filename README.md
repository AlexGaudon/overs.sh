# OverSSH

## What is OverSSH?

OverSSH is a simple way to transfer files between computers, without having to open any ports.

## Why did I make it?

I was tired of using `scp` to copy things between my laptop and various servers, especially when one server was on an external network the necessary ports open to facilitate `scp` or `sftp`.

# Getting Started

You need to configure a few things. After you install [Golang](https://go.dev/doc/install) on your system, you will need to generate an ssh key.

The easiest way to do this is with the `ssh-keygen` command.

It will ask you where to save the key, enter `./key`

You now have a .key and .key.pub file.

The application will automatically read the .key file if it exists, and use a random key if one isn't created.

Note: Not generating a key will lead to a warning of a possible man in the middle attack, and you will have to remove the entry from your `known_hosts` file every time you restart the server.

You will also need a SSL certificate, you can provide your own or generate one using [Certbot](https://certbot.eff.org/).

Store both the private (privkey.pem) and public (fullchain.pem) files in this directory.

Lastly, set an environment variable called `URL` with the domain that you'll be hosting the application on. (Example: https://overs.sh)

After you have all of the configuration done, simply run `make prod` to run a production instance of the application, or `make dev` for a development instance.

# FAQs

## Is it "production ready"?

Probably not. I've used it for approximately 6 months and haven't run into any major blocks, but my usecase is usually limited to copying a single config file, or sometimes a .tar.gz file, usually maxing out at around 10-15 MB.

## Known Issues

The file size that can be transferred is limited by the memory on the server. Running on a $5 VPS (1 vCPU, 1GB RAM) on Linode I can easily transfer files up to around 700 MB.