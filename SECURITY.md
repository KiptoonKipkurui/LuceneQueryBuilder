# Security Policy

## Supported Versions

This project is currently maintained on the `main` branch only.

| Version | Supported |
|--------|-----------|
| main   | ✅ Yes     |
| older releases | ❌ No |

---

## Reporting a Vulnerability

If you discover a security vulnerability, **please do not open a public GitHub issue**.

Instead, report it responsibly by contacting:

**Email:** kiptoonkipkurui@gmail.com

Please include:
- A clear description of the vulnerability
- Steps to reproduce the issue
- Any relevant code snippets or proof-of-concept
- The potential impact (e.g. query injection, malformed output, denial of service)

---

## Response Process

Once a vulnerability is reported:
1. It will be acknowledged within **72 hours**
2. The issue will be investigated and validated
3. A fix will be developed and tested
4. A patch or release will be published if required

If the issue is confirmed, we kindly ask that you **avoid public disclosure** until a fix is available.

---

## Scope

This library focuses on **query construction only**. It does **not**:
- Execute queries
- Communicate with external services
- Handle authentication or authorisation

However, malformed query generation, unsafe escaping, or injection-style issues are considered **in scope** and should be reported.

---

Thank you for helping keep this project secure.
