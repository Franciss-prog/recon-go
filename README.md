# recon-go# 🔍 ReconGo

> Automated domain reconnaissance — fast, simple, and built in Go.

ReconGo is a CLI-driven tool that automatically scans a domain to discover infrastructure information. Give it a domain name, and it hunts down subdomains, open ports, and active web services for you.

---

## Table of Contents

- [Features](#features)
- [How It Works](#how-it-works)
- [Installation](#installation)
- [Usage](#usage)
  - [Scanning a Domain](#scanning-a-domain)
  - [Checking Results](#checking-results)
- [What Gets Scanned](#what-gets-scanned)
  - [Subdomain Discovery](#1-subdomain-discovery)
  - [Port Scanning](#2-port-scanning)
  - [Web Service Detection](#3-web-service-detection)
- [Example Output](#example-output)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Subdomain Discovery** — Probes common subdomain names to find what's live
- **Port Scanning** — Checks common ports (80, 443, 22, and more) for open services
- **Web Service Detection** — Sends HTTP requests to confirm active web services
- **Async Scanning** — Scans run server-side; retrieve results when ready
- **Simple CLI** — Two commands to scan and retrieve results

---

## How It Works

```
User runs CLI command
        │
        ▼
Server receives domain
        │
        ▼
Server scans the domain
(subdomains → ports → web services)
        │
        ▼
Server returns results
```

Scans are handled asynchronously on the server. You kick off a scan, get a scan ID, and poll for results when you're ready.

---

## Installation

### Prerequisites

- [Go](https://go.dev/dl/) 1.21+
- Git

### Build from Source

```bash
# Clone the repository
git clone https://github.com/youruser/recongo.git
cd recongo

# Build the CLI
go build -o recongo ./cmd/cli

# Build the server
go build -o recongo-server ./cmd/server

# (Optional) Move CLI to your PATH
mv recongo /usr/local/bin/recongo
```

### Start the Server

```bash
./recongo-server
# Server listening on :8080
```

---

## Usage

### Scanning a Domain

Run a reconnaissance scan against a domain:

```bash
recongo scan example.com
```

**Response:**

```
Scan started.
Scan ID: 1
Use 'recongo results 1' to retrieve results.
```

The scan runs in the background on the server. You can continue working or wait a moment before checking results.

---

### Checking Results

Retrieve the results of a completed scan by its ID:

```bash
recongo results 1
```

**Response:**

```
Subdomains found:
  api.example.com
  dev.example.com

Open ports:
  api.example.com : 80
  api.example.com : 443

Web services detected:
  http://api.example.com  →  200 OK
  https://api.example.com →  200 OK
```

---

### Full Example Walkthrough

```bash
# 1. Start the server (in a separate terminal)
./recongo-server

# 2. Run a scan
recongo scan tesla.com
# → Scan ID: 1

# 3. Grab results
recongo results 1
```

---

## What Gets Scanned

### 1. Subdomain Discovery

ReconGo tries a list of common subdomain prefixes against the target domain:

```
api.example.com
dev.example.com
admin.example.com
mail.example.com
staging.example.com
www.example.com
...
```

If a subdomain resolves (DNS lookup succeeds) → it is recorded.

---

### 2. Port Scanning

For each discovered subdomain, ReconGo checks common ports:

| Port | Service    |
| ---- | ---------- |
| 80   | HTTP       |
| 443  | HTTPS      |
| 22   | SSH        |
| 21   | FTP        |
| 3306 | MySQL      |
| 5432 | PostgreSQL |
| 8080 | HTTP Alt   |
| 8443 | HTTPS Alt  |

If a port accepts a TCP connection → it is marked open.

---

### 3. Web Service Detection

For open HTTP/HTTPS ports, ReconGo makes a real HTTP request:

```
GET http://api.example.com
```

If the server responds → the web service is recorded along with its status code.

---

## Example Output

Input:

```bash
recongo scan example.com
```

Full output from `recongo results 1`:

```
=========================================
  ReconGo Results — example.com
=========================================

Subdomains found:
  ✓ api.example.com
  ✓ dev.example.com

Open ports:
  api.example.com : 80
  api.example.com : 443
  dev.example.com : 22
  dev.example.com : 80

Web services detected:
  http://api.example.com   → 200 OK
  https://api.example.com  → 200 OK
  http://dev.example.com   → 301 Moved

=========================================
Scan completed in 4.2s
```

---

## Architecture

ReconGo has two components:

### Client — Go CLI

The CLI you interact with. Sends requests to the server and displays results.

```
recongo scan <domain>     →  POST /scan
recongo results <id>      →  GET  /results/:id
```

### Server — Go HTTP Server

Runs the actual scanners. Manages scan jobs and returns results via a REST API.

```
POST /scan           Start a new scan
GET  /results/:id    Get results for a scan
```

---

## Tech Stack

| Component | Technology |
| --------- | ---------- |
| CLI       | Go         |
| Server    | Go         |
| Transport | HTTP/REST  |

---

## Contributing

Contributions are welcome!

1. Fork the repo
2. Create a feature branch: `git checkout -b feat/my-feature`
3. Commit your changes: `git commit -m "feat: add my feature"`
4. Push and open a Pull Request

Please keep PRs focused and include a description of what changed and why.

---

## License

MIT License. See [LICENSE](LICENSE) for details.

---

> Built with Go. Designed to be simple.
