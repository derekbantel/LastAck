# LastAck

![Status](https://img.shields.io/badge/status-alpha-orange)
![License](https://img.shields.io/badge/license-MIT-blue)
![Docker](https://img.shields.io/badge/docker-ready-green)

**LastAck** is a self-hosted, digital dead man's switch.

It ensures that if you (the sysadmin) suddenly become unavailable, your critical information (password manager master keys, server access details, crypto seeds, or personal letters) is securely decrypted and delivered to your trusted contacts.

> **The "Bus Factor" Solution:** If you were hit by a bus tomorrow, would your family be able to access your digital life?

## 🚀 Features

* **Zero-Knowledge Storage:** Secrets are stored using **AES-256** encryption. LastAck does not store the decryption key in the database; it is derived from your configuration/env vars.
* **Multiple "Pulse" Methods:** Reset your timer via:
    * Web UI (One-click "I'm Alive" button).
    * Webhook (Integration with Home Assistant, Uptime Kuma, or CRON).
    * Email Link.
* **Escalation Policies:** Define a grace period (e.g., *Warning Email* after 30 days -> *Release Secrets* after 35 days).
* **Docker First:** Built to run in a container with minimal footprint.
* **SMTP Support:** Works with any standard SMTP provider (SendGrid, Mailgun, or your local Postfix).

## 🛠️ How It Works

1.  **Setup:** You define a **Timeout Period** (e.g., 14 days).
2.  **Vault:** You input text or upload files into the encrypted vault and define **Recipients**.
3.  **Check-in:** As long as you "Check In" before the timeout, the timer resets.
4.  **The Trigger:** If the timer reaches zero, LastAck decrypts the vault and emails the contents to your recipients.
