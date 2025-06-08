---
title: "My makeshift hosting setup"
date: "2025-06-06"
tags: ["hosting", "cloud", "kubernetes"]
excerpt: "A post about my current hosting setup"
---

Over the years, I've jumped between many different cloud providers to host my little pet projects. In theory, 90% of it could be hosted on static file hosting such as GitHub Pages, Cloudflare Pages, or even plain old S3. However, being the engineer that I am, I love over-engineering things. So we bring in Kubernetes, baby!

# Migration off the cloud and onto Raspberry PIs

My latest solution is a setup of Raspberry PIs, an old MacBook, and a few Proxmox VMs, all of different sizes and architectures to make things spicy. Using [k3s](https://k3s.io/) and a [few Ansible scripts](https://github.com/scott-the-programmer/pi-cluster-cm), I'm able to create and add new nodes to my cluster easily (provided I'm on my home network). Outside of power cuts and ISP disruptions, so far this has been pretty reliable and fast—at least if you're in the southern hemisphere.

# OK cool, what are you hosting exactly?

Well, none other than [term.nz](https://term.nz), [scott.murray.kiwi](https://scott.murray.kiwi), and a few APIs to fuel things. There's also a product that I'm working on to be revealed eventually!

# How do you expose it to the internet?

As much as I love doxxing my IP, Cloudflare Tunnel has been a game changer. All the public containers run with a [sidecar that tunnels through Cloudflare/Argo Tunnel](https://github.com/scott-the-programmer/meshed/blob/main/stacks/applications/apps/blog-api.go). The nice bit is that this keeps my IP private (_not that I own a static IP anyway_) but is also flexible enough that if I were to switch hosting, my tunnel config can stay exactly the same.

# You may be asking, how are you updating your apps?

That's where [keel.sh](https://keel.sh) comes in! I've simply added the [polling annotations](https://keel.sh/docs/#deployment-polling-example) to my deployments, and as soon as an image is updated in the respective repository, it gets refreshed via the same mechanism that was used to pull the containers in the first place. Obviously, most of my apps are open-source, but for the few that I'd rather keep private, Keel is seamless enough to work with those without any extra config.

# Different architectures, you say?

Probably the biggest pain point, and a great argument to not embark on this journey, is that each container image needs to be cross-architecture compatible. It is essentially a gamble on which node it gets allocated to, and unfortunately, many household name images (think databases) still don't support ARM. This means that each deployment and Helm chart needs to be verified on both architectures before releasing into the wild. Luckily, when there truly is a roadblock, simply adding the `kubernetes.io/arch: amd64` node selector can get around it, but it does defeat the point of the cluster in a sense.

# What's the point of all this?

Honestly, no idea. Cloud providers are expensive, and having old machines lying around makes for a perfect setup. Part of it is enjoying the journey and the problems that come up along the way, and the other part is knowing no machine has gone to waste.

cya
