# GitHub Activity Analyzer

A lightweight Go service that analyzes public GitHub activity events and extracts meaningful development metrics such as commit counts, repository usage, and push statistics using the GitHub REST API.

## 🚀 Features

Fetches public GitHub events for any user

Detects PushEvents and PullRequests events

Reconstructs real commit counts

Uses GitHub’s compare API to obtain accurate commit statistics

Designed to be fast, simple, and dependency-light

## 🏗 Architecture
```bash
Client
   │
   ▼
GitHub Events API
   │
   ▼
PushEvent Detector
   │
   ▼
Commit Range Comparator ──► GitHub Compare API
   │
   ▼
Commit Statistics Output
```

## 📦 Installation
```bash
git clone https://github.com/dondarrion91/github-activity
cd github-activity
go mod tidy
```

## ▶️ Usage
```bash
go build main.go
./main <github-username>
```

## 🔐 Security

Works safely with signed commits

No private data required

Does not store tokens or credentials

Fully public GitHub API compatible

## Roadmap.sh project
This project is based of a example project in [https://roadmap.sh/projects/github-user-activity](https://roadmap.sh/projects/github-user-activity)

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

